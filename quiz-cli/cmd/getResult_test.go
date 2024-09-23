package cmd

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func Test_ExecuteGetResultCommand(t *testing.T) {

	authCmd := NewAuthCmd()
	authCmdBuff := bytes.NewBufferString("")
	authCmd.SetOut(authCmdBuff)
	authCmd.SetArgs([]string{"--username", "alice", "--password", "rainbow"})
	authCmd.Execute()
	
	t.Run("should return the quiz result successfully when all answers are correct", func(t *testing.T) {
		
		selectOptionsCommand := NewSelectOptionsCommand()
		selectOptionsBuff := bytes.NewBufferString("")
		selectOptionsCommand.SetOut(selectOptionsBuff)
		selectOptionsCommand.SetArgs([]string{"1:a,2:b,3:b"})
		selectOptionsCommand.Execute()

		getResultCommand := NewGetResultCmd()
		getResultBuff := bytes.NewBufferString("")
		getResultCommand.SetOut(getResultBuff)
		getResultCommand.Execute()
		actual, err := ioutil.ReadAll(getResultBuff)
		if err != nil {
			t.Fatal(err)
		}
		expected := "Number of correct answers: 3"
		if string(actual) != expected {
			t.Fatalf("expected \"%s\" got \"%s\"", expected, string(actual))
		}
	})
	t.Run("should return the quiz result successfully when one answer is correct", func(t *testing.T) {
		
		selectOptionsCommand := NewSelectOptionsCommand()
		selectOptionsBuff := bytes.NewBufferString("")
		selectOptionsCommand.SetOut(selectOptionsBuff)
		selectOptionsCommand.SetArgs([]string{"1:b,2:a,3:b"})
		selectOptionsCommand.Execute()
		
		getResultCommand := NewGetResultCmd()
		getResultBuff := bytes.NewBufferString("")
		getResultCommand.SetOut(getResultBuff)
		getResultCommand.Execute()
		actual, err := ioutil.ReadAll(getResultBuff)
		if err != nil {
			t.Fatal(err)
		}
		expected := "Number of correct answers: 1"
		if string(actual) != expected {
			t.Fatalf("expected \"%s\" got \"%s\"", expected, string(actual))
		}
	})
}