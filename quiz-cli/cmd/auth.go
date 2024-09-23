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

type Auth struct {
	Username string
	Password string
}

var credentials = &Auth{} 

func NewAuthCmd() *cobra.Command {

	var username string
	var password string

	authCmd := &cobra.Command{
		Use:   "auth",
		Short: "Command for basic authentication",
		Run: func(cmd *cobra.Command, args []string) {
			username, err := cmd.Flags().GetString("username")
			if err != nil {
				log.Fatal(err)
			}
			password, err := cmd.Flags().GetString("password")
			if err != nil {
				log.Fatal(err)
			}
			credentials.Username = username
			credentials.Password = password
			out := credentials.authRequest()
			fmt.Fprint(cmd.OutOrStdout(), out)
		},
	}
	authCmd.Flags().StringVarP(&username, "username", "u", "", "Username (required if password is set)")
	authCmd.Flags().StringVarP(&password, "password", "p", "", "Password (required if username is set)")
	authCmd.MarkFlagRequired("username")
	authCmd.MarkFlagRequired("password")
	authCmd.MarkFlagsRequiredTogether("username", "password")
	return authCmd
}

func (auth *Auth) authRequest() string {

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
		return "user authenticated successfully"
	default:
		return strings.TrimSpace(string(body))
	}
}

func init() {
	authCmd := NewAuthCmd()
	rootCmd.AddCommand(authCmd)
}
