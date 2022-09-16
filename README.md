## Gophercises
This repo contains the exercises I made to gain familiarity with the Go Language.

### Quiz-game
This game asks the user to answer simple mathematical questions of a provided file, 
within a specified time limit. It covers (a.o.) the following packages:
- time
- flag
- encoding/csv

It also included the concept of go routines and channels, 
required for implementing the timer functionality.

### URL-shortner
This url-shortner package handles requests to urls for which no handlers have been implemented.
Instead requests to these 'short' urls are redirected to urls which can be handled, 
albeit to 3rd party servers.

The package provides two mapping functions, a regular map and yaml format (not exactly ServeMuxes as no handlers are registered) 

<b>What did I learn doing this exercise?</b>
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

### Choose Your Own Adventure
<b>What did I learn doing this exercise?</b>
- Opening files from the operating system --> package: os
- Parsing JSON data into Go Structures --> package: encode/json
- Building a (HTML) template file (.gohtml) --> package: text/template & html/template
- Setting up a HTTP server and handling requests by writing the template file --> package: net/http
- Handle the requests through a handler struct, which takes optional parameters
like story and template to deviate from the default one.

The following design pattern is covered:
- Functional options pattern

### Link-Parser
A link parser may come in handy when trying to scrape data from websites.

<b>What did I learn doing this exercise?</b>
- Creation of packages, also made me consider folder structure and http webservice architecture
- Writing tests --> package: testing
- Getting more comforatble with reading package documentation
- Getting familiar with DFS and BFS algorithms --> leetcode.com

### Sitemap Builder


<b>What did I learn doing this exercise?</b>






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



