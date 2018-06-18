package main

import (
	"fmt"
	"log"
	"net/http"
	"html/template"

	"github.com/gorilla/mux"
)

type Event struct {
	EventId int
}

type PageData struct {
	PageTitle string
	Events     []Event
}

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/receive", receiveIndex)
	router.HandleFunc("/receive/{data}", receiveShow)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintln(w, "Welcome! This is the landing page")

	tmpl := template.Must(template.ParseFiles("consumer/main/templates/Index.html"))

	data := PageData{
		PageTitle: "My TODO list",
		Events: []Event{
			{EventId: 1},
			{EventId: 2},
			{EventId: 3},
		},
	}
	tmpl.Execute(w, data)
}

func receiveIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This could be the overview of received data!")

	tmpl := template.Must(template.ParseFiles("consumer/main/templates/eventList.html"))
	data := PageData{
		PageTitle: "My TODO list",
		Events: []Event{
			{EventId: 1},
			{EventId: 2},
			{EventId: 3},
		},
	}
	tmpl.Execute(w, data)
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
// take a look at https://stackoverflow.com/questions/8567114/how-to-make-an-ajax-call-without-jquery