# Babbleon

Babbleon is a simple notification tool which, although primarily built for Slack, is easily extensible due to the 
architecture used in the implementation. 


## Environment Requirements

`BABBLEON_DEBUG` : Sets loggers into debug mode
`BABBLEON_OAUTH_TOKEN` : The OAuth token to be used for authenticating to Slack

## Commandline Parameters
`--message` : The body (content) of the Slack message
`--target` : The channel or username reference (as `@username`) to send the message to.



## Using with Babylon
When calling babbleon, you'll need to set the OAuth token somehow. But how!? Babylon has a `--additional_post_work_envs`
flag, which takes a comma delimited list of env variables. It will set those into the thread execution the 
post-execution script, it will make sure they're available:

Example:
```
bbi nonmem run --nm_version nm74 local --overwrite=true --post_work_executable babbleon --additional_post_work_envs "BABBLEON_OAUTH_TOKEN=<TOKEN>,BABBLEON_TARGET=@user,BABBLEON_MESSAGE=this is the message we're sending" 001.mod
```
