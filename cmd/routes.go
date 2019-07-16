package teachingstrategiessample

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Routes - mappings for the endpoints
func Routes() {
	router := mux.NewRouter()

	subrouter := router.PathPrefix("/v1/api").Subrouter()

	subrouter.HandleFunc("/list", TodoListHandler)
	subrouter.HandleFunc("/list/{listId}", TodoListHandler)
	subrouter.HandleFunc("/todo/{listId}/todo", TodoItemHandler)
	subrouter.HandleFunc("/todo/{listId}/todo/{todoId}", TodoItemHandler)

	http.Handle("/", router)
}
