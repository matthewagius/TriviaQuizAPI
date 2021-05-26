package models

type SelectedAnswer struct {
	QuestionID int    `json:"questionID"`
	Answer     string `json:"answer"`
}

type SelectedAnswers []SelectedAnswer

type QuestionResult struct {
	Question       string `json:"question"`
	Correct        bool   `json:"correct"`
	SelectedAnswer string `json:"selectedAnswer"`
	CorrectAnswer  string `json:"correctAnswer"`
}

type FinalResult struct {
	PersonalScore     int              `json:"personalScore"`
	QuestionResults   []QuestionResult `json:"questionResults"`
	PercentageRanking int              `json:"percentageRanking"`
}

//get percentage of people who scored less
func CalculatePercentageRanking(prevScores []int, userScore int) int {
	scoredLess := 0
	for _, prevScore := range prevScores {
		if prevScore < userScore {
			scoredLess++
		}
	}
	percentageRanking := (100 * scoredLess) / len(prevScores)
	return percentageRanking
}