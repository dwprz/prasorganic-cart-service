package entity

type CartQueryRes struct {
	TotalCart int    `json:"total_cart" gorm:"column:total_cart"`
	UserId    string `json:"user_id" gorm:"column:user_id"`
	ProductId uint32 `json:"product_id" gorm:"column:product_id"`
	Quantity  uint32 `json:"quantity" gorm:"column:quantity"`
}
