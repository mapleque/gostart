package main

import "github.com/mapleque/gostart/http/todolist"

func main() {
	todolist.RunTodoListService(":80")
}
