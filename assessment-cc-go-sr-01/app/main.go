package main

import (
	"battle-of-monsters/app/config"
	"battle-of-monsters/app/db"
	"battle-of-monsters/app/router"
	"fmt"
	"log"
)

func main() {
	config.Load()
	db.Connect()

	server := router.Router()
	port := fmt.Sprintf(":%s", config.ENV.Port)

	if err := server.Run(port); err != nil {
		log.Fatalln("error when server is initializing")
	}
}
