package main

import (
	"fmt"
	"lesson/app/models"
)

func main() {
	// fmt.Println(models.Db)
	// u := &models.User{}
	// u.Name = "test"
	// u.Email = "test@test"
	// u.Password = "testpassword"
	// fmt.Println(u)

	// u.CreateUser()

	u, _ := models.GetUser(1)
	fmt.Println(u)

	u.Name = "Test2"
	u.Email = "test2@test2"
	u.UpdateUser()

	u, _ = models.GetUser(1)
	fmt.Println(u)
}
