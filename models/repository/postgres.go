package repository

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host string
	Port string
	User string
	Password string
	DBName string
	SSLMode string
}

func   NewPostgresDB (config Config ) (*gorm.DB, error ){
    configQuery := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
    	config.Host, config.Port, config.User, config.DBName, config.SSLMode)
	db, err := gorm.Open(postgres.Open(configQuery), &gorm.Config{})
	if err !=nil {
		panic("failed to connect database")
	}
	return  db, nil
}
