package main

import (
	"flag"
	"strings"

	common "github.com/dokku/dokku/plugins/common"
)

// runs 'git gc --agressive' against the application's repo
func gitGC() {
	flag.Parse()
	appName := flag.Arg(1)
	if appName == "" {
		common.LogFail("Please specify an app to run the command on")
	}
	common.VerifyAppName(appName)

	appRoot := strings.Join([]string{common.GetEnv("DOKKU_ROOT"), appName}, "/")
	cmdEnv := map[string]string{
		"GIT_DIR": appRoot,
	}
	gitGcCmd := common.NewDokkuCmd("git gc --aggressive")
	gitGcCmd.Env = cmdEnv
	gitGcCmd.Execute()
}
