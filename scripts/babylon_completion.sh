#!/bin/bash

####
#  Babylon post message execution script
#  Written by Darrell Breeden
#       04/01/20
#  Post Execution:
#  Failure: Slack notification via babbleon binary
#  Success: use BBI to write summary to file
####



###
# --additional_post_work_envs="BABBLEON_TARGET=@john.doe,BABBLEON_OAUTH_TOKEN=<insert_token>" from babylon will ensure these are set
# If the BABBLEON_OAUTH_TOKEN value is not set in the environment, slack notification will not be possible
# If the BABBLEON_TARGET value is not set, the slack API will not know what user to send to. BABBLEON_TARGET specifies a user as @<slackusername>
# BABBLEON_TARGET="@john.doe" for example.
###

###Babylon Variables
BBI=/data/apps/bbi
SUCCESS=$BABYLON_SUCCESSFUL
OUTPUT_DIR=$BABYLON_OUTPUT_DIR
MODEL=$BABYLON_MODEL


###AWS Variables
region=$(curl -s http://169.254.169.254/latest/meta-data/placement/availability-zone | rev | cut -c2- | rev)
instance_id=$(curl -s http://169.254.169.254/latest/meta-data/instance-id)
workflow=`aws ec2 describe-tags --filters "Name=resource-id,Values=${instance_id}" --region=$region | jq '.Tags[] | select(.Key == "aws:cloudformation:stack-name") | .Value' | sed 's/\"//g' | awk -F "-" '{print $2}'`


if [[ -z $SUCCESS ]] ; then
        #Somehow the env wasn't set, so we will assume failure
        SUCCESS=false
        echo "It appears that the BABYLON environment variables are not set. Cannot process"
        exit 1
fi



if [[ "$SUCCESS" = "true" ]] ; then
        SUMMARY_FILE="bbi_summary.txt"

        #Success. Run BBI Summary
        cd $OUTPUT_DIR || exit 1
        $BBI nonmem summary $MODEL > $SUMMARY_FILE
        exit 0
fi


if [[ "$SUCCESS" = "false" ]]; then
        #Failure. Send slack message
        babbleon --additional_message_values="Workflow: $workflow"
fi