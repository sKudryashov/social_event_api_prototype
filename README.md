# Event api prototype:

<a href="https://codebeat.co/projects/github-com-skudryashov-social_event_api_prototype"><img alt="codebeat badge" src="https://codebeat.co/badges/451abd5d-0ac6-4d56-9dd1-c1a7e966c40e" /></a>

**Usage example**

There is actually four actions: create, get all, get by event type, get by 
time range. Just type appropriate commands in your console.

Get by type: `curl -XPOST -H 'Content-Type: application/json' -d '{"eventType":"Usual"}' http://localhost:3030/readbytype`

Create: `curl -XPOST -H 'Content-Type: application/json' -d '{"eventType":"Usual","sessionStart":1476628565,"sessionEnd":1476628965,"linkClicked":"https://blog.golang.org/c-go-cgo","timestamp":12039109203,"params":{"C":"c++","D":"D++","R":"R is not a real language"}}' http://localhost:3030/add`

Read all: `curl -XGET -H 'Content-Type: application/json' http://localhost:3030/read`

Read by time range: `curl -XGET -H 'Content-Type: application/json' http://localhost:3030/readbytimerange/:start/:end`

Note that each action dataset has it's own validation rules and some fields are required. If the query didn't pass validation, panic would be triggered. Of course you may change this behaviour easily.

**How to launch it ?**
