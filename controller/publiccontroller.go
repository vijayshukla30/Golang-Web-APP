package controller

import (
	"fmt"
	dbconn "github.com/vijayshukla30/web_app/config"
	models "github.com/vijayshukla30/web_app/model"
	"html/template"
	"net/http"
)

func Home(writer http.ResponseWriter, request *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err == nil {
		db := dbconn.DbConn()
		selDB, err := db.Query("SELECT * FROM user ORDER BY id DESC")
		if err != nil {
			panic(err.Error())
		}

		user := models.User{}

		var res []models.User

		for selDB.Next() {
			var id int
			var firstName, lastName, email string
			err = selDB.Scan(&id, &email, &firstName, &lastName)

			if err != nil {
				panic(err.Error())
			}

			//User Assignment
			user.Id = id
			user.Email = email
			user.FirstName = firstName
			user.LastName = lastName

			res = append(res, user)
		}
		fmt.Println(res)
		tmpl.Execute(writer, res)
		defer db.Close()
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
