package repository

import (
	"context"

	"github.com/dwprz/prasorganic-cart-service/src/model/dto"
	"github.com/stretchr/testify/mock"
)

type CartMock struct {
	mock.Mock
}

func NewCartMock() *CartMock {
	return &CartMock{
		Mock: mock.Mock{},
	}
}

func (c *CartMock) Create(ctx context.Context, data *dto.CreateCartReq) error {
	arguments := c.Mock.Called(ctx, data)

	return arguments.Error(0)
}

func (c *CartMock) FindManyByUserId(ctx context.Context, userId string, limit, offset int) (*dto.CartWithCountRes, error) {
	arguments := c.Mock.Called(ctx, userId)

	if arguments.Get(0) == nil {
		return nil, arguments.Error(1)
	}

	return arguments.Get(0).(*dto.CartWithCountRes), arguments.Error(1)
}

func (c *CartMock) CountByUserId(ctx context.Context, userId string) (totalCart int64, err error) {
	arguments := c.Mock.Called(ctx, userId)

	return arguments.Get(0).(int64), arguments.Error(1)
}

func (c *CartMock) DeleteItem(ctx context.Context, data *dto.DeleteItemCartReq) error {
	arguments := c.Mock.Called(ctx, data)

	return arguments.Error(0)
}
