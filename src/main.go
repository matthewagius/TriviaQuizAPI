package main

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/matthewagius/quiz/controllers"
	"github.com/matthewagius/quiz/data"

	_ "github.com/matthewagius/quiz/docs"
	"github.com/matthewagius/quiz/routes"
)



func main() {

	qstns, loadingError := controllers.LoadInitialData()
	data.GlobalData = data.Initialise(qstns)
	
	if(loadingError != nil){
		panic(loadingError)
	}
	app := fiber.New()
	setupRoutes(app)
	err := app.Listen(":8000")

	if err != nil {
		panic(err)
	}
}

func setupRoutes(app *fiber.App){
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success":  true,
			"message": "Trivia Quiz API endpoint. For documentation visit /swagger/",
		})
	})

	app.Get("/swagger/",swagger.Handler)

	api := app.Group("/api")
	api.Get("", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "You are at the api endpoint ",
		})
	})
	
	routes.AnswerRoute(api.Group("/answers"))
	routes.QuestionRoute(api.Group("/questions"))
	routes.SwaggerRoute(app.Group("/swagger"))
}

