package cmd

import (
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "project-list",
	Short: "get a list of projects",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		name, err := os.UserHomeDir()
		if err != nil {
			log.Fatal(err.Error())
		}
		path := name + "/.healer/"
		files, err := ioutil.ReadDir(path)
		if err != nil {
			log.Fatal(err.Error())
		}
		for _, f := range files {
			cmd.Println(strings.ReplaceAll(f.Name(), ".json", ""))
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
