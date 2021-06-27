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
	// fmt.Println(models.Db)

	// u := &models.User{}
	// u.Name = "test"
	// u.Email = "test@gmail.com"
	// u.Password = "testtest"

	// fmt.Println(u)

	// u.CreateUser()
	u, _ := models.GetUser(1)

	fmt.Println(u)

	u.Name = "test2"
	u.Email = "test2@gmail.com"
	u.UpdateUser()
	u, _ = models.GetUser(1)
	fmt.Println(u)
}
