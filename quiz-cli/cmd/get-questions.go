package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
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

var getQuestionsCmd = &cobra.Command{
	Use:   "get-questions",
	Short: "Lists all the quiz questions",
	Long: `This command will display all questions and options`,
	Run: func(cmd *cobra.Command, args []string) {
		// HARD CODED - ideally this should be populated by the auth command
		auth = &Auth{APIKey: "key1"} 
		auth.getQuestionsRequest() 
	},
}

func (auth *Auth) getQuestionsRequest() {
	
	client := &http.Client{Timeout: 5 * time.Second}

	request, err := http.NewRequest("GET", "http://localhost:8081/questions", http.NoBody)
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
		
		var data []Question
		decoder := json.NewDecoder(response.Body)
		err = decoder.Decode(&data)
		if err != nil {
			log.Fatal(err)
		}
		for _, q := range data {
			fmt.Printf("%v: %v\n", q.ID, q.Description)
			for _, o := range q.Options {
				fmt.Printf("\t%v: %v\n", o.ID, o.Description)
			}
		}

	default:
		body, err := io.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%v", string(body))
	}
}

func init() {
	rootCmd.AddCommand(getQuestionsCmd)
}
