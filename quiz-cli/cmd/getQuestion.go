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

func NewGetQuestionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "get-question <question-id>",
		Short: "Displays the question and options",
		Long: `This command will display a question and options provided the question ID`,
		Args: func(cmd *cobra.Command, args []string) error {
			if err := cobra.ExactArgs(1)(cmd, args); err != nil {
				return err
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			out := credentials.questionRequest(args[0])
			fmt.Fprint(cmd.OutOrStdout(), out)
		},
	}
}

func (auth *Auth) questionRequest(id string) string {

	client := &http.Client{Timeout: 5 * time.Second}

	request, err := http.NewRequest("GET", "http://localhost:8081/question?id="+id, http.NoBody)
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
		var data Question
		decoder := json.NewDecoder(response.Body)
		err = decoder.Decode(&data)
		if err != nil {
			log.Fatal(err)
		}
		return parseQuestion(data)
	default:
		body, err := io.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		return strings.TrimSpace(string(body))
	}
}

func parseQuestion(data Question) string {
	out := fmt.Sprintf("%v: %v\n", data.ID, data.Description)
	for _, o := range data.Options {
		out += fmt.Sprintf("\t%v: %v\n", o.ID, o.Description)
	}
	return out
}

func init() {
	getQuestionCmd := NewGetQuestionCmd()
	rootCmd.AddCommand(getQuestionCmd)
}
