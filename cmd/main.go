package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/dwprz/prasorganic-cart-service/src/core/restful/handler"
	"github.com/dwprz/prasorganic-cart-service/src/core/restful/middleware"
	"github.com/dwprz/prasorganic-cart-service/src/core/restful/server"
	"github.com/dwprz/prasorganic-cart-service/src/infrastructure/database"
	"github.com/dwprz/prasorganic-cart-service/src/repository"
	"github.com/dwprz/prasorganic-cart-service/src/service"
	"github.com/go-playground/validator/v10"
)

func handleCloseApp(closeCH chan struct{}) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		close(closeCH)
	}()
}

func main() {
	closeCH := make(chan struct{})
	handleCloseApp(closeCH)

	validate := validator.New()
	postgresDB := database.NewPostgres()

	cartRepository := repository.NewCart(postgresDB)
	cartService := service.NewCart(cartRepository, validate)
	
	cartRestfulHandler := handler.NewCart(cartService)
	middleware := middleware.New()

	restfulServer := server.New(cartRestfulHandler, middleware)
	defer restfulServer.Stop()

	go restfulServer.Run()

	<-closeCH
}
