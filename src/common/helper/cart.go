package helper

import "github.com/dwprz/prasorganic-cart-service/src/model/entity"

func GetProductIdsInCart(cart []entity.Cart) []uint32 {
	var productIds []uint32

	for _, item := range cart {
		productIds = append(productIds, uint32(item.ProductId))
	}

	return productIds
}
