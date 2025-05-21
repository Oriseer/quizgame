package main

import (
	"log"
	"os"

	quiz "github.com/Oriseer/quizgame"
)

func main() {
	file, _ := os.Open("questions.csv")
	defer file.Close()

	quizGame := quiz.NewQuizGame(file, os.Stdout, os.Stdin)
	records, err := quizGame.Read()
	if err != nil {
		log.Fatalf("Error reading CSV file: %v", err)
	}
	quizGame.QuizStart(records)
}
