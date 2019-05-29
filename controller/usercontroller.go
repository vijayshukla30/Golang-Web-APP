package controller

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/vijayshukla30/web_app/config"
	"github.com/vijayshukla30/web_app/model"
	"html/template"
	"log"
	"net/http"
)

func Delete(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]
	db := config.DbConn()
	delForm, err := db.Prepare("DELETE FROM user WHERE id=?")
	if err != nil {
		panic(err.Error())
	}

	delForm.Exec(id)
	log.Println("Deleted")
	defer db.Close()
	http.Redirect(writer, request, "/", 301)
}
func Update(writer http.ResponseWriter, request *http.Request) {
	db := config.DbConn()
	if request.Method == "POST" {
		id := request.FormValue("uid")
		firstName := request.FormValue("firstName")
		lastName := request.FormValue("lastName")
		email := request.FormValue("email")

		insForm, err := db.Prepare("UPDATE  user SET first_name=?, last_name=?, email=? WHERE id=?")
		if err != nil {
			panic(err.Error())
		}

		insForm.Exec(firstName, lastName, email, id)
		log.Println("Updated")
	}
	defer db.Close()
	http.Redirect(writer, request, "/", 301)
}
func Edit(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	nid := vars["id"]
	fmt.Println(nid)
	db := config.DbConn()
	sel, err := db.Query("SELECT * FROM user WHERE id=?", nid)

	if err != nil {
		panic(err.Error())
	}

	user := model.User{}

	for sel.Next() {
		var id int
		var firstName, lastName, email string

		err = sel.Scan(&id, &firstName, &lastName, &email)

		if err != nil {
			panic(err.Error())
		}

		user.Id = id
		user.FirstName = firstName
		user.LastName = lastName
		user.Email = email
	}
	tmpl, err := template.New("").ParseFiles("templates/edit.html", "templates/base.html")
	defer db.Close()

	if err == nil {
		tmpl.ExecuteTemplate(writer, "base", user)
	} else {
		fmt.Fprint(writer, err)
	}
}

func Show(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	nid := vars["id"]
	fmt.Println("Checking DB")
	db := config.DbConn()
	sel, err := db.Query("SELECT * FROM user WHERE id=?", nid)
	if err != nil {
		panic(err.Error())
	}
	user := model.User{}
	for sel.Next() {
		var id int
		var firstName, lastName, email string

		err = sel.Scan(&id, &firstName, &lastName, &email)

		if err != nil {
			panic(err.Error())
		}

		user.Id = id
		user.FirstName = firstName
		user.LastName = lastName
		user.Email = email
	}
	tmpl, _ := template.New("").ParseFiles("templates/view.html", "templates/base.html")
	tmpl.ExecuteTemplate(writer,"base", user)
	defer db.Close()

}

func Save(writer http.ResponseWriter, request *http.Request) {
	db := config.DbConn()
	if request.Method == "POST" {
		firstName := request.FormValue("firstName")
		lastName := request.FormValue("lastName")
		email := request.FormValue("email")

		insForm, err := db.Prepare("INSERT INTO user(first_name, last_name, email) VALUES (?,?,?)")
		if err != nil {
			panic(err.Error())
		}

		insForm.Exec(firstName, lastName, email)
		log.Println("Inserted")
	}
	defer db.Close()
	http.Redirect(writer, request, "/", 301)
}
