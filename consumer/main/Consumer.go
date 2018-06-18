package main

import (
	"fmt"
	"log"
	"net/http"
	"html/template"

	"github.com/gorilla/mux"
	"strconv"
)

type Event struct {
	EventId int
}

var Events []Event

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/getEvents", getEvents)
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

func addEvent(w http.ResponseWriter, r *http.Request) {
	eventIdText := r.FormValue("eventId")

	//parsing string to int
	eventId, err := strconv.Atoi(eventIdText)
	if err != nil {
		// handle error
		fmt.Println(err)
		fmt.Println("Request body was: " + eventIdText)
		//os.Exit(2)
	}
	if !contains(Events, eventId) {
		Events = append(Events, Event{eventId})
	} else {
		fmt.Printf("Producer produced duplicate eventId!\n")
	}
}

func contains(s []Event, e int) bool {
	for _, a := range s {
		if a.EventId == e {
			return true
		}
	}
	return false
}

// take a look at https://github.com/ET-CS/golang-response-examples/blob/master/ajax-json.go
// plan of action: build template html with content div.
// poll for update via ajax call to rest api
// todo evaluate if all content or only new content should be sent
// if all content : delete content of div, put in result of ajax call
// if new content : append new content to div, check for duplicates?
// take a look at https://stackoverflow.com/questions/8567114/how-to-make-an-ajax-call-without-jquery

// todo improve templates to load only once
