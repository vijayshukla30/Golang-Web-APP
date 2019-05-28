package controller

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/vijayshukla30/web_app/config"
	"github.com/vijayshukla30/web_app/model"
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
	vars := mux.Vars(request)
	id := vars["id"]

	fmt.Fprint(writer, "<p>You are going to Update "+id+"</p>")
}
func Edit(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]

	fmt.Fprint(writer, "<p>You are going to Edit "+id+"</p>")
}
func Save(writer http.ResponseWriter, request *http.Request) {
	db := config.DbConn()
	if request.Method == "POST" {
		firstName := request.FormValue("firstName")
		lastName := request.FormValue("lastName")
		email := request.FormValue("email")
		details := model.User{
			Email:     email,
			FirstName: firstName,
			LastName:  lastName,
		}

		fmt.Println(details)
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
