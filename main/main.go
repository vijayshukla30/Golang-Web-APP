package main

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
