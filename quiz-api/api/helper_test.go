package api
import (
	"fmt"
	"testing"

	"github.com/go-playground/assert/v2"
)

func Test_getQuizResult(t *testing.T) {
	t.Run(`should calculate the quiz result successfully`, func(t *testing.T) {
		UserSolution["1"] = "b"
		UserSolution["2"] = "b"
		UserSolution["3"] = "b"
		actual := getQuizResult()
		expected := 2
		assert.Equal(t, expected, actual)
	})
}

func Test_getQuizStats(t *testing.T) {
	t.Run(`should calculate the quiz stats successfully`, func(t *testing.T) {
		UserSolution["1"] = "b"
		UserSolution["2"] = "b"
		UserSolution["3"] = "b"
		actual := getQuizStats("alice")
		expected := float64(50)
		assert.Equal(t, expected, actual)
	})
}

func Test_validateAnswer(t *testing.T) {
	tests := []struct {
		name     string
		answer   Answer
		expected error
	}{
		{
			name: "should return nil when a valid answer is provided",
			answer: Answer{
				QuestionID: "1",
				OptionID:   "a",
			},
			expected: nil,
		},
		{
			name: "should return an error when an invalid question ID is provided",
			answer: Answer{
				QuestionID: "10c",
				OptionID:   "a",
			},
			expected: fmt.Errorf("invalid question ID provided: 10c"),
		},
		{
			name: "should return an error when an invalid option ID is provided",
			answer: Answer{
				QuestionID: "3",
				OptionID:   "12ecainvdf",
			},
			expected: fmt.Errorf("question 3 only has the options [{a Females} {b Males}]"),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := validateAnswer(test.answer)
			assert.Equal(t, test.expected, actual)
		})
	}
}
