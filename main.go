package main

import (
	"fmt"
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
	http.HandleFunc("/", toddFunc)
	http.ListenAndServe(":8080", nil)
}