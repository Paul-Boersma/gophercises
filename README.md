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

!TODO: To make this exercise more challenging I decided to later improve it by:
- Allowing the use of a file instead of raw bytes data
- Build a JSON handler next to the YAML handler
- Implement a connection with a database to get the mapping

#### Choose Your Own Adventure


The following packages are covered:
- text/template
- html/template

It includes the concept of building HTML pages for Go Clients throught actions '{{ . }}'




##### Notes:
Templates are defined as .gohtml files.

Process:
- Create template (template.New) 
- Parse template (template.Parse)
- Execute template (template.Execute)


Terminology
Actions:
- Data evaluations
    - Arguments (simply go values: nil, primitives, ., vars, fields, niladic methods, niladic functions)
    - Pipelines (chained seqeuence of commands; arguments including functions & methods) 
        arguments and pipelines may be assigned to variables '$variable = pipeline'
- Control structures


Functions are declared in:
- template struct (go code)
- global function map (functions & binary comparison operators)
        template functions may be created through code 'template.Funcs()'

Combining templates:
- Associated templates
- Nested templates
        assocate templates through code 'template.New()', all associated templates are stored in a single namespace within the Template struct
        create nested template at top of file through {{ define "templateName" }} <template> {{end}}
        refer to template -> {{ template "templateName" }} 



