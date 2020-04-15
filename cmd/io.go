package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"strings"
)

func init() {
	rootCmd.AddCommand(ioCmd)
}

var ioCmd = &cobra.Command{
	Use:   "io",
	Short: "The io example",
	Long:  `The io examples`,
	Run: func(cmd *cobra.Command, args []string) {
		raw, err := ioutil.ReadFile("/etc/shells")
		if err != nil {
			log.Fatal(err)
		}

		var shells []string
		for _, line := range strings.Split(string(raw), "\n") {
			shells = append(shells, line)
		}

		filename := "hoge"
		if err := ioutil.WriteFile(filename, []byte(strings.Join(shells, "\r\n")), 0666); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("successfully written into %s\n", filename)
	},
}
