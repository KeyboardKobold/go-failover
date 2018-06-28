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
	router.HandleFunc("/health", checkHealth)

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

func checkHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func contains(s []Event, e int) bool {
	for _, a := range s {
		if a.EventId == e {
			return true
		}
	}
	return false
}
