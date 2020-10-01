package cmd

import (
	"io/ioutil"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "project-list",
	Short: "get a list of projects",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := os.UserHomeDir()
		path := name + "/.healer/"
		files, _ := ioutil.ReadDir(path)

		for _, f := range files {
			cmd.Println(strings.ReplaceAll(f.Name(), ".json", ""))
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
