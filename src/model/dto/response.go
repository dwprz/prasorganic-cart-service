package dto

import "github.com/dwprz/prasorganic-cart-service/src/model/entity"

type Paging struct {
	TotalData int `json:"total_data"`
	Page      int `json:"page"`
	TotalPage int `json:"total_page"`
}

type DataWithPaging[T any] struct {
	Data   T       `json:"data"`
	Paging *Paging `json:"paging"`
}

type CartQueryRes struct {
	Cart      []byte `json:"cart"`
	TotalCart int    `json:"total_cart"`
}

type CartWithCountRes struct {
	Cart      []entity.Cart `json:"cart"`
	TotalCart int           `json:"total_cart"`
}

type ProductCartRes struct {
	UserId      string  `json:"user_id"`
	ProductId   uint32  `json:"product_id"`
	Quantity    uint32  `json:"quantity"`
	ProductName string  `json:"product_name"`
	Image       string  `json:"image"`
	Price       uint32  `json:"price"`
	Stock       uint32  `json:"stock"`
	Length      uint32  `json:"length"`
	Width       uint32  `json:"width"`
	Height      uint32  `json:"height"`
	Weight      float32 `json:"weight"`
}
