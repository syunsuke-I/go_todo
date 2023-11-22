package main

import (
	"lesson/app/models"
)

func main() {

	t, _ := models.GetTodo(3)
	t.DeleteTodo()
}
