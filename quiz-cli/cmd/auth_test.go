package cmd

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func Test_ExecuteAuthCommand(t *testing.T) {
	t.Run("should return authentication successfull message", func(t *testing.T) {
		cmd := NewAuthCmd()
		b := bytes.NewBufferString("")
		cmd.SetOut(b)
		cmd.SetArgs([]string{"--username", "alice", "--password", "rainbow"})
		cmd.Execute()
		actual, err := ioutil.ReadAll(b)
		if err != nil {
			t.Fatal(err)
		}
		expected := "user authenticated successfully"
		if string(actual) != expected {
			t.Fatalf("expected \"%s\" got \"%s\"", expected, string(actual))
		}
	})

	t.Run("should return authentication failed message", func(t *testing.T) {
		cmd := NewAuthCmd()
		b := bytes.NewBufferString("")
		cmd.SetOut(b)
		cmd.SetArgs([]string{"--username", "alice", "--password", "seasalt"})
		cmd.Execute()
		actual, err := ioutil.ReadAll(b)
		if err != nil {
			t.Fatal(err)
		}
		expected := "user authentication failed"
		if string(actual) != expected {
			t.Fatalf("expected \"%s\" got \"%s\"", expected, string(actual))
		}
	})
}