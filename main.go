package main

import (
	"fmt"
	"lesson/app/models"
)

func main() {

	t, _ := models.GetTodo(1)
	t.Content = "Updated"
	t.UpdateTodo()
	fmt.Println(t)

}
