### Gophercises
This repo contains the exercises I made to gain familiarity with the Go Language.

#### Quiz-game
This game asks the user to answer simple mathematical questions of a provided file, 
within a specified time limit. It covers (a.o.) the following packages:
- time
- flag
- encoding/csv

It also included the concept of go routines and channels, 
required for implementing the timer functionality.

#### URL-shortner
This url-shortner package handles requests to urls for which no handlers have been implemented.
Instead requests to these 'short' urls are redirected to urls which can be handled, 
albeit to 3rd party servers.

The package provides two mapping functions, a regular map and yaml format (not exactly ServeMuxes as no handlers are registered) 

During this exercise I mostly gained familiarity with the net/http package and
the Types associated with this package, in particular:
- Server, Client, ServeMux
- Response, Request
- ResponseWriter (interface)

The following packages are covered:
- net/http
- gopkg.in/yaml.v3 (3rd party)

It also included the concept of go workspaces, go modules and package management, 
required for maintaining multiple go modules on a single github page.

#### Choose Your Own Adventure



