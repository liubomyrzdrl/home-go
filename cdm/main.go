package main

import (
	homego "github.com/liubomyrzdrl/home-go"
	"github.com/liubomyrzdrl/home-go/config"
	"github.com/liubomyrzdrl/home-go/models"
	"github.com/liubomyrzdrl/home-go/pkg/routes"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func main() {
	db, err := gorm.Open(postgres.Open(config.DSNREMOTE), &gorm.Config{})
	config.DB = db
	if err != nil {
		panic("failed to connect database")
	}
	config.DB.AutoMigrate(&models.User{}, &models.CreditCard{}, &models.Role{})

	srv := new(homego.Server)
	if err := srv.Run(config.PORT, routes.InitRouter()); err != nil {
		log.Fatalf("error occured while running http server %s", err.Error())
	}

	//TODO run server from Gin
	//r := routes.InitRouter()
	//r.Run()
}
