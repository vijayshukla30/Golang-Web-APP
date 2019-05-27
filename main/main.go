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
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type MyHandler struct {
}

func (this *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]

	log.Println(path)
	data, err := ioutil.ReadFile("/" + string(path))

	if (err == nil) {
		var contentType string

		if strings.HasSuffix(path, ".css") {
			contentType = "text/css"
		} else if strings.HasSuffix(path, ".html") {
			contentType = "text/html"
		} else if strings.HasSuffix(path, ".js") {
			contentType = "text/javascript"
		} else if strings.HasSuffix(path, ".png") {
			contentType = "text/png"
		} else if strings.HasSuffix(path, ".jpg") {
			contentType = "text/jpeg"
		} else if strings.HasSuffix(path, ".mp4") {
			contentType = "video/mp4"
		} else {
			contentType = "text/plain"
		}
		w.Header().Add("Content-Type", contentType)
		w.Write(data)
	} else {
		w.WriteHeader(404)
		w.Write([]byte("404 - " + http.StatusText(404)))
	}
}

func main() {
	http.Handle("/", new(MyHandler))
	http.ListenAndServe(":8080", nil)
}

/*func home(writer http.ResponseWriter, request *http.Request) {

	writer.Write([]byte("Hello Universe"))
}*/

//http://localhost:8080/home/vijay/Downloads/coderthemes.com/minton/material/index.html
