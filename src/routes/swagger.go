package routes

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
)

func SwaggerRoute(route fiber.Router){
	route.Get("*",swagger.Handler)
}