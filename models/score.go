package models

import "time"

type Score struct {
	Date           string
	Score          float64
	CorrectAnswers int
}

func NewScore(times string, cAwnsers int) Score {
	return Score{
		Date:           time.Now().Format(time.DateOnly),
		Score:          float64(cAwnsers) / float64(20) * 100,
		CorrectAnswers: cAwnsers,
	}
}
