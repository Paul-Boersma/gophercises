### Gophercises
This repo contains the exercises I made to gain familiarity with the Go Language.

#### Quiz-game
This game asks the user to answer simple mathematical questions of a provided file, 
within a specified time limit. It covers (a.o.) the following packages:
- time
- flag
- encoding/csv

It also introduces the concept of go routines and channels.

#### URL-shortner
This shortner package handles redirects of short urls to more complicated ones.
2 Types of Handlers are used as a ServeMux; a map and a yaml file. 

During this exercise I mostly gained familiarity with the net/http package and
the Types associated with this package, in particular:
- ServeMux, Server, Client
- Response, Request
- ResponseWriter (interface)

It covers the following packages:
- net
- net/http
- gopkg.in/yaml.v3 (3rd party)

It also introduces the concept of go workspaces, go modules and package management. 

#### Choose Your Own Adventure



