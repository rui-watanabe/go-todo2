package main

import (
	"fmt"
	"todo/app/models"
)

// "fmt"
// "log"
// "todo/config"

func main() {
	// fmt.Println(config.Config.Port)
	// fmt.Println(config.Config.LogFile)
	// fmt.Println(config.Config.SQLDriver)
	// fmt.Println(config.Config.DbName)

	// log.Println("test")
	fmt.Println(models.Db)

	// u := &models.User{}
	// u.Name = "test2"
	// u.Email = "test2@gmail.com"
	// u.Password = "testtest"
	// fmt.Println(u)
	// u.CreateUser()

	// u.CreateUser()
	// user, _ := models.GetUser(2)
	// user.CreateTodo("first todo")
	// fmt.Println(user)

	// u.Name = "test"
	// u.Email = "test@gmail.com"
	// u.UpdateUser()
	// u, _ = models.GetUser(1)
	// fmt.Println(u)

	// u.DeleteUser()
	// u, _ = models.GetUser(1)
	// fmt.Println(u)

	// t, _ := models.GetTodo(1)
	// fmt.Println(t)

	// user, _ := models.GetUser(3)
	// user.CreateTodo("third todo")

	// todos, _ := models.GetTodos()

	// for _, v := range todos {
	// 	fmt.Println(v)
	// }

	// user3, _ := models.GetUser(3)
	// todos, _ := user3.GetTodosByUser()

	// for _, v := range todos {
	// 	fmt.Println(v)
	// }

	t, _ := models.GetTodo(4)
	// t.Content = "fourth todo"
	t.DeleteTodo()
}
