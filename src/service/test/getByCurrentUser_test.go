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
	"github.com/dwprz/prasorganic-cart-service/src/model/entity"
	serviceimpl "github.com/dwprz/prasorganic-cart-service/src/service"
	pb "github.com/dwprz/prasorganic-proto/protogen/product"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
)

// go test -p=1 -v ./src/service/test/... -count=1
// go test -run ^TestService_GetByCurrentUser$ -v ./src/service/test/ -count=1

type GetByCurrentUserTestSuite struct {
	suite.Suite
	cartService service.Cart
	productGrpc *deliverry.ProductGrpcMock
	cartRepo    *repository.CartMock
}

func (g *GetByCurrentUserTestSuite) SetupSuite() {
	g.cartRepo = repository.NewCartMock()

	g.productGrpc = deliverry.NewProductGrpcMock()
	productConn := new(grpc.ClientConn)

	grpcClient := client.NewGrpc(g.productGrpc, productConn)
	g.cartService = serviceimpl.NewCart(g.cartRepo, grpcClient)
}

func (g *GetByCurrentUserTestSuite) Test_Success() {
	req := &dto.GetCartByCurrentUserReq{
		UserId: "hyfa_5Sq7nQcaY6ACksXP",
		Page:   1,
	}

	cart := &dto.CartWithCountRes{
		Cart: []entity.Cart{
			{UserId: "hyfa_5Sq7nQcaY6ACksXP", ProductId: 10, Quantity: 5},
			{UserId: "hyfa_5Sq7nQcaY6ACksXP", ProductId: 20, Quantity: 10},
		},
		TotalCart: 2,
	}

	productCart := []*pb.ProductCart{
		{
			ProductId:   10,
			ProductName: "Product 1",
			Price:       10000,
		},
		{
			ProductId:   20,
			ProductName: "Product 2",
			Price:       20000,
		},
	}

	g.MockCartRepo_FindManyByUserId(req.UserId, cart, nil)
	g.MockProductGrpc_FindManyByIds([]uint32{10, 20}, productCart, nil)

	res, err := g.cartService.GetByCurentUser(context.Background(), req)
	assert.NoError(g.T(), err)

	assert.NotEmpty(g.T(), res)
}

func (g *GetByCurrentUserTestSuite) Test_NotFound() {
	req := &dto.GetCartByCurrentUserReq{
		UserId: "hyfa_5SqwascaY6ACksXP",
		Page:   1,
	}

	errRes := &errors.Response{HttpCode: 404, Message: "cart not found"}

	g.MockCartRepo_FindManyByUserId(req.UserId, nil, errRes)
	res, err := g.cartService.GetByCurentUser(context.Background(), req)
	assert.Error(g.T(), err)

	resErr, ok := err.(*errors.Response)
	assert.True(g.T(), ok)

	assert.Equal(g.T(), 404, resErr.HttpCode)
	assert.Nil(g.T(), res)
}

func (g *GetByCurrentUserTestSuite) MockCartRepo_FindManyByUserId(userId string, returnArg1 *dto.CartWithCountRes, returnArg2 error) {

	g.cartRepo.Mock.On("FindManyByUserId", mock.Anything, userId, mock.Anything, mock.Anything).Return(returnArg1, returnArg2)
}

func (g *GetByCurrentUserTestSuite) MockProductGrpc_FindManyByIds(productIds []uint32, returnArg1 []*pb.ProductCart, returnArg2 error) {

	g.productGrpc.Mock.On("FindManyByIds", mock.Anything, productIds).Return(returnArg1, returnArg2)
}

func TestService_GetByCurrentUser(t *testing.T) {
	suite.Run(t, new(GetByCurrentUserTestSuite))
}
