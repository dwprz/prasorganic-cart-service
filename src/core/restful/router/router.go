package router

import (
	"github.com/dwprz/prasorganic-cart-service/src/core/restful/handler"
	"github.com/dwprz/prasorganic-cart-service/src/core/restful/middleware"
	"github.com/gofiber/fiber/v2"
)

func Create(app *fiber.App, h *handler.CartRESTful, m *middleware.Middleware) {
	// all
	app.Add("POST", "/api/carts", m.VerifyJwt, h.Create)
	app.Add("GET", "/api/carts/users/current", m.VerifyJwt, h.GetByCurrentUser)
	app.Add("DELETE", "/api/carts/products/:productId", m.VerifyJwt, h.DeleteItem)
}
