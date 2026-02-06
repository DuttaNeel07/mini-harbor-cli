package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"mini-harbor-cli/pkg/auth"

	"github.com/spf13/cobra"
)

var repo string

var auditCmd = &cobra.Command{
	Use:   "audit",
	Short: "Audit related commands",
}

var auditStreamCmd = &cobra.Command{
	Use:   "stream",
	Short: "Stream audit events",
	Run: func(cmd *cobra.Command, args []string) {
		token := auth.GetToken()
		if token == "" {
			fmt.Println("not logged in. run: mini login --token <token>")
			return
		}

		if repo == "" {
			fmt.Println("repo is required. example: --repo owner/name")
			return
		}

		url := fmt.Sprintf(
			"https://api.github.com/repos/%s/events",
			repo,
		)

		var lastEventID string

		for {
			req, err := http.NewRequest("GET", url, nil)
			if err != nil {
				fmt.Println("failed to create request:", err)
				return
			}

			req.Header.Set("Authorization", "Bearer "+token)
			req.Header.Set("Accept", "application/vnd.github+json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				fmt.Println("failed to fetch events:", err)
				return
			}

			if resp.StatusCode != 200 {
				fmt.Println("API returned status:", resp.Status)
				resp.Body.Close()
				return
			}

			var events []map[string]interface{}
			err = json.NewDecoder(resp.Body).Decode(&events)
			resp.Body.Close()
			if err != nil {
				fmt.Println("failed to parse events:", err)
				return
			}

			// Print only new events
			for i := len(events) - 1; i >= 0; i-- {
				event := events[i]
				id, _ := event["id"].(string)

				if id == lastEventID {
					break
				}

				eventType, _ := event["type"].(string)
				repoObj, _ := event["repo"].(map[string]interface{})
				repoName, _ := repoObj["name"].(string)

				fmt.Printf("[%s] %s\n", eventType, repoName)
			}

			if len(events) > 0 {
				lastEventID, _ = events[0]["id"].(string)
			}

			time.Sleep(10 * time.Second)
		}
	},
}

func init() {
	rootCmd.AddCommand(auditCmd)
	auditCmd.AddCommand(auditStreamCmd)

	auditStreamCmd.Flags().StringVar(
		&repo,
		"repo",
		"",
		"Repository in owner/name format",
	)
}
