package main

import (
	"golang-todo-list/internal/todolists"
	"golang-todo-list/internal/util"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	util.Initialize("postgres://dbuser:password@localhost/teachingstrategies-sample?sslmode=disable")

	router := mux.NewRouter()

	subrouter := router.PathPrefix("/v1/api").Subrouter()

	subrouter.HandleFunc("/lists", todolists.HandleRequests).Methods("GET")
	subrouter.HandleFunc("/lists/{listID}", todolists.HandleRequests).Methods("GET")

	http.Handle("/", router)

	http.ListenAndServe(":3000", nil)
}
