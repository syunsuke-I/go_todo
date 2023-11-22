package main

import (
	"fmt"
	"lesson/app/models"
)

func main() {

	user, _ := models.GetUser(2)
	user.CreateTodo("todo-2")
	todos, _ := models.GetTodos()
	for _, v := range todos {
		fmt.Println(v)
	}
}
