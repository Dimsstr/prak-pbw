package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// model data
type Todo struct {
	ID        string `json:"id"`
	TitleID   string `json:"title_id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

// initialize slice to store data
var todos []Todo

func main() {
	// initialize router
	router := mux.NewRouter()

	// add handlers for routes "/todos"
	router.HandleFunc("/todos", GetTodos).Methods("GET")
	router.HandleFunc("/todos", CreateTodo).Methods("POST")

	// start server on port 8000
	log.Fatal(http.ListenAndServe(":8000", router))
}

// handler to get all todos
func GetTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

// handler to create a new todo
func CreateTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newTodo Todo
	json.NewDecoder(r.Body).Decode(&newTodo)
	newTodo.ID = uuid.New().String()
	todos = append(todos, newTodo)
	json.NewEncoder(w).Encode(newTodo)
}
