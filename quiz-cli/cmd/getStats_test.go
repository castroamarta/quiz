package cmd

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func Test_ExecuteGetStatsCommand(t *testing.T) {
	t.Run("should return the quiz stats successfully for configuration 1", func(t *testing.T) {
		credentials = &Auth{APIKey: "VAFJWEKSFS"} // alice guesses one option
		selectOptionsCommand := NewSelectOptionsCommand()
		selectOptionsBuff := bytes.NewBufferString("")
		selectOptionsCommand.SetOut(selectOptionsBuff)
		selectOptionsCommand.SetArgs([]string{"1:a,2:a,3:a"})
		selectOptionsCommand.Execute()

		credentials = &Auth{APIKey: "FEJRGIERGJ"} // bob misses all options
		selectOptionsCommand = NewSelectOptionsCommand()
		selectOptionsBuff = bytes.NewBufferString("")
		selectOptionsCommand.SetOut(selectOptionsBuff)
		selectOptionsCommand.SetArgs([]string{"1:b,2:a,3:a"})
		selectOptionsCommand.Execute()

		credentials = &Auth{APIKey: "PQIENFJRGR"} // eve misses all options
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
		credentials = &Auth{APIKey: "VAFJWEKSFS"} // alice guesses one option
		selectOptionsCommand := NewSelectOptionsCommand()
		selectOptionsBuff := bytes.NewBufferString("")
		selectOptionsCommand.SetOut(selectOptionsBuff)
		selectOptionsCommand.SetArgs([]string{"1:a,2:a,3:a"})
		selectOptionsCommand.Execute()

		credentials = &Auth{APIKey: "FEJRGIERGJ"} // bob misses all options
		selectOptionsCommand = NewSelectOptionsCommand()
		selectOptionsBuff = bytes.NewBufferString("")
		selectOptionsCommand.SetOut(selectOptionsBuff)
		selectOptionsCommand.SetArgs([]string{"1:b,2:a,3:a"})
		selectOptionsCommand.Execute()

		credentials = &Auth{APIKey: "PQIENFJRGR"} // eve guesses two options
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
		expected := "You were better than 100% of all quizzers"
		if string(actual) != expected {
			t.Fatalf("expected \"%s\" got \"%s\"", expected, string(actual))
		}
	})

	// t.Run("should return the quiz stats successfully for configuration 3", func(t *testing.T) {
	// 	credentials = &Auth{APIKey: "1:a,2:b,3:b"} // alice guesses all options
	// 	selectOptionsCommand := NewSelectOptionsCommand()
	// 	selectOptionsBuff := bytes.NewBufferString("")
	// 	selectOptionsCommand.SetOut(selectOptionsBuff)
	// 	selectOptionsCommand.SetArgs([]string{"1:a,2:b,3:b"})
	// 	selectOptionsCommand.Execute()

	// 	credentials = &Auth{APIKey: "FEJRGIERGJ"} // bob guesses one option
	// 	selectOptionsCommand = NewSelectOptionsCommand()
	// 	selectOptionsBuff = bytes.NewBufferString("")
	// 	selectOptionsCommand.SetOut(selectOptionsBuff)
	// 	selectOptionsCommand.SetArgs([]string{"1:a,2:a,3:a"})
	// 	selectOptionsCommand.Execute()

	// 	credentials = &Auth{APIKey: "PQIENFJRGR"} // eve guesses two options
	// 	selectOptionsCommand = NewSelectOptionsCommand()
	// 	selectOptionsBuff = bytes.NewBufferString("")
	// 	selectOptionsCommand.SetOut(selectOptionsBuff)
	// 	selectOptionsCommand.SetArgs([]string{"1:a,2:b,3:a"})
	// 	selectOptionsCommand.Execute()

	// 	getStatsCommand := NewGetStatsCmd()
	// 	getStatsBuff := bytes.NewBufferString("")
	// 	getStatsCommand.SetOut(getStatsBuff)
	// 	getStatsCommand.Execute()
	// 	actual, err := ioutil.ReadAll(getStatsBuff)
	// 	if err != nil {
	// 		t.Fatal(err)
	// 	}
	// 	expected := "You were better than 67% of all quizzers"
	// 	if string(actual) != expected {
	// 		t.Fatalf("expected \"%s\" got \"%s\"", expected, string(actual))
	// 	}
	// })
}