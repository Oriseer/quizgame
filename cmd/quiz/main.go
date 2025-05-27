package main

import (
	"flag"
	"log"
	"os"
	"time"

	quiz "github.com/Oriseer/quizgame"
)

var fileName = flag.String("file", "questions.csv", "CSV file containing quiz questions")
var timeLimit = flag.Int("time", 30, "quiz time limit in seconds")
var isShuffle = flag.Bool("shuffle", false, "shuffle questions in the quiz")

func main() {
	flag.Parse()
	file, _ := os.Open(*fileName)
	defer file.Close()

	quizGame := quiz.NewQuizGame(file, os.Stdout, os.Stdin, time.Duration(*timeLimit*int(time.Second)), *isShuffle)
	records, err := quizGame.Read()
	if err != nil {
		log.Fatalf("Error reading CSV file: %v", err)
	}
	quizGame.QuizStart(records)
}
