package main

import (
	"fmt"
	"lesson/app/models"
)

func main() {

	fmt.Println(models.Db)
	u := &models.User{}
	u.Name = "test"
	u.Email = "test@test"
	u.Password = "testpassword"

	u.CreateUser()

	user, _ := models.GetUser(2)
	user.CreateTodo("first Todo")
}