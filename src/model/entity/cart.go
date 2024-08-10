package entity

type Cart struct {
	CartItemId uint   `json:"cart_item_id" gorm:"column:cart_item_id;<-:create"`
	UserId     string `json:"user_id" gorm:"column:user_id;<-:create"`
	ProductId  uint   `json:"product_id" gorm:"column:product_id;<-:create"`
	Quantity   uint   `json:"quantity" gorm:"column:quantity;<-:create"`
}

func (c *Cart) TableName() string {
	return "carts"
}
