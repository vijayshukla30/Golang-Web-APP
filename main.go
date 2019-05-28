package main

import (
	"github.com/gorilla/mux"
	controllers "github.com/vijayshukla30/web_app/controller"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	//Static File Serve
	r.
		PathPrefix("/static/").
		Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	//Controllers
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/login", controllers.Login)
	r.HandleFunc("/contact", controllers.Contact)

	//Create a New User
	r.HandleFunc("/register", controllers.Register).Methods("GET")
	r.HandleFunc("/save", controllers.Save).Methods("POST")
	//Update User
	r.HandleFunc("/edit/{id}", controllers.Edit).Methods("GET")
	r.HandleFunc("/update/{id}", controllers.Update).Methods("PUT")
	//Delete User
	r.HandleFunc("/delete/{id}", controllers.Delete).Methods("DELETE")

	r.NotFoundHandler = http.HandlerFunc(controllers.NotFound)

	log.Fatal(http.ListenAndServe(":8080", r))
}
