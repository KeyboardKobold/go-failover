package main

import (
	"fmt"
	"log"
	"net/http"
	"html/template"

	"github.com/gorilla/mux"
	"strconv"
	"os"
)

type Event struct {
	EventId int
}

var Events []Event

func main() {

	Events = []Event{
		{EventId: 1},
		{EventId: 2},
		{EventId: 3},
	}

	Events = append(Events, Event{4})

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/getEvents", getEvents)
	router.HandleFunc("/receive/{data}", receiveShow)
	router.HandleFunc("/addEvent", addEvent).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("consumer/main/templates/Index.html"))
	tmpl.Execute(w, Events)
}

func getEvents(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("consumer/main/templates/eventList.html"))
	tmpl.Execute(w, Events)
}

func receiveShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	data := vars["data"]
	fmt.Fprintln(w, "This is how I can accept data:", data)
}

func addEvent(w http.ResponseWriter, r *http.Request) {
	eventIdText := r.FormValue("eventId")

	//parsing string to int
	eventId, err := strconv.Atoi(eventIdText)
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}
	Events = append(Events, Event{eventId})
}

// take a look at https://github.com/ET-CS/golang-response-examples/blob/master/ajax-json.go
// plan of action: build template html with content div.
// poll for update via ajax call to rest api
// todo evaluate if all content or only new content should be sent
// if all content : delete content of div, put in result of ajax call
// if new content : append new content to div, check for duplicates?
// take a look at https://stackoverflow.com/questions/8567114/how-to-make-an-ajax-call-without-jquery