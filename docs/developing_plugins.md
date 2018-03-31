# Developing plugins

The fastest way to start developing a plugin is to use the CLI tool.

install:   
`go install github.com/aunem/transpose`   

check out the help command:   
`transpose plugin init --help`   

example:   

```bash
PLUGIN_NAME="mynewplugin"
PLUGIN_PACKAGE="github.com/myuser/myrepo"
PLUGIN_TYPE="middleware"  # middleware, listener, or roundtrip allowed

transpose plugin init -n ${PLUGIN_NAME} -p ${PLUGIN_PACKAGE} -t ${PLUGIN_TYPE}
```
this will scaffold you out a good start for your plugin, now lets walk though the files present...


...in progress