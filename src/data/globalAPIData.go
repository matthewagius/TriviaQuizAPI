package data

import (
	"github.com/matthewagius/quiz/models"
)

type GlobalAPIData struct{
	QuestionSet models.QuizQuestions
	Scores []int
}

var GlobalData *GlobalAPIData


func Initialise(questions models.QuizQuestions) *GlobalAPIData{
	return &GlobalAPIData{
		QuestionSet: questions,
		Scores: []int{2,6,10}, //make up scores for previous players so that first player does not get 100% better than other players regardless of score
	}
}




