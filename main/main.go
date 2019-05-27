package main

/*
-http.ListenAndServe
--Listens for, and responds to, http requests
--Handles each request using go routines
---Lightweight concurrency
---This is multiplexing
-http.Handle
--Handle a URL request
--Maps a URL to any Type (“object”) implementing the handler interface
-http.HandleFunc
--Handles a URL request
--Maps a URL to a func
---“Wrapper” around a function
---Turns any function to a handler
*/

/*
SERVER-MUX
compares incoming requests against a list of predefined URL paths,
amd ca;;s tje associated handles for the path whenever a match found.

HANDLERS
responsible for writing response headers and bodies.
Almost any type ("Object") can be a handles, so long as it satisfies the http.Handler interface.
*/
import (
	"net/http"
)

type Person struct {
	fName string
}

func (p *Person) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("First Name: " + p.fName))
}

func main() {
	//myMux := http.NewServeMux()
	//myMux.HandleFunc("/", home)
	//log.Fatal(http.ListenAndServe(":8080", myMux))

	personOne := &Person{fName: "Vijay"}

	http.ListenAndServe(":8080", personOne)
}

/*func home(writer http.ResponseWriter, request *http.Request) {

	writer.Write([]byte("Hello Universe"))
}*/
