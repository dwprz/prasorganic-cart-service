package delivery

import (
	"context"
	"fmt"
	"log"

	"github.com/dwprz/prasorganic-cart-service/src/core/grpc/interceptor"
	"github.com/dwprz/prasorganic-cart-service/src/infrastructure/cbreaker"
	"github.com/dwprz/prasorganic-cart-service/src/infrastructure/config"
	"github.com/dwprz/prasorganic-cart-service/src/interface/deliverry"
	pb "github.com/dwprz/prasorganic-proto/protogen/product"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ProductGrpcImpl struct {
	client   pb.ProductServiceClient
}

func NewProductGrpc(unaryRequest *interceptor.UnaryRequest) (deliverry.ProductGrpc, *grpc.ClientConn) {
	var opts []grpc.DialOption
	opts = append(
		opts,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(unaryRequest.AddBasicAuth),
	)

	conn, err := grpc.NewClient(config.Conf.ApiGateway.BaseUrl, opts...)
	if err != nil {
		log.Fatalf("new otp grpc client: %v", err.Error())
	}

	client := pb.NewProductServiceClient(conn)

	return &ProductGrpcImpl{
		client:   client,
	}, conn
}

func (p *ProductGrpcImpl) FindManyByIds(ctx context.Context, productIds []uint32) ([]*pb.ProductCart, error) {
	res, err := cbreaker.ProductGrpc.Execute(func() (any, error) {
		res, err := p.client.FindManyByIdsForCart(ctx, &pb.ProductIds{
			Ids: productIds,
		})

		return res, err
	})

	if err != nil {
		return nil, err
	}

	products, ok := res.(*pb.ProductsCartRes)
	if !ok {
		return nil, fmt.Errorf("client.CartGrpcImpl/FindManyByIds | unexpected type: %T", res)
	}

	return products.Data, err
}