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

func Home(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/html")
	tmpl, err := template.New("name").Parse(doc)
	if err == nil {
		context := Context{"todd", "more go, please"}
		tmpl.Execute(writer, context)
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", Home)
	r.HandleFunc("/contact", Contact)
	r.NotFoundHandler = http.HandlerFunc(NotFound)
	log.Fatal(http.ListenAndServe(":8080", r))
}

func Contact(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "<h1>Contact Us</h1>")
}
func NotFound(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusNotFound)
	fmt.Fprint(writer, "<h1>Page Not Exist</h1>")
}
