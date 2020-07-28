# GO Windows Service

This is based on the [GO Windows service example](https://godoc.org/golang.org/x/sys/windows/svc/example) provided by the GO Project. 
It is a project shell to create a Windows service.

## Getting Started

The program tested on GO 1.14.2.  The generated executable accepts a single parameter.  The parameter values include:

* debug - runs the program in the interactive mode from the command-line
* install - installs a windows service
* remove - remove the windows service
* start
* stop

## Customizing

The code exists in two parts

* cmd/main.go - initialize service and check params
* service.go - windows service wrapper

* hello.go - simple http hello server

For you service need change package name in the service.go
and change server initialization code in the Execute function.

* `svcName` - This constant is the name of the installed service.  This is used for NET START and NET STOP commands.
* `svcNameLong` - This is the longer service name that appears in the Services control panel.

