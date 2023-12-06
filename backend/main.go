package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// structure for a task - title, description
type (
	todo struct {
		Title string `json:"title"`
		Desc  string `json:"description"`
	}
)

// array of tasks as a todo list - memory
var todoList []todo

func main() {
	// fmt.Print("running...")
	// endpoint for listing and adding
	http.HandleFunc("/", ToDoListHandler)
	// port to listen on 8080
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func ToDoListHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// check
	// fmt.Fprintf(w, "OK\n")

	switch r.Method {
	// when request is a GET operation
	case "GET":
		http.ServeFile(w, r, "App.tsx")
	// when request is a POST operation
	case "POST":
		title := r.FormValue("title")
		desc := r.FormValue("description")

		// add new task to todo list
		new := todo{title, desc}
		todoList := append(todoList, new)

		// return encoded response of todolist in .json format
		json.NewEncoder(w).Encode(todoList)
	}
}
