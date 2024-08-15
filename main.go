package main

import (
	"github.com/lucas-o-ferreira/goscore/service"
	"github.com/lucas-o-ferreira/goscore/ui"
)

const scoreFile = "score.json"

func main() {
	scores, err := service.LoadScores(scoreFile)

	if err != nil {
		ui.PrintError("Failed to load save scores", err)
	}

	for {
		ui.ClearScreen()
		ui.PrintMenu()

		choice := ui.GetUserChoice()

		switch choice {
		case 1:
			ui.ClearScreen()
			service.AddScore(&scores)
			if err := service.SaveScores(scoreFile, scores); err != nil {
				ui.PrintError("Failed to save file:", err)
			}
		case 2:
			ui.ClearScreen()
			ui.PrintSeparator()
			service.ListScores(scores)
			ui.PrintSeparator()
			ui.WaitForUserInput()
		case 3:
			return
		default:
			ui.PrintInvalidChoice()
		}
	}
}
