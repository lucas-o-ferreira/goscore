package ui

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func PrintMenu() {
	fmt.Println("1. Add score")
	fmt.Println("2. List scores")
	fmt.Println("3. Exit")
	fmt.Print("Enter choice: ")
}

func GetUserChoice() int {
	var choice int
	fmt.Scanln(&choice)
	return choice
}

func PrintInvalidChoice() {
	fmt.Println("Invalid choice")
}

func PrintSeparator() {
	fmt.Println("_______________________________________")
}

func ClearScreen() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func WaitForUserInput() {
	fmt.Println("Press 'Enter' to continue...")
	fmt.Scanln()
}

func PrintError(message string, err error) {
	fmt.Printf("%s %v\n", message, err)
}
