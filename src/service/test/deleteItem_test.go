package test

import (
	"context"
	"testing"

	"github.com/dwprz/prasorganic-cart-service/src/core/grpc/client"
	"github.com/dwprz/prasorganic-cart-service/src/interface/service"
	"github.com/dwprz/prasorganic-cart-service/src/mock/deliverry"
	"github.com/dwprz/prasorganic-cart-service/src/mock/repository"
	"github.com/dwprz/prasorganic-cart-service/src/model/dto"
	serviceimpl "github.com/dwprz/prasorganic-cart-service/src/service"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
)

// go test -p=1 -v ./src/service/test/... -count=1
// go test -run ^TestService_DeleteItem$ -v ./src/service/test/ -count=1

type DeleteItemTestSuite struct {
	suite.Suite
	cartService service.Cart
	productGrpc *deliverry.ProductGrpcMock
	cartRepo    *repository.CartMock
}

func (g *DeleteItemTestSuite) SetupSuite() {
	g.cartRepo = repository.NewCartMock()

	g.productGrpc = deliverry.NewProductGrpcMock()
	productConn := new(grpc.ClientConn)

	grpcClient := client.NewGrpc(g.productGrpc, productConn)
	g.cartService = serviceimpl.NewCart(g.cartRepo, grpcClient)
}

func (g *DeleteItemTestSuite) Test_Success() {
	req := &dto.DeleteItemCartReq{
		UserId:    "hyfa_5Sq7nQcaY6ACksXP",
		ProductId: 1,
	}

	g.MockCartRepo_DeleteItem(req, nil)

	err := g.cartService.DeleteItem(context.Background(), req)
	assert.NoError(g.T(), err)
}

func (g *DeleteItemTestSuite) Test_InvalidUserId() {
	req := &dto.DeleteItemCartReq{
		UserId:    "invalid-user-id",
		ProductId: 1,
	}

	err := g.cartService.DeleteItem(context.Background(), req)
	assert.Error(g.T(), err)

	resErr, ok := err.(validator.ValidationErrors)
	assert.True(g.T(), ok)

	assert.Equal(g.T(), resErr[0].Field(), "UserId")
}

func (g *DeleteItemTestSuite) MockCartRepo_DeleteItem(data *dto.DeleteItemCartReq, returnArg1 error) {

	g.cartRepo.Mock.On("DeleteItem", mock.Anything, data).Return(returnArg1)
}

func TestService_DeleteItem(t *testing.T) {
	suite.Run(t, new(DeleteItemTestSuite))
}
