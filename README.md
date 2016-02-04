# run-script-as-service
Practically run anything as a service (shell, batch, python, etc)

# Important notice / disclaimer

This script is run as a service. So if your script is not an ongoing script (web server or a indefinite for-loop) the service will run it, and could in some scenarios keep running it on a continual basis.

This service-wrapper is currently for long-running processes and not for scheduled tasks (or cron jobs).


# Getting started

## Get code

Get the code which installs a `run-script-as-service` binary into your `$GOPATH/bin` dir.
`go get github.com/golang-devops/run-script-as-service`

## Install python script as a service
`run-script-as-service -service install -name "Python Svc" python "/path/to/python/script"`

## Uninstall the same service
`run-script-as-service -service uninstall -name "Python Svc"`

Note that the `-name` **must** be the same to install/uninstall the same service.