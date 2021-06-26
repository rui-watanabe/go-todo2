package main

import (
	"fmt"
	"log"
	"todo/config"
)

func main() {
	fmt.Println(config.Config.Port)
	fmt.Println(config.Config.LogFile)
	fmt.Println(config.Config.SQLDriver)
	fmt.Println(config.Config.DbName)

	log.Println("test")
}
