package main

import (
	"GinIntroduction/internal/config"
	"GinIntroduction/internal/entity"
	"GinIntroduction/internal/server"
	"github.com/spf13/viper"
	"log"
)

func init() {
	config.InitEnv()
	config.Connection()

	config.DB.AutoMigrate(&entity.Product{})
}

func main() {
	r := server.NewRouter()
	if err := r.Run(viper.GetString("PORT")); err != nil {
		log.Fatal(err)
	}
}
