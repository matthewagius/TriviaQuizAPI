package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateScore(t *testing.T) {
	assert := assert.New(t)
	testData :=[]int{2,4,6,8}

	//case 1: scores in the middle - 2 out of 4 scored less ie 50%
	var score = CalculatePercentageRanking(testData, 5)
	assert.Equal(50,score, "Percentages should match")

	//case 2: highest score - 100%
	score = CalculatePercentageRanking(testData,10)
	assert.Equal(100,score, "Percentages should match")

	//case 3: lowest score - 0%
	score = CalculatePercentageRanking(testData,1)
	assert.Equal(0,score, "Percentages should match")
}