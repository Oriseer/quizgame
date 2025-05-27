package quiz

import (
	"bufio"
	"encoding/csv"
	"io"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type Timer interface {
	NewTimer(d time.Duration) *time.Timer
}

type quizTimer struct{}

func (n quizTimer) NewTimer(d time.Duration) *time.Timer {
	return time.NewTimer(d)
}

func (q *QuizGame) SetTimer(timer Timer) *time.Timer {
	return timer.NewTimer(q.timeDuration)
}

type QuizGame struct {
	csvReader    io.Reader
	out          io.Writer
	in           *bufio.Scanner
	Counter      int
	timeDuration time.Duration
	isShuffle    bool
}

func NewQuizGame(csvReader io.Reader, out io.Writer, in io.Reader,
	timeDuration time.Duration, shuffle bool) *QuizGame {
	return &QuizGame{
		csvReader:    csvReader,
		out:          out,
		in:           bufio.NewScanner(in),
		Counter:      0,
		timeDuration: timeDuration,
		isShuffle:    shuffle,
	}
}

func (q *QuizGame) Read() ([][]string, error) {
	records, err := csv.NewReader(q.csvReader).ReadAll()

	if err != nil {
		return nil, err
	}

	if q.isShuffle {
		rand.Shuffle(len(records), func(i, j int) {
			records[i], records[j] = records[j], records[i]
		})
	}

	return records, nil
}

func (q *QuizGame) QuizStart(records [][]string) {
	q.out.Write([]byte("Please hit enter\n"))
	started := q.readLine()
	started = strings.TrimSuffix(started, "\n")
	timer := q.SetTimer(quizTimer{})

	go func() {
		<-timer.C
		q.out.Write([]byte("time is up!\n"))
		q.displayResults(len(records))
		os.Exit(0) // Exit the program after time is up
	}()

	if started == "" {
		q.start(records)
	}

	q.displayResults(len(records))
}

func (q *QuizGame) start(records [][]string) {
	for _, record := range records {
		question := record[0]
		correctAnswer := record[1]
		q.out.Write([]byte(question + "\n"))
		userAnswer := strings.TrimSpace(q.readLine())
		if correctAnswer == userAnswer {
			q.Counter++
		}
	}

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
