package repository

import (
	"context"

	"github.com/dwprz/prasorganic-cart-service/src/model/dto"
)

type Cart interface {
	Create(ctx context.Context, data *dto.CreateCartReq) error
	FindManyByUserId(ctx context.Context, userId string, limit, offset int) (*dto.CartWithCountRes, error)
	CountByUserId(ctx context.Context, userId string) (totalCart int64, err error)
	DeleteItem(ctx context.Context, data *dto.DeleteItemCartReq) error
}
