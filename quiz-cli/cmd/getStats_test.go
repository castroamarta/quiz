package cmd

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func Test_ExecuteGetStatsCommand(t *testing.T) {
	t.Run("should return the quiz stats successfully for configuration 1", func(t *testing.T) {

		// alice guesses one question
		authCmd := NewAuthCmd()
		authCmdBuff := bytes.NewBufferString("")
		authCmd.SetOut(authCmdBuff)
		authCmd.SetArgs([]string{"--username", "alice", "--password", "rainbow"})
		authCmd.Execute()

		selectOptionsCommand := NewSelectOptionsCommand()
		selectOptionsBuff := bytes.NewBufferString("")
		selectOptionsCommand.SetOut(selectOptionsBuff)
		selectOptionsCommand.SetArgs([]string{"1:a,2:a,3:a"})
		selectOptionsCommand.Execute()

		// bob misses all question
		authCmd = NewAuthCmd()
		authCmdBuff = bytes.NewBufferString("")
		authCmd.SetOut(authCmdBuff)
		authCmd.SetArgs([]string{"--username", "bob", "--password", "flower"})
		authCmd.Execute()

		selectOptionsCommand = NewSelectOptionsCommand()
		selectOptionsBuff = bytes.NewBufferString("")
		selectOptionsCommand.SetOut(selectOptionsBuff)
		selectOptionsCommand.SetArgs([]string{"1:b,2:a,3:a"})
		selectOptionsCommand.Execute()

		// eve misses all question
		authCmd = NewAuthCmd()
		authCmdBuff = bytes.NewBufferString("")
		authCmd.SetOut(authCmdBuff)
		authCmd.SetArgs([]string{"--username", "eve", "--password", "boat"})
		authCmd.Execute()

		selectOptionsCommand = NewSelectOptionsCommand()
		selectOptionsBuff = bytes.NewBufferString("")
		selectOptionsCommand.SetOut(selectOptionsBuff)
		selectOptionsCommand.SetArgs([]string{"1:b,2:a,3:a"})
		selectOptionsCommand.Execute()

		getStatsCommand := NewGetStatsCmd()
		getStatsBuff := bytes.NewBufferString("")
		getStatsCommand.SetOut(getStatsBuff)
		getStatsCommand.Execute()
		actual, err := ioutil.ReadAll(getStatsBuff)
		if err != nil {
			t.Fatal(err)
		}
		expected := "You were better than 0% of all quizzers"
		if string(actual) != expected {
			t.Fatalf("expected \"%s\" got \"%s\"", expected, string(actual))
		}
	})

	t.Run("should return the quiz stats successfully for configuration 2", func(t *testing.T) {

		// alice guesses all question
		authCmd := NewAuthCmd()
		authCmdBuff := bytes.NewBufferString("")
		authCmd.SetOut(authCmdBuff)
		authCmd.SetArgs([]string{"--username", "alice", "--password", "rainbow"})
		authCmd.Execute()

		selectOptionsCommand := NewSelectOptionsCommand()
		selectOptionsBuff := bytes.NewBufferString("")
		selectOptionsCommand.SetOut(selectOptionsBuff)
		selectOptionsCommand.SetArgs([]string{"1:a,2:b,3:b"})
		selectOptionsCommand.Execute()

		// bob guesses one question
		authCmd = NewAuthCmd()
		authCmdBuff = bytes.NewBufferString("")
		authCmd.SetOut(authCmdBuff)
		authCmd.SetArgs([]string{"--username", "bob", "--password", "flower"})
		authCmd.Execute()

		selectOptionsCommand = NewSelectOptionsCommand()
		selectOptionsBuff = bytes.NewBufferString("")
		selectOptionsCommand.SetOut(selectOptionsBuff)
		selectOptionsCommand.SetArgs([]string{"1:a,2:a,3:a"})
		selectOptionsCommand.Execute()

		// eve gueses two questions
		authCmd = NewAuthCmd()
		authCmdBuff = bytes.NewBufferString("")
		authCmd.SetOut(authCmdBuff)
		authCmd.SetArgs([]string{"--username", "eve", "--password", "boat"})
		authCmd.Execute()

		selectOptionsCommand = NewSelectOptionsCommand()
		selectOptionsBuff = bytes.NewBufferString("")
		selectOptionsCommand.SetOut(selectOptionsBuff)
		selectOptionsCommand.SetArgs([]string{"1:a,2:b,3:a"})
		selectOptionsCommand.Execute()

		getStatsCommand := NewGetStatsCmd()
		getStatsBuff := bytes.NewBufferString("")
		getStatsCommand.SetOut(getStatsBuff)
		getStatsCommand.Execute()
		actual, err := ioutil.ReadAll(getStatsBuff)
		if err != nil {
			t.Fatal(err)
		}
		expected := "You were better than 50% of all quizzers"
		if string(actual) != expected {
			t.Fatalf("expected \"%s\" got \"%s\"", expected, string(actual))
		}
	})

	t.Run("should return the quiz stats successfully for configuration 3", func(t *testing.T) {

		// alice misses all question
		authCmd := NewAuthCmd()
		authCmdBuff := bytes.NewBufferString("")
		authCmd.SetOut(authCmdBuff)
		authCmd.SetArgs([]string{"--username", "alice", "--password", "rainbow"})
		authCmd.Execute()

		selectOptionsCommand := NewSelectOptionsCommand()
		selectOptionsBuff := bytes.NewBufferString("")
		selectOptionsCommand.SetOut(selectOptionsBuff)
		selectOptionsCommand.SetArgs([]string{"1:b,2:a,3:a"})
		selectOptionsCommand.Execute()

		// bob guesses one question
		authCmd = NewAuthCmd()
		authCmdBuff = bytes.NewBufferString("")
		authCmd.SetOut(authCmdBuff)
		authCmd.SetArgs([]string{"--username", "bob", "--password", "flower"})
		authCmd.Execute()

		selectOptionsCommand = NewSelectOptionsCommand()
		selectOptionsBuff = bytes.NewBufferString("")
		selectOptionsCommand.SetOut(selectOptionsBuff)
		selectOptionsCommand.SetArgs([]string{"1:a,2:a,3:a"})
		selectOptionsCommand.Execute()

		// eve gueses all questions
		authCmd = NewAuthCmd()
		authCmdBuff = bytes.NewBufferString("")
		authCmd.SetOut(authCmdBuff)
		authCmd.SetArgs([]string{"--username", "eve", "--password", "boat"})
		authCmd.Execute()

		selectOptionsCommand = NewSelectOptionsCommand()
		selectOptionsBuff = bytes.NewBufferString("")
		selectOptionsCommand.SetOut(selectOptionsBuff)
		selectOptionsCommand.SetArgs([]string{"1:a,2:b,3:b"})
		selectOptionsCommand.Execute()

		getStatsCommand := NewGetStatsCmd()
		getStatsBuff := bytes.NewBufferString("")
		getStatsCommand.SetOut(getStatsBuff)
		getStatsCommand.Execute()
		actual, err := ioutil.ReadAll(getStatsBuff)
		if err != nil {
			t.Fatal(err)
		}
		expected := "You were better than 100% of all quizzers"
		if string(actual) != expected {
			t.Fatalf("expected \"%s\" got \"%s\"", expected, string(actual))
		}
	})
}