package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/dwprz/prasorganic-cart-service/src/core/grpc/client"
	"github.com/dwprz/prasorganic-cart-service/src/core/grpc/delivery"
	"github.com/dwprz/prasorganic-cart-service/src/core/grpc/interceptor"
	"github.com/dwprz/prasorganic-cart-service/src/core/restful/handler"
	"github.com/dwprz/prasorganic-cart-service/src/core/restful/middleware"
	"github.com/dwprz/prasorganic-cart-service/src/core/restful/server"
	"github.com/dwprz/prasorganic-cart-service/src/infrastructure/database"
	"github.com/dwprz/prasorganic-cart-service/src/repository"
	"github.com/dwprz/prasorganic-cart-service/src/service"
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

	postgresDB := database.NewPostgres()

	unaryRequestInterceptor := interceptor.NewUnaryRequest()
	productGrpcDelivery, productGrpcConn := delivery.NewProductGrpc(unaryRequestInterceptor)

	cartRepository := repository.NewCart(postgresDB)
	grpcClient := client.NewGrpc(productGrpcDelivery, productGrpcConn)
	defer grpcClient.Close()

	cartService := service.NewCart(cartRepository, grpcClient)

	cartRestfulHandler := handler.NewCartRESTful(cartService)
	middleware := middleware.New()

	restfulServer := server.New(cartRestfulHandler, middleware)
	defer restfulServer.Stop()

	go restfulServer.Run()

	<-closeCH
}
