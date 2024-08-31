package helper

import (
	"github.com/dwprz/prasorganic-cart-service/src/model/dto"
	"github.com/dwprz/prasorganic-cart-service/src/model/entity"
	pb "github.com/dwprz/prasorganic-proto/protogen/product"
)

func MapCartToProductCartRes(cart []*entity.Cart, products []*pb.ProductCart) []*dto.ProductCartRes {
	productMap := make(map[uint32]*pb.ProductCart)
	for _, product := range products {
		productMap[product.ProductId] = product
	}

	var result []*dto.ProductCartRes
	for _, c := range cart {
		if product, exists := productMap[c.ProductId]; exists {
			res := &dto.ProductCartRes{
				UserId:      c.UserId,
				ProductId:   c.ProductId,
				Quantity:    c.Quantity,
				ProductName: product.ProductName,
				Image:       product.Image,
				Price:       product.Price,
				Stock:       product.Stock,
				Length:      product.Length,
				Width:       product.Width,
				Height:      product.Height,
				Weight:      product.Weight,
			}

			result = append(result, res)
		}
	}

	return result
}

func MapCartQueryToEntities(data []*entity.CartQueryRes) (carts []*entity.Cart, totalCart int) {
    for _, item := range data {
        carts = append(carts, &entity.Cart{
            UserId:    item.UserId,
            ProductId: item.ProductId,
            Quantity:  item.Quantity,
        })
    }

    return carts, data[0].TotalCart
}