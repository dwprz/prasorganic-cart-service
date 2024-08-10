package handler

import (
	"github.com/dwprz/prasorganic-cart-service/src/interface/service"
	"github.com/dwprz/prasorganic-cart-service/src/model/dto"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type Cart struct {
	cartService service.Cart
}

func NewCart(cs service.Cart) *Cart {
	return &Cart{
		cartService: cs,
	}
}

func (h *Cart) Create(c *fiber.Ctx) error {
	req := new(dto.CreateCartReq)

	if err := c.BodyParser(req); err != nil {
		return err
	}

	userData := c.Locals("user_data").(jwt.MapClaims)
	req.UserId = userData["user_id"].(string)

	err := h.cartService.Create(c.Context(), req)
	if err != nil {
		return err
	}

	return c.Status(201).JSON(fiber.Map{"data": "created item cart successfully"})
}
