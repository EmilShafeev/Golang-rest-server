package main

import (
	"Den_task1/package/handler"
	"Den_task1/package/repository"
	"Den_task1/package/repository/repository_disc"
	"Den_task1/package/service"
	"log"
	"net/http"
	"time"

	"github.com/spf13/viper"
)

func main() {

	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	repository := repository.NewDiscRepository(repository_disc.DiscSettings{SendAllwaysError: false})
	service := service.NewService(repository)
	handler := handler.NewHandler(service)

	server := &http.Server{
		Addr:         viper.GetString("port"),
		Handler:      handler,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	log.Printf("Starting server on %s", viper.GetString("port"))
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("could not start server: %v\n", err)
	}

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("disc")
	return viper.ReadInConfig()
}
