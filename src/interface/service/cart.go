package service

import (
	"context"

	"github.com/dwprz/prasorganic-cart-service/src/model/dto"
)

type Cart interface {
	Create(ctx context.Context, data *dto.CreateCartReq) (error) 
}