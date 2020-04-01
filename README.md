# Babbleon

Babbleon is a simple notification tool which, although primarily built for Slack, is easily extensible due to the 
architecture used in the implementation. The core functionality is built around pre-structuring a message based on known
environment variables provided by Babylon's post-execution environment. 


## Environment Requirements

`BABBLEON_OAUTH_TOKEN` : The OAuth token to be used for authenticating to Slack
`BABBLEON_DEBUG`: Runs babbleon with additional logging output


## Commandline Parameters
`--target` : The channel or username reference (as `@username`) to send the message to.
`--additional_message_values`: A comma-separated slice of strings that are added to the message. Useful for collecting
and reporting cloud environment details, Grid manager details and anything else relevant to the business.

## Using with [Babylon](https://www.github.com/metrumresearchgroup/babylon)
When calling babbleon, you'll need to set the OAuth token somehow. But how!? Babylon has a `--additional_post_work_envs`
flag, which takes a comma delimited list of env variables. It will set those into the thread execution the 
post-execution script, it will make sure they're available:

Example:
```
bbi nonmem run --nm_version nm74 local --overwrite=true --post_work_executable /shared/path/to/babbleon --additional_post_work_envs "BABBLEON_OAUTH_TOKEN=<TOKEN>,BABBLEON_TARGET=@user" 001.mod
```
### Note
In the above scenario, we're running in Babylon local mode, but if you use SGE mode, you'll want to make sure that the path
for babbleon (or any other post execution script) is available at that location by all SGE workers.