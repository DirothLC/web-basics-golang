package config

import (
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	_ "os"
)

var DB *gorm.DB

func Connection() {
	var err error
	DB, err = gorm.Open(postgres.Open(viper.GetString("DATABASE")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database:" + err.Error())
	}
}
