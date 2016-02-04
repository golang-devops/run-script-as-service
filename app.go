package main

import (
	"github.com/ayufan/golang-kardianos-service"
	"os"
	"os/exec"
)

type app struct {
	logger             service.Logger
	watcherDoneChannel chan bool
}

func (a *app) OnStop() {
	defer recover()
	if a.watcherDoneChannel != nil {
		close(a.watcherDoneChannel)
	}
}

func (a *app) runScript(exe string, args ...string) {
	cmd := exec.Command(exe, args...)
	if output, err := cmd.CombinedOutput(); err != nil {
		a.logger.Infof("Script (%s, %#v) failed to run. OUTPUT: %s. ERROR: %s", exe, args, cleanOutput(string(output)), err.Error())
	} else {
		a.logger.Infof("Script ran successfully with output: %s", cleanOutput(string(output)))
	}
}

func (a *app) Run(logger service.Logger) {
	a.logger = logger
	defer func() {
		if r := recover(); r != nil {
			a.logger.Errorf("Run app error: %s", getStringFromRecovery(r))
		}
	}()

	//At this point we expect arguments to be:
	//  0 - This EXE path
	//  1 - -name
	//  2 - SERVICE NAME
	//  3 - Starting the actual script exe + args
	//  4 - ...
	scriptExeAndArgs := os.Args[3:]

	scriptExe := scriptExeAndArgs[0]
	scriptArgs := []string{}
	if len(scriptExeAndArgs) > 1 {
		scriptArgs = scriptExeAndArgs[1:]
	}
	a.runScript(scriptExe, scriptArgs...)
}
