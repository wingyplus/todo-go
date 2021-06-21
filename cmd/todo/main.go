// GET
// POST
// PUT
// DELETE
// OPTIONS

// Features:
// 	* List todos
//		* GET /todos
//  * Create new todo
//		* POST /todos
//  * Update todo
//		* PUT /todos?id=1

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

type Storage interface {
	listTodos() []Todo
	createTodo(todo Todo)
	updateTodo(id int, todo Todo) error
}

func main() {
	var storage Storage = &InMemoryTodoStorage{
		todos: make([]Todo, 0),
	}

	http.HandleFunc("/todos", func(writer http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case "GET":
			todos := storage.listTodos()
			b, _ := json.Marshal(todos)
			writer.Header().Add("Content-Type", "application/json;charset=utf-8")
			writer.Write(b)
		case "POST":
			// read json
			b, _ := io.ReadAll(request.Body)
			defer request.Body.Close()
			// serialize json data to struct
			var todo Todo
			if err := json.Unmarshal(b, &todo); err != nil {
				http.Error(writer, fmt.Sprint("Bad Request: ", err), http.StatusBadRequest)
				return
			}
			// store it
			storage.createTodo(todo)
			// response
			writer.WriteHeader(http.StatusCreated)
		case "PUT":
			// read json
			b, _ := io.ReadAll(request.Body)
			defer request.Body.Close()
			// serialize json data to struct
			var todo Todo
			if err := json.Unmarshal(b, &todo); err != nil {
				http.Error(writer, fmt.Sprint("Bad Request: ", err), http.StatusBadRequest)
				return
			}
			// update todo by using id from query
			idquery := request.URL.Query().Get("id")
			id, _ := strconv.Atoi(idquery)
			if err := storage.updateTodo(id, todo); err != nil {
				http.Error(writer, err.Error(), http.StatusBadRequest)
				return
			}
			// response
			writer.WriteHeader(http.StatusOK)
		}
	})
	if err := http.ListenAndServe(":4000", nil); err != nil {
		log.Fatal(err)
	}
}
