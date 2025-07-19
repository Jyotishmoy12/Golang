package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux" // importing Gorilla Mux package
)

// main is the entry point of the program
func main() {
	fmt.Println("Welcome to mod in golang!") // prints a welcome message
	greeter() // call the greeter function
}

// greeter sets up the router and starts the HTTP server
func greeter() {
	fmt.Println("Hey there mod users")

	// Create a new router using mux
	r := mux.NewRouter()

	// Handle GET requests to "/" route with serveHome function
	r.HandleFunc("/", serveHome).Methods("GET")

	// Start the server on port 3000 and log any fatal errors
	// log.Fatal stops the program if ListenAndServe returns an error
	log.Fatal(http.ListenAndServe(":3000", r))
}

// serveHome is a handler function that writes HTML to the response
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to golang series</h1>"))
}
