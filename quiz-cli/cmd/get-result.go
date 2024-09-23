package cmd

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/spf13/cobra"
)

var getResultCmd = &cobra.Command{
	Use:   "get-result",
	Short: "Displays the quiz result",
	Long: `This command will display the total number of correct answers`,
	Run: func(cmd *cobra.Command, args []string) {
		auth = &Auth{APIKey: "key1"} 
		auth.getResultRequest()
	},
}

func (auth *Auth) getResultRequest() {
	client := &http.Client{Timeout: 5 * time.Second}

	request, err := http.NewRequest("GET", "http://localhost:8081/result", http.NoBody)
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
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v", string(body))
}

func init() {
	rootCmd.AddCommand(getResultCmd)
}
