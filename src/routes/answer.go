package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/matthewagius/quiz/controllers"
)

func AnswerRoute(route fiber.Router) {
	route.Post("", controllers.CheckAnswers)
}