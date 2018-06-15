package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/receive", receiveIndex)
	router.HandleFunc("/receive/{data}", receiveShow)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome! This is the landing page")
}

func receiveIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This could be the overview of received data!")
}

func receiveShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	data := vars["data"]
	fmt.Fprintln(w, "This is how I can accept data:", data)
}

// take a look at https://github.com/ET-CS/golang-response-examples/blob/master/ajax-json.go
// plan of action: build template html with content div.
// poll for update via ajax call to rest api
// todo evaluate if all content or only new content should be sent
// if all content : delete content of div, put in result of ajax call
// if new content : append new content to div, check for duplicates?