package cbreaker

import (
	"github.com/sony/gobreaker/v2"
)

type CircuitBreaker struct {
	ProductGrpc *gobreaker.CircuitBreaker[any]
}

func New() *CircuitBreaker {
	productGrpcCBreaker := setupForProductGrpc()

	return &CircuitBreaker{
		ProductGrpc: productGrpcCBreaker,
	}
}