package deliverry

import (
	"context"

	pb "github.com/dwprz/prasorganic-proto/protogen/product"
	"github.com/stretchr/testify/mock"
)

type ProductGrpcMock struct {
	mock.Mock
}

func NewProductGrpcMock() *ProductGrpcMock {
	return &ProductGrpcMock{
		Mock: mock.Mock{},
	}
}

func (p *ProductGrpcMock) FindManyByIds(ctx context.Context, productIds []uint32) ([]*pb.ProductCart, error) {
	arguments := p.Mock.Called(ctx, productIds)

	if arguments.Get(0) == nil {
		return nil, arguments.Error(1)
	}

	return arguments.Get(0).([]*pb.ProductCart), arguments.Error(1)
}
