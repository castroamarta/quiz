package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

type Option struct {
	ID        string
	Description string
}

type Question struct {
	ID string
	Description string
	Options     []Option
}

func NewGetQuestionsCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "get-questions",
		Short: "Lists all the quiz questions",
		Long: `This command will display all questions and options`,
		Run: func(cmd *cobra.Command, args []string) {
			out := credentials.questionsRequest()
			fmt.Fprint(cmd.OutOrStdout(), out)
		},
	}
}

func (auth *Auth) questionsRequest() string {
	
	client := &http.Client{Timeout: 5 * time.Second}

	request, err := http.NewRequest("GET", "http://localhost:8081/questions", http.NoBody)
	if err != nil {
		log.Fatal(err)
	}
	request.Header.Add("Accept", "application/json")
	request.Header.Add("Content-Type", "application/json")
	request.SetBasicAuth(auth.Username, auth.Password)

	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	switch response.StatusCode {
	case http.StatusOK:
		var data []Question
		decoder := json.NewDecoder(response.Body)
		err = decoder.Decode(&data)
		if err != nil {
			log.Fatal(err)
		}
		return parseQuestions(data)
	default:
		body, err := io.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		return strings.TrimSpace(string(body))
	}
}

func parseQuestions(data []Question) string {
	var out string
	for _, q := range data {
		out += fmt.Sprintf("%v: %v\n", q.ID, q.Description)
		for _, o := range q.Options {
			out += fmt.Sprintf("\t%v: %v\n", o.ID, o.Description)
		}
	}
	return out
}

func init() {
	getQuestionsCmd := NewGetQuestionsCmd()
	rootCmd.AddCommand(getQuestionsCmd)
}
