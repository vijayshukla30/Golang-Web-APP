package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"text/template"
)

type Context struct {
	FirstName string
	Message   string
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
<p>Hello {{.FirstName}}</p>
<p>My Message: {{.Message}}</p>
</body>
</html
`

func toddFunc(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/html")

	if request.URL.Path == "/" {
		tmpl, err := template.New("name").Parse(doc)

		if err == nil {
			context := Context{"todd", "more go, please"}
			tmpl.Execute(writer, context)
		}
	} else {
		writer.WriteHeader(http.StatusNotFound)
		fmt.Fprint(writer, "<h1>We Couldn't Found Page</h1>")
	}
}

func main() {
	/*
		PROCESS - 1
		http.HandleFunc("/", toddFunc)
		http.ListenAndServe(":8080", nil)*/

	/*
		PROCESS - 2
	*/
	/*mux := &http.ServeMux{}
	mux.HandleFunc("/", toddFunc)
	log.Fatal(http.ListenAndServe(":8080", mux))*/

	r := mux.NewRouter()
	r.HandleFunc("/", toddFunc)
	log.Fatal(http.ListenAndServe(":8080", r))
}
