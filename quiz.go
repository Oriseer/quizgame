package quiz

import (
	"bufio"
	"encoding/csv"
	"io"
	"strconv"
	"time"
)

type QuizGame struct {
	csvReader    io.Reader
	out          io.Writer
	in           *bufio.Scanner
	Counter      int
	timeDuration time.Duration
}

func NewQuizGame(csvReader io.Reader, out io.Writer, in io.Reader, timeDuration time.Duration) *QuizGame {
	return &QuizGame{
		csvReader:    csvReader,
		out:          out,
		in:           bufio.NewScanner(in),
		Counter:      0,
		timeDuration: timeDuration,
	}
}

func (q *QuizGame) Read() ([][]string, error) {
	records, err := csv.NewReader(q.csvReader).ReadAll()

	if err != nil {
		return nil, err
	}

	return records, nil
}

func (q *QuizGame) QuizStart(records [][]string) {
	q.out.Write([]byte("Please hit enter\n"))
	correct := make(chan struct{})
	for _, record := range records {
		q.out.Write([]byte(record[0] + "\n"))
		go func() {
			userInput := q.readLine()
			if userInput == record[1] {
				correct <- struct{}{}
			}
		}()
		select {
		case <-correct:
			q.Counter++
		case <-time.After(q.timeDuration):
			q.out.Write([]byte("time is up!\n"))
			q.displayResults(len(records))
			return
		}
		q.displayResults(len(records))
	}

	q.displayResults(len(records))

}

func (q *QuizGame) readLine() string {
	if q.in.Scan() {
		return q.in.Text()
	}
	return ""
}

func (q *QuizGame) displayResults(numOfQuestion int) {
	lengthOfQuestion := strconv.Itoa(numOfQuestion)
	counter := strconv.Itoa(q.Counter)

	message := "You got " + counter + " out of " + lengthOfQuestion + " questions correct\n"

	q.displayWriter(message)
}

func (q *QuizGame) displayWriter(message string) {
	q.out.Write([]byte(message))
}
