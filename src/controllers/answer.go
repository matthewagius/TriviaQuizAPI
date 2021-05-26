package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/matthewagius/quiz/data"
	"github.com/matthewagius/quiz/models"
)

// func to check answers and generate final result
// @Description check answers and generate final result
// @Summary check answers and generate final result
// @Tags Answers
// @Accept json
// @Produce json
// @Success 200 {array} models.FinalResult
// @Router /api/answers [post]
func CheckAnswers(c *fiber.Ctx) error {
	var body models.SelectedAnswers
	err := c.BodyParser(&body)

	// if error return
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
		})
	}

	finalResult := CalculateScore(body)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    finalResult,
	})
}

func CalculateScore(answers models.SelectedAnswers) models.FinalResult {
	score := 0
	var finalResult models.FinalResult

	for _, userAnswer := range answers {

		//get the question and correct answer, if answer was correct update correct bool and score accordingly
		question, correctAnswer := (*data.GlobalData).QuestionSet.ReturnQuestionAndCorrectAnswer(userAnswer.QuestionID)
		correct := correctAnswer == userAnswer.Answer
		if correct {
			score++
		}

		finalResult.QuestionResults = append(finalResult.QuestionResults, models.QuestionResult{
			Question:       question,
			Correct:        correct,
			SelectedAnswer: userAnswer.Answer,
			CorrectAnswer:  correctAnswer,
		})
	}

	//save total score in and work out percentage ranking
	finalResult.PersonalScore = score
	finalResult.PercentageRanking = models.CalculatePercentageRanking((*data.GlobalData).Scores, score)
	(*data.GlobalData).Scores = append((*data.GlobalData).Scores, score)
	return finalResult
}
