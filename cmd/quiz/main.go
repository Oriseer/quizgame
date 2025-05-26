package main

import (
	"flag"
	"log"
	"os"
	"time"

	quiz "github.com/Oriseer/quizgame"
)

var fileName = flag.String("file", "questions.csv", "CSV file containing quiz questions")

func main() {
	flag.Parse()
	file, _ := os.Open(*fileName)
	defer file.Close()

	quizGame := quiz.NewQuizGame(file, os.Stdout, os.Stdin, time.Second*5)
	records, err := quizGame.Read()
	if err != nil {
		log.Fatalf("Error reading CSV file: %v", err)
	}
	quizGame.QuizStart(records)
}
