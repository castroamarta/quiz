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
		// bob misses all questions
		UserSolution["1"] = "b"
		UserSolution["2"] = "a"
		UserSolution["3"] = "a"
		actual := getQuizComparisonStats("bob")

		expected := float64(0)
		assert.Equal(t, expected, actual)

		// eve guesses all questions
		UserSolution["1"] = "a"
		UserSolution["2"] = "b"
		UserSolution["3"] = "b"
		actual = getQuizComparisonStats("eve")

		expected = float64(100)
		assert.Equal(t, expected, actual)

		// alice guesses 1 question
		UserSolution["1"] = "a"
		UserSolution["2"] = "a"
		UserSolution["3"] = "a"
		actual = getQuizComparisonStats("alice")

		expected = float64(50)
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

func Test_validateCredentials(t *testing.T) {
	t.Run(`should validate basic auth request credentials successfully`, func(t *testing.T) {
		actual := validateCredentials("alice", "rainbow")
		assert.Equal(t, nil, actual)
	})
	t.Run(`should return error when the credentials provided are invalid`, func(t *testing.T) {
		actual := validateCredentials("alice", "rainbo")
		expected := fmt.Errorf("user unauthorized to perform method call")
		assert.Equal(t, expected, actual)
	})
}