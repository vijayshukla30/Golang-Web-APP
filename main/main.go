package main

/*
-http.ListenAndServe
--Listens for, and responds to, http requests
--Handles each request using go routines
---Lightweight concurrency
---This is multiplexing
-http.Handle
--Handle a URL request
--Maps a URL to any Type (“object”) implementing the handler interface
-http.HandleFunc
--Handles a URL request
--Maps a URL to a func
---“Wrapper” around a function
---Turns any function to a handler
*/

/*
SERVER-MUX
compares incoming requests against a list of predefined URL paths,
amd ca;;s tje associated handles for the path whenever a match found.

HANDLERS
responsible for writing response headers and bodies.
Almost any type ("Object") can be a handles, so long as it satisfies the http.Handler interface.
*/
import (
	"net/http"
	"text/template"
)

func main() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":8080", nil)
}

func home(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("Content-Type", "text/html")
	tmpl, err := template.New("home").Parse(doc)

	if err == nil {
		tmpl.Execute(writer, nil)
	}
}

const doc = `
<html>
<head>
<title>
First Template
</title>
</head>
<body>
<h4>First Template Sample</h4>
</body>
</html
`

//http://localhost:8080/home/vijay/Downloads/coderthemes.com/minton/material/index.html
