package dto

type CreateCartReq struct {
	UserId    string `json:"user_id" validate:"required,min=21,max=21"`
	ProductId uint32 `json:"product_id" validate:"required"`
	Quantity  uint32 `json:"quantity" validate:"required"`
}

type GetCartByCurrentUserReq struct {
	UserId string `json:"user_id" validate:"required,min=21,max=21"`
	Page   int    `json:"page" validate:"required,min=1,max=2"`
}

type DeleteItemCartReq struct {
	UserId    string `json:"user_id" validate:"required,min=21,max=21"`
	ProductId uint32 `json:"product_id" validate:"required"`
}
