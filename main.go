package main

import (
	"fmt"
	"lesson/app/models"
)

func main() {

	todo, _ := models.GetTodo(1)
	fmt.Println(todo)
}
