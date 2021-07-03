package main

import (
	"fmt"
	"todo/app/controllers"
	"todo/app/models"
)

// "fmt"
// "log"
// "todo/config"

func main() {
	fmt.Println(models.Db)
	controllers.StartMainServer()
}
