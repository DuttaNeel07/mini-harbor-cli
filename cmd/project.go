package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// projectCmd represents the project command
var projectCmd = &cobra.Command{
	Use:   "project",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("project called")
	},
}
var projectListCmd = &cobra.Command{
	Use:   "list",
	Short: "List projects",
	Run: func(cmd *cobra.Command, args []string) {
		token := viper.GetString("token")
		if token == "" {
			fmt.Println("not logged in, please run: mini login --token <token>")
			return
		}
	
		req, err := http.NewRequest("GET", "https://api.github.com/user/repos", nil)
		if err != nil {
			fmt.Println("failed to create request:", err)
			return
		}
	
		req.Header.Set("Authorization", "Bearer "+token)
		req.Header.Set("Accept", "application/vnd.github+json")
	
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("API request failed:", err)
			return
		}
		defer resp.Body.Close()
	
		if resp.StatusCode != 200 {
			fmt.Println("API returned status:", resp.Status)
			return
		}
	
		var repos []map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&repos)
		if err != nil {
			fmt.Println("failed to parse response:", err)
			return
		}
	
		for _, repo := range repos {
			fmt.Println(repo["name"])
		}
	},
	
}

func init() {
	rootCmd.AddCommand(projectCmd)
	projectCmd.AddCommand(projectListCmd)
}
