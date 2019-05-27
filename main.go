package main

import (
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
	writer.Header().Add("Content-Type", "text/html")
	tmpl, err := template.New("name").Parse(doc)

	if err == nil {
		context := Context{"todd", "more go, please"}
		tmpl.Execute(writer, context)
	}
}

func main() {
	http.HandleFunc("/", toddFunc)
	//http.HandleFunc("/todd", toddFunc)
	http.ListenAndServe(":8080", nil)
}

//http://localhost:8080/home/vijay/Downloads/coderthemes.com/minton/material/index.html
