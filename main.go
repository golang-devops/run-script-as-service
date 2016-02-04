package main

import (
	"flag"
	"github.com/zero-boilerplate/go-api-helpers/service"
	"log"
)

var (
	serviceName = flag.String("name", "", "Name of this script-wrapper service")
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Fatalf("Service ERROR: %s", getStringFromRecovery(r))
		}
	}()

	a := &app{}

	args := []string{}
	if len(*service.ServiceFlag) > 0 {
		args = flag.Args()

		if *serviceName == "" {
			panic("The service name is required")
		}

		if *service.ServiceFlag == "install" {
			if len(args) == 0 {
				panic("The list of arguments cannot be empty")
			}
		}
	}

	combinedArgs := []string{
		"-name",
		*serviceName,
	}
	combinedArgs = append(combinedArgs, args...)

	service.NewServiceRunnerBuilder(*serviceName, a).
		WithAdditionalArguments(combinedArgs...).
		WithOnStopHandler(a).
		WithServiceUserName_AsCurrentUser().
		Run()
}
