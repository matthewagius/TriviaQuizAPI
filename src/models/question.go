package models

//A single question
type QuizQuestion struct {
	Id              int	`json:"id"`
	Topic           string  `json:"topic"`
	Question        string  `json:"question"`
	Answers         [4]string `json:"answers"`
	CorrectAnswer string `json:"correctAnswer"`
}

//A slice of questions
type QuizQuestions  []QuizQuestion  

func (questions QuizQuestions) ReturnQuestionAndCorrectAnswer(questionID int) (string,string){
	var question QuizQuestion
	for i:= range questions{
		if questions[i].Id == questionID{
			question = questions[i]
			break
		}
	}
	return question.Question, question.CorrectAnswer
}

