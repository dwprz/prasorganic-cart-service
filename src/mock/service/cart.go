package service

// import (
// 	"context"

// 	"github.com/dwprz/prasorganic-cart-service/src/model/dto"
// )

// type Cart interface {
// 	Create(ctx context.Context, data *dto.CreateCartReq) error
// 	GetByCurentUser(ctx context.Context, data *dto.GetCartByCurrentUserReq) (*dto.DataWithPaging[[]*dto.ProductCartRes], error)
// 	DeleteItem(ctx context.Context, data *dto.DeleteItemCartReq) error
// }

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

func (c *CartMock) GetByCurentUser(ctx context.Context, data *dto.GetCartByCurrentUserReq) (*dto.DataWithPaging[[]*dto.ProductCartRes], error){
	arguments := c.Mock.Called(ctx, data)

	if arguments.Get(0) == nil {
		return nil, arguments.Error(1)
	}

	return arguments.Get(0).(*dto.DataWithPaging[[]*dto.ProductCartRes]), arguments.Error(1)
}

func (c *CartMock) 	DeleteItem(ctx context.Context, data *dto.DeleteItemCartReq) error{
	arguments := c.Mock.Called(ctx, data)

	return arguments.Error(0)
}