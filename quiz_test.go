package quiz

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
	"strings"
	"testing"
	"time"
)

var dummyWriter io.Writer
var dummyReader io.Reader
var dummyTimeDuration time.Duration

func TestQuiz(t *testing.T) {
	t.Run("Open and Read CSV file", func(t *testing.T) {

		tests := []struct {
			name           string
			inputCSV       string
			expectedOutput [][]string
		}{
			{name: "valid CSV",
				inputCSV: `5+5,10
7+3,10
1+1,2
8+3,11
1+2,3`,
				expectedOutput: [][]string{
					{"5+5", "10"},
					{"7+3", "10"},
					{"1+1", "2"},
					{"8+3", "11"},
					{"1+2", "3"},
				},
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				gotCSV := NewQuizGame(strings.NewReader(tt.inputCSV), dummyWriter, dummyReader, dummyTimeDuration)

				rec, err := gotCSV.Read()

				assertNoErr(t, err)

				assertCSV(t, rec, tt.expectedOutput)
			})
		}

	})

	t.Run("display correct questions", func(t *testing.T) {
		buf := &bytes.Buffer{}

		inputCSV := `5+5,10
7+3,10
1+1,2
8+3,11
1+2,3
`
		want := `Please hit enter
5+5
7+3
1+1
8+3
1+2
`
		displayMsg := fmt.Sprintf("You got %d out of %d questions correct\n", 5, 5)
		want += displayMsg
		gotCSV := NewQuizGame(strings.NewReader(inputCSV), buf, stringReader("10", "10", "2", "11", "3"), time.Duration(5*time.Second))

		csvData, err := gotCSV.Read()

		assertNoErr(t, err)

		gotCSV.QuizStart(csvData)
		assertString(t, buf.String(), want)
		assertInt(t, gotCSV.Counter, 5)

	})
	t.Run("run quiz with timer", func(t *testing.T) {
		buf := &bytes.Buffer{}

		inputCSV := `5+5,10
`
		want := `Please hit enter
5+5
time is up!
`
		displayMsg := fmt.Sprintf("You got %d out of %d questions correct\n", 0, 1)
		want += displayMsg
		gotCSV := NewQuizGame(strings.NewReader(inputCSV), buf, stringReader("10"), time.Duration(0*time.Millisecond))

		csvData, err := gotCSV.Read()

		assertNoErr(t, err)

		gotCSV.QuizStart(csvData)
		assertString(t, buf.String(), want)
	})

}

func stringReader(s ...string) io.Reader {
	joined := strings.Join(s, "\n")
	return strings.NewReader(joined)
}

func assertNoErr(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("error encountered while reading csv %v", err)
	}
}

func assertCSV(t testing.TB, got, want [][]string) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("CSV got %v, want %v", got, want)
	}
}

func assertString(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("String got %q, want %q", got, want)
	}
}

func assertInt(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("Int got %d, want %d", got, want)
	}
}
