package service

import (
	"context"

	"github.com/dwprz/prasorganic-cart-service/src/interface/repository"
	"github.com/dwprz/prasorganic-cart-service/src/interface/service"
	"github.com/dwprz/prasorganic-cart-service/src/model/dto"
	"github.com/go-playground/validator/v10"
)

type CartImpl struct {
	cartRepo repository.Cart
	validate *validator.Validate
}

func NewCart(cr repository.Cart, v *validator.Validate) service.Cart {
	return &CartImpl{
		cartRepo: cr,
		validate: v,
	}
}

func (c *CartImpl) Create(ctx context.Context, data *dto.CreateCartReq) error {
	if err := c.validate.Struct(data); err != nil {
		return err
	}

	err := c.cartRepo.Create(ctx, data)
	return err
}
