package quiz

import (
	"bufio"
	"encoding/csv"
	"io"
	"strconv"
)

type QuizGame struct {
	csvReader io.Reader
	out       io.Writer
	in        *bufio.Scanner
	Counter   int
}

func NewQuizGame(csvReader io.Reader, out io.Writer, in io.Reader) *QuizGame {
	return &QuizGame{
		csvReader: csvReader,
		out:       out,
		in:        bufio.NewScanner(in),
		Counter:   0,
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
	for _, record := range records {
		q.out.Write([]byte(record[0] + "\n"))
		userInput := q.readLine()
		if userInput == record[1] {
			q.Counter++
		}
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

	q.displayWriter("You got " + counter + " out of " + lengthOfQuestion + " questions correct\n")
}

func (q *QuizGame) displayWriter(message string) {
	q.out.Write([]byte(message))
}
