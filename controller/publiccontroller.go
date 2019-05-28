package controller

import (
	"fmt"
	models "github.com/vijayshukla30/web_app/model"
	"html/template"
	"net/http"
)

func Home(writer http.ResponseWriter, request *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err == nil {
		user := models.User{"Vijay", "Shukla"}
		tmpl.Execute(writer, user)
	} else {
		fmt.Fprint(writer, err)
	}
}

func Login(writer http.ResponseWriter, request *http.Request) {
	tmpl, err := template.ParseFiles("templates/login.html")
	if err == nil {
		tmpl.Execute(writer, nil)
	} else {
		fmt.Fprint(writer, err)
	}
}

func Register(writer http.ResponseWriter, request *http.Request) {
	tmpl, err := template.ParseFiles("templates/register.html")
	if err == nil {
		tmpl.Execute(writer, nil)
	} else {
		fmt.Fprint(writer, err)
	}
}

func NotFound(writer http.ResponseWriter, request *http.Request) {
	tmpl, err := template.ParseFiles("templates/404.html")
	if err == nil {
		tmpl.Execute(writer, nil)
	} else {
		fmt.Fprint(writer, err)
	}
}
func Contact(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/html")
	fmt.Fprint(writer, "<div>To get in touch, please send an email to "+
		"<a href=\"mailto:vijay@nexthoughts.com\">vijay@nexthoughts.com</a></div>")
}
