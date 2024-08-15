package service

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/lucas-o-ferreira/goscore/models"
)

type Score = models.Score

func LoadScores(filename string) ([]Score, error) {
	file, err := os.Open(filename)

	if os.IsNotExist(err) {
		return []Score{}, nil
	} else if err != nil {
		return nil, err
	}
	defer file.Close()

	var scores []Score
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&scores); err != nil {
		return nil, err
	}

	return scores, nil
}

func SaveScores(filename string, scores []Score) error {
	data, err := json.MarshalIndent(scores, "", "")

	if err != nil {
		return err
	}

	if err := os.WriteFile(filename, data, 0644); err != nil {
		return err
	}

	return nil
}

func ListScores(scores []Score) {
	if len(scores) == 0 {
		fmt.Println("No scores to list")
		return
	}

	fmt.Println("Date\t\tCorrect\tTotal\tPercentage")
	var totalPercentage float64

	for _, score := range scores {
		fmt.Printf("%s\t%d\t20\t%.2f%%\n", score.Date, score.CorrectAnswers, score.Score)
		totalPercentage += score.Score
	}

	avg := totalPercentage / float64(len(scores))
	fmt.Printf("\nAverage Percentage: %.2f%%\n", avg)

}

func AddScore(scores *[]Score) {
	var correct string
	fmt.Print("Enter the number of correct answers (out of 20): ")
	fmt.Scanln(&correct)

	correctI, err := strconv.Atoi(correct)
	if err != nil || correctI < 0 || correctI > 20 {
		fmt.Println("Invalid input. Please enter a number between 0 and 20.")
		return
	}

	newScore := models.NewScore("", correctI)
	*scores = append(*scores, newScore)
	fmt.Println("Score added successfully.")
}
