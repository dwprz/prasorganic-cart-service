package repository

import (
	"context"

	"github.com/dwprz/prasorganic-cart-service/src/common/errors"
	"github.com/dwprz/prasorganic-cart-service/src/interface/repository"
	"github.com/dwprz/prasorganic-cart-service/src/model/dto"
	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
)

type CartImpl struct {
	db *gorm.DB
}

func NewCart(db *gorm.DB) repository.Cart {
	return &CartImpl{
		db: db,
	}
}

func (c *CartImpl) Create(ctx context.Context, data *dto.CreateCartReq) error {
	err := c.db.WithContext(ctx).Table("carts").Create(data).Error

	if errPG, ok := err.(*pgconn.PgError); ok && errPG.Code == "23505" {
		return &errors.Response{HttpCode: 409, Message: "cart already exists"}
	}

	return err
}
