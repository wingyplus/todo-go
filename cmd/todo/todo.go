package main

import "errors"

type Todo struct {
	Text      string `json:"text"`
	Completed bool   `json:"completed"`
}

type InMemoryTodoStorage struct {
	todos []Todo
}

func (storage *InMemoryTodoStorage) listTodos() []Todo {
	return storage.todos
}

func (storage *InMemoryTodoStorage) createTodo(todo Todo) {
	storage.todos = append(storage.todos, todo)
}

func (storage *InMemoryTodoStorage) updateTodo(id int, todo Todo) error {
	if id > len(storage.todos) {
		return errors.New("not found")
	}

	existingTodo := storage.todos[id-1]
	if todo.Text != "" {
		existingTodo.Text = todo.Text
	}
	existingTodo.Completed = todo.Completed
	storage.todos[id-1] = existingTodo
	return nil
}
