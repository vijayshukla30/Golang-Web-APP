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
	r.HandleFunc("/register", controllers.Register)
	r.HandleFunc("/contact", controllers.Contact)
	r.NotFoundHandler = http.HandlerFunc(controllers.NotFound)

	log.Fatal(http.ListenAndServe(":8080", r))
}
