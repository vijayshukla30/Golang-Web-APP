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
import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", home)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func home(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Hello Universe"))
}
