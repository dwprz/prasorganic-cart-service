package dto

type CreateCartReq struct {
	UserId    string `json:"user_id" validate:"required,min=21,max=21"`
	ProductId uint   `json:"product_id" validate:"required"`
	Quantity  uint   `json:"quantity" validate:"required"`
}
