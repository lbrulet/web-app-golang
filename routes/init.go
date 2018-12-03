package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/lbrulet/web-app-golang/routes/user"
)

func init() {
	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/", user.YourHandler)
	r.HandleFunc("/users", user.AllUsers).Methods("GET")
	r.HandleFunc("/users", user.CreateUser).Methods("POST")
	r.HandleFunc("/users", user.UpdateUser).Methods("PUT")
	r.HandleFunc("/users", user.DeleteUser).Methods("DELETE")
	r.HandleFunc("/users/{id}", user.FindUser).Methods("GET")
	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", r))
}
