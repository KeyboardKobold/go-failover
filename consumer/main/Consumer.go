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