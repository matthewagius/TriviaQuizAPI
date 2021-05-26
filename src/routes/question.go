package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/matthewagius/quiz/controllers"
)

func QuestionRoute(route fiber.Router) {
	route.Get("", controllers.Questions)
}
