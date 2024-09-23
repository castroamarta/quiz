package cmd

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/spf13/cobra"
)

var getStatsCmd = &cobra.Command{
	Use:   "get-stats",
	Short: "Displays the quiz statistics",
	Long: `This command will display the percentage of correct anwsers when compared to other quizzers`,
	Run: func(cmd *cobra.Command, args []string) {
		// HARD CODED - ideally this should be populated by the auth command
		auth = &Auth{APIKey: "key1"} 
		auth.getStatsRequest()
	},
}

func (auth *Auth) getStatsRequest() {

	client := &http.Client{Timeout: 5 * time.Second}

	request, err := http.NewRequest("GET", "http://localhost:8081/stats", http.NoBody)
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
	rootCmd.AddCommand(getStatsCmd)
}
