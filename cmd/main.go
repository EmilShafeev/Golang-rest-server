package main

import (
	"Den_task1/package/handler"
	"Den_task1/package/repository"
	"Den_task1/package/repository/repository_disc"
	"Den_task1/package/service"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
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

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatalf("could not start server: %v\n", err)
		}
	}()
	log.Print("Starting started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Print("Server got signal for stop")

	if err := server.Shutdown(context.Background()); err != nil {
		log.Fatalf("Oh shit, we have not any chances for stop here. Error: %s", err)
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("disc")
	return viper.ReadInConfig()
}
