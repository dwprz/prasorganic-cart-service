package test

import (
	"context"
	"testing"

	"github.com/dwprz/prasorganic-cart-service/src/common/errors"
	"github.com/dwprz/prasorganic-cart-service/src/core/grpc/client"
	"github.com/dwprz/prasorganic-cart-service/src/interface/service"
	"github.com/dwprz/prasorganic-cart-service/src/mock/deliverry"
	"github.com/dwprz/prasorganic-cart-service/src/mock/repository"
	"github.com/dwprz/prasorganic-cart-service/src/model/dto"
	serviceimpl "github.com/dwprz/prasorganic-cart-service/src/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
)

// go test -p=1 -v ./src/service/test/... -count=1
// go test -run ^TestService_CreateCart$ -v ./src/service/test/ -count=1

type CreateCartTestSuite struct {
	suite.Suite
	cartService service.Cart
	cartRepo    *repository.CartMock
}

func (c *CreateCartTestSuite) SetupSuite() {
	c.cartRepo = repository.NewCartMock()

	productDelivery := deliverry.NewProductGrpcMock()
	productConn := new(grpc.ClientConn)

	grpcClient := client.NewGrpc(productDelivery, productConn)
	c.cartService = serviceimpl.NewCart(c.cartRepo, grpcClient)
}

func (c *CreateCartTestSuite) Test_Success() {
	req := &dto.CreateCartReq{
		UserId:    "hyfa_5Sq7nQcaY6ACksXP",
		ProductId: 10,
		Quantity:  5,
	}

	totalCart := 10

	c.MockCartRepo_CountByUserId(req.UserId, int64(totalCart), nil)
	c.MockCartRepo_Create(req, nil)

	err := c.cartService.Create(context.Background(), req)
	assert.NoError(c.T(), err)
}

func (c *CreateCartTestSuite) Test_Limit() {
	req := &dto.CreateCartReq{
		UserId:    "hyfa_5Sq7nQcaY6fCdaXP",
		ProductId: 10,
		Quantity:  5,
	}

	totalCart := 40

	c.MockCartRepo_CountByUserId(req.UserId, int64(totalCart), nil)

	err := c.cartService.Create(context.Background(), req)
	assert.Error(c.T(), err)

	resErr, ok := err.(*errors.Response)
	assert.True(c.T(), ok)

	assert.Equal(c.T(), 400, resErr.HttpCode)
}

func (c *CreateCartTestSuite) MockCartRepo_CountByUserId(userId string, returnArg1 int64, returnArg2 error) {

	c.cartRepo.Mock.On("CountByUserId", mock.Anything, userId).Return(returnArg1, returnArg2)
}

func (c *CreateCartTestSuite) MockCartRepo_Create(data *dto.CreateCartReq, returnArg1 error) {

	c.cartRepo.Mock.On("Create", mock.Anything, data).Return(returnArg1)
}

func TestService_CreateCart(t *testing.T) {
	suite.Run(t, new(CreateCartTestSuite))
}
