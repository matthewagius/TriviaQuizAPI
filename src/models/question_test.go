package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// function to map multiple return values to a slice of values that the assert function can compare
func F(a, b interface{}) []interface{} {
	return []interface{}{a, b}
}

func TestReturnQuestionAndCorrectAnswer(t *testing.T){
	assert:= assert.New(t)
	var testData = QuizQuestions{
		{
			Id: 1,
			Topic: "test",
			Question: "question 1",
			Answers: [4]string {"answer 1","answer 2","answer 3","answer 4"},
			CorrectAnswer: "answer 1",
		},
		{
			Id: 2,
			Topic: "test",
			Question: "question 2",
			Answers: [4]string {"answer 1","answer 2","answer 3","answer 4"},
			CorrectAnswer: "answer 2",
		},
		{
			Id: 3,
			Topic: "test",
			Question: "question 3",
			Answers: [4]string {"answer 1","answer 2","answer 3","answer 4"},
			CorrectAnswer: "answer 3",
		},
		{
			Id: 4,
			Topic: "test",
			Question: "question 3",
			Answers: [4]string {"answer 1","answer 2","answer 3","answer 4"},
			CorrectAnswer: "answer 4",
		},
	}

	for _, test := range testData {
		question,answer := testData.ReturnQuestionAndCorrectAnswer(test.Id)
        assert.Equal(F(test.Question,test.CorrectAnswer),F(question,answer), "Questions should be the same")
    }
}