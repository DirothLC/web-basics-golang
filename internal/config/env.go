package config

import (
	"github.com/spf13/viper"
	"log"
)

func InitEnv() {
	viper.AddConfigPath("./internal/config")
	viper.SetConfigName("viper")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Fatal error config file: %s \n", err)
	}
	/*	err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading viper.env file")
		}
	*/

}
