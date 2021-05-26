package controllers

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/matthewagius/quiz/data"
	"github.com/matthewagius/quiz/models"
)

//function to load all questions from the questions.json file
func  LoadInitialData() (models.QuizQuestions,error){
	//open json file
	jsonFile, err := os.Open("questions.json")
	//return if there's an error
	if(err != nil){
		return nil, err
	}
	//read open xmlFile as a byte array
	fileData, readErr := ioutil.ReadAll(jsonFile)
	//return if there's an error
	if(readErr != nil){
		return nil,readErr
	}
	var quizQuestions models.QuizQuestions
	//unmarshal the file's contents into questions receiver
	unmarshalErr := json.Unmarshal(fileData,&quizQuestions)
	if(unmarshalErr != nil){
		//close json file
	    jsonFile.Close()
		return nil,unmarshalErr		
	}
	//close json file
	jsonFile.Close()
	//return full list of questions
	return quizQuestions,nil

}

// Questions func gets all questions in the quiz.
// @Description Get all available questions.
// @Summary Get all available questions.
// @Tags Questions
// @Accept json
// @Produce json
// @Success 200 {array} models.QuizQuestion
// @Router /api/questions [get]
func Questions(c *fiber.Ctx) error {
	//return questions from global variable in api package
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":   (*data.GlobalData).QuestionSet,
		
	})
}
