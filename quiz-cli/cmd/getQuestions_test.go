package cmd

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func Test_ExecuteGetQuestionsCommand(t *testing.T) {

	authCmd := NewAuthCmd()
	authCmdBuff := bytes.NewBufferString("")
	authCmd.SetOut(authCmdBuff)
	authCmd.SetArgs([]string{"--username", "alice", "--password", "rainbow"})
	authCmd.Execute()

	t.Run("should return the quiz questions successfully", func(t *testing.T) {
		questionsCmd := NewGetQuestionsCmd()
		questionsCmdBuff := bytes.NewBufferString("")
		questionsCmd.SetOut(questionsCmdBuff)
		questionsCmd.Execute()
		actual, err := ioutil.ReadAll(questionsCmdBuff)
		if err != nil {
			t.Fatal(err)
		}
		expected := parseQuestions(quizMock)
		if string(actual) != expected {
			t.Fatalf("expected \"%s\" got \"%s\"", expected, string(actual))
		}
	})
}

var quizMock = []Question{
	{
		ID:          "1",
		Description: "Which OS is more popular?",
		Options: []Option{
			{
				ID:          "a",
				Description: "MacOS",
			},
			{
				ID:          "b",
				Description: "Windows",
			},
			{
				ID:          "c",
				Description: "Linux",
			},
		},
	},
	{
		ID:          "2",
		Description: "Which bike was sold more often on the month of July?",
		Options: []Option{
			{
				ID:          "a",
				Description: "Road Bike",
			},
			{
				ID:          "b",
				Description: "Moutain Bike",
			},
		},
	},
	{
		ID:          "3",
		Description: "Which genre buys more apples?",
		Options: []Option{
			{
				ID:          "a",
				Description: "Females",
			},
			{
				ID:          "b",
				Description: "Males",
			},
		},
	},
}
