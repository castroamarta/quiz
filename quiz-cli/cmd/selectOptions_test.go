package cmd

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func Test_ExecuteSelectOptionsCommand(t *testing.T) {
	t.Run("should select an option successfully", func(t *testing.T) {
		cmd := NewSelectOptionsCommand()
		b := bytes.NewBufferString("")
		cmd.SetOut(b)
		cmd.SetArgs([]string{"1:b"})
		cmd.Execute()
		actual, err := ioutil.ReadAll(b)
		if err != nil {
			t.Fatal(err)
		}
		expected := parseSelectedOptions([]Answer{answersMock[0]})
		if string(actual) != expected {
			t.Fatalf("expected \"%s\" got \"%s\"", expected, string(actual))
		}
	})
	t.Run("should select all quiz options successfully", func(t *testing.T) {
		cmd := NewSelectOptionsCommand()
		b := bytes.NewBufferString("")
		cmd.SetOut(b)
		cmd.SetArgs([]string{"1:b,2:b,3:a"})
		cmd.Execute()
		actual, err := ioutil.ReadAll(b)
		if err != nil {
			t.Fatal(err)
		}
		expected := parseSelectedOptions(answersMock)
		if string(actual) != expected {
			t.Fatalf("expected \"%s\" got \"%s\"", expected, string(actual))
		}
	})
}

var answersMock = []Answer{
	{
		QuestionID: "1",
		OptionID: "b",
	},
	{
		QuestionID: "2",
		OptionID: "b",
	},
	{
		QuestionID: "3",
		OptionID: "a",
	},
}