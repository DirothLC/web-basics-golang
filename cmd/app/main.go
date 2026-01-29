package main

import (
	"GinIntroduction/internal/config"
	"GinIntroduction/internal/entity"
	"GinIntroduction/internal/server"
	"log"
)

func init() {
	config.InitEnv()
	config.Connection()

	config.DB.AutoMigrate(&entity.Product{})
}

func main() {
	r := server.NewRouter()
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
