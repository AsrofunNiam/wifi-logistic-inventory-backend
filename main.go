package main

import (
	"log"
	"net/http"

	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/app"
	c "github.com/AsrofunNiam/wifi-logistic-inventory-backend/configuration"
	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/helper"
	"github.com/go-playground/validator/v10"
)

func main() {
	configuration, err := c.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	port := configuration.Port
	db := app.ConnectDatabase(configuration.User, configuration.Host, configuration.Password, configuration.PortDB, configuration.Db)
	redisClient := app.ConnectClientCRedis(configuration.RedisHost, configuration.RedisPort, configuration.RedisPassword)

	validate := validator.New()
	router := app.NewRouter(db, redisClient, validate)
	server := http.Server{
		Addr:    ":" + port,
		Handler: router,
	}
	log.Printf("Server is running on port %s", port)

	err = server.ListenAndServe()
	helper.PanicIfError(err)
}
