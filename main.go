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

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", ToddFunc)
	r.HandleFunc("/contact", ContactFunc)
	r.NotFoundHandler = http.HandlerFunc(NotFoundFunc)
	//Static File Serve
	/*r.
	PathPrefix("/static/").
	Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	*/
	log.Fatal(http.ListenAndServe(":8080", r))
}

func NotFoundFunc(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/html")
	writer.WriteHeader(http.StatusNotFound)
	fmt.Fprint(writer, "<h1>We Couldn't Found Page</h1>")
}
func ContactFunc(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/html")
	fmt.Fprint(writer, "<div>To get in touch, please send an email to "+
		"<a href=\"mailto:vijay@nexthoughts.com\">vijay@nexthoughts.com</a></div>")
}

func ToddFunc(writer http.ResponseWriter, request *http.Request) {
	tmpl, err := template.New("name").Parse(doc)

	if err == nil {
		context := Context{"todd", "more go, please"}
		log.Fatal(tmpl.Execute(writer, context))
	}
}
