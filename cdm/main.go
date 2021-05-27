package main

import (
	homego "github.com/liubomyrzdrl/home-go"
	"github.com/liubomyrzdrl/home-go/config"
	"github.com/liubomyrzdrl/home-go/models"
	"github.com/liubomyrzdrl/home-go/models/repository"
	"github.com/liubomyrzdrl/home-go/pkg/handler"
	"github.com/liubomyrzdrl/home-go/pkg/service"
	"github.com/spf13/viper"
	"log"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatal("error init configs: %s", err.Error())
	}

	db,err := repository.NewPostgresDB(repository.Config{
		Host:  viper.GetString("db.host"),
		User: viper.GetString("db.user"),
		DBName: viper.GetString("db.dbName"),
		Port: viper.GetString("db.port"),
	})
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}

        repos := repository.NewRepository(db)
		services := service.NewService(repos)
		handler := handler.NewHandler(services)


	db.AutoMigrate(&models.User{}, &models.CreditCard{}, &models.Role{})

	srv := new(homego.Server)
	if err := srv.Run(config.PORT, handler.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server %s", err.Error())
	}

	//TODO run server from Gin
	//r := routes.InitRouter()
	//r.Run()
}

func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}