package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

type Answer struct {
	QuestionID string `json:"question_id"`
	OptionID   string `json:"option_id"`
}

func NewSelectOptionsCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "select-options <question-id>:<option-id>,<question-id>:<option-id> ...",
		Short: "Selects a set of quiz answers",
		Long: `This command will post a set of quiz answers`,
		Args: func(cmd *cobra.Command, args []string) error {
			if err := cobra.ExactArgs(1)(cmd, args); err != nil {
				return err
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			var answers []Answer
			var answer Answer
			argsSlice := strings.Split(args[0], ",")
			for _, arg := range argsSlice {
				answerSlice := strings.Split(arg, ":")
				answer.QuestionID = answerSlice[0]
				answer.OptionID = answerSlice[1]
				answers = append(answers, answer)
			}
			out := credentials.selectOptionsRequest(answers)
			fmt.Fprint(cmd.OutOrStdout(), out)
		},
	}
}

func (auth *Auth) selectOptionsRequest(answers []Answer) string {

	client := &http.Client{Timeout: 5 * time.Second}
	var buf bytes.Buffer

    err := json.NewEncoder(&buf).Encode(answers)
    if err != nil {
		log.Fatal(err)
    }

	request, err := http.NewRequest("POST", "http://localhost:8081/answers", &buf)
	if err != nil {
		log.Fatal(err)
	}
	request.Header.Add("Accept", "application/json")
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-KEY", auth.APIKey)

	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	switch response.StatusCode {
	case http.StatusOK:
		var data []Answer
		decoder := json.NewDecoder(response.Body)
		err = decoder.Decode(&data)
		if err != nil {
			log.Fatal(err)
		}
		return parseSelectedOptions(data)
	default:
		body, err := io.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		return strings.TrimSpace(string(body))
	}
}

func parseSelectedOptions(data []Answer) string {
	var out string
	for _, a := range data {
		out += fmt.Sprintf("%v: %v\n", a.QuestionID, a.OptionID)
	}
	return out
}

func init() {
	selectOptionsCmd := NewSelectOptionsCommand()
	rootCmd.AddCommand(selectOptionsCmd)
}
