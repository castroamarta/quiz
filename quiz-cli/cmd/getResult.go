package cmd

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/spf13/cobra"
)


func NewGetResultCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "get-result",
		Short: "Displays the quiz result",
		Long: `This command will display the total number of correct answers`,
		Run: func(cmd *cobra.Command, args []string) {
			out := credentials.resultRequest()
			fmt.Fprint(cmd.OutOrStdout(), out)
		},
	}
}

func (auth *Auth) resultRequest() string {
	client := &http.Client{Timeout: 5 * time.Second}

	request, err := http.NewRequest("GET", "http://localhost:8081/result", http.NoBody)
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
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	return strings.TrimSpace(string(body))
}

func init() {
	getResultCmd := NewGetResultCmd()
	rootCmd.AddCommand(getResultCmd)
}
