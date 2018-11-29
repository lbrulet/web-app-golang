package main

import (
	"log"
	"net/http"
	route "web-app/routes"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/", route.YourHandler)
	r.HandleFunc("/users", route.AllUsers).Methods("GET")
	r.HandleFunc("/users", route.CreateUser).Methods("POST")
	r.HandleFunc("/users", route.UpdateUser).Methods("PUT")
	r.HandleFunc("/users", route.DeleteUser).Methods("DELETE")
	r.HandleFunc("/users/{id}", route.FindUser).Methods("GET")
	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", r))
}
