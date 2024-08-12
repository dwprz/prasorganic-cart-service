package service

import (
	"context"

	"github.com/dwprz/prasorganic-cart-service/src/common/errors"
	"github.com/dwprz/prasorganic-cart-service/src/common/helper"
	"github.com/dwprz/prasorganic-cart-service/src/core/grpc/client"
	"github.com/dwprz/prasorganic-cart-service/src/interface/repository"
	"github.com/dwprz/prasorganic-cart-service/src/interface/service"
	"github.com/dwprz/prasorganic-cart-service/src/model/dto"
	"github.com/go-playground/validator/v10"
)

type CartImpl struct {
	cartRepo   repository.Cart
	grpcClient *client.Grpc
	validate   *validator.Validate
}

func NewCart(cr repository.Cart, gc *client.Grpc, v *validator.Validate) service.Cart {
	return &CartImpl{
		cartRepo:   cr,
		grpcClient: gc,
		validate:   v,
	}
}

func (c *CartImpl) Create(ctx context.Context, data *dto.CreateCartReq) error {
	if err := c.validate.Struct(data); err != nil {
		return err
	}

	totalCart, err := c.cartRepo.CountByUserId(ctx, data.UserId)
	if err != nil {
		return err
	}

	if totalCart >= 40 {
		return &errors.Response{HttpCode: 400, Message: "sorry, this user already has 40 cart item"}
	}

	err = c.cartRepo.Create(ctx, data)
	return err
}

func (c *CartImpl) GetByCurentUser(ctx context.Context, data *dto.GetCartByCurrentUserReq) (*dto.DataWithPaging[[]*dto.ProductCartRes], error) {
	if err := c.validate.Struct(data); err != nil {
		return nil, err
	}

	limit, offset := helper.CreateLimitAndOffset(data.Page)

	res, err := c.cartRepo.FindManyByUserId(ctx, data.UserId, limit, offset)
	if err != nil {
		return nil, err
	}

	productIds := helper.GetProductIdsInCart(res.Cart)
	products, err := c.grpcClient.Product.FindManyByIds(ctx, productIds)
	if err != nil {
		return nil, err
	}

	productsCart := helper.MapCartToProductCartRes(res.Cart, products)

	return helper.FormatPagedData(productsCart, res.TotalCart, data.Page, limit), nil
}
