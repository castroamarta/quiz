package cmd

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Auth struct {
	Username string
	Password string
	APIKey string
}

var auth = &Auth{}

var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Command for basic authentication",
	PreRun: func(cmd *cobra.Command, args []string) {
		viper.BindPFlag("username", cmd.Flags().Lookup("username"))
		viper.BindPFlag("password", cmd.Flags().Lookup("password"))
   },
	Run: func(cmd *cobra.Command, args []string) {
		username, err := cmd.Flags().GetString("username")
		if err != nil {
			log.Fatal(err)
		}
		password, err := cmd.Flags().GetString("password")
		if err != nil {
			log.Fatal(err)
		}
		auth.Username = username
		auth.Password = password
		auth.authRequest()
	},
}

func (auth *Auth) authRequest() {

	client := &http.Client{Timeout: 5 * time.Second}

	request, err := http.NewRequest("GET", "http://localhost:8081/auth", http.NoBody)
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
	switch response.StatusCode {
	case http.StatusOK:
		auth.APIKey = strings.Join(strings.Fields(string(body)),"")
		fmt.Print("user authenticated successfully\n")
	default:
		fmt.Printf("%v", string(body))
	}
}

func init() {
	
	var username string
	var password string

	authCmd.Flags().StringVarP(&username, "username", "u", "", "Username (required if password is set)")
	authCmd.Flags().StringVarP(&password, "password", "p", "", "Password (required if username is set)")
	authCmd.MarkFlagRequired("username")
	authCmd.MarkFlagRequired("password")
	authCmd.MarkFlagsRequiredTogether("username", "password")
	rootCmd.AddCommand(authCmd)
}
