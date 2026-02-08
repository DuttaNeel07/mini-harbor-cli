package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)
var token string

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if token == ""{
			fmt.Println("Token is required")
			return
		}
		viper.Set("token", token)
		err := viper.WriteConfig()

		if err != nil{
			fmt.Println("Failed to save token: ", err)
			return
		}
		fmt.Println("Token Saved Sucessfully");

			req, err := http.NewRequest("GET", "https://api.github.com/user", nil)
	if err != nil {
		fmt.Println("Token saved, but failed to verify user")
		return
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Accept", "application/vnd.github+json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Token saved, but failed to verify user")
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Println("Token saved, but failed to verify user")
		return
	}

	var user map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&user)
	if err != nil {
		fmt.Println("Token saved, but failed to parse user info")
		return
	}

	fmt.Println("Logged in as:", user["login"])
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
	loginCmd.Flags().StringVar(&token, "token", "", "Authentication token")
}
