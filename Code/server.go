package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func home(writer http.ResponseWriter, req *http.Request) {
	fmt.Println("Home request handled")
	writer.WriteHeader(http.StatusOK)
}

type Todo struct {
	Title   string
	Content string
}

type TodosPage struct {
	PageTitle string
	PageTodos []Todo
}

var todos = []Todo{Todo{Title: "Oe", Content: "First"}, {Title: "2", Content: "Second"}}

func handleTodosReq(writer http.ResponseWriter, req *http.Request) {
	// writer.WriteHeader(http.StatusOK)
	// fmt.Fprint(writer, "Todos")

	t, err := template.ParseFiles("todos.html")
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		log.Println("Template parsing error", err)
	}

	err = t.Execute(writer, TodosPage{PageTitle: "Todos title", PageTodos: todos})
}

func addTodo(writer http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		log.Println("Reuest todos parsing error", err)
	}

	var todo Todo = Todo{}

	todo.Title = req.Form.Get("title")
	todo.Content = req.Form.Get("content")

	todos = append(todos, todo)

	// BAD PRACTICE. Only for quick learnign purposes
	http.Redirect(writer, req, "/todos", http.StatusSeeOther)
}

func main() {
	fmt.Println("Starting web server on 8080")
	http.HandleFunc("/", home)
	http.HandleFunc("/todos/", handleTodosReq)
	http.HandleFunc("/add-todo/", addTodo)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
