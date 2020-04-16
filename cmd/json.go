package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func init() {
	rootCmd.AddCommand(jsonCmd)
}

type Repository struct {
	Name  string
	Owner struct {
		Name string `json:"login"`
	}
	Language string
	URL      string `json:"html_url"`
}

var jsonCmd = &cobra.Command{
	Use:   "json",
	Short: "json package examples",
	Long:  `json package examples`,
	Run: func(cmd *cobra.Command, args []string) {
		decoder := json.NewDecoder(os.Stdin)
		var repos []Repository
		if err := decoder.Decode(&repos); err != nil {
			fmt.Println(err)
		}
		fmt.Println(repos[0].Name)
		fmt.Println(repos[0].Owner.Name)
		for _, repo := range repos {
			fmt.Println(repo.Name)
		}
	},
}
