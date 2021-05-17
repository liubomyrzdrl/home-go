package main

import (
	"github.com/liubomyrzdrl/home-go/config"
	"github.com/liubomyrzdrl/home-go/models"
	"github.com/liubomyrzdrl/home-go/pkg/routes"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(postgres.Open(config.DSN), &gorm.Config{})
	config.DB = db
	if err != nil {
		panic("failed to connect database")
	}
	config.DB.AutoMigrate(&models.User{})

	r := routes.InitRouter()
	r.Run()
}
