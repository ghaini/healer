package cmd

import (
	"github.com/spf13/cobra"
	"healer/storage"
	"log"
)

var projectCmd = &cobra.Command{
	Use:   "project-create",
	Short: "build a project that has a number of microservices",
	Long:  "",
	Run: runProjectCmd,
}

func init() {
	rootCmd.AddCommand(projectCmd)
	projectCmd.Flags().StringP("name", "n", "", "project name (required)")
	err := projectCmd.MarkFlagRequired("name")
	if err != nil {
		log.Fatal(err.Error())
	}
}

func runProjectCmd(cmd *cobra.Command, args []string) {
	projectName, err := cmd.Flags().GetString("name")
	if err != nil {
		log.Fatal(err.Error())
	}
	err = storage.CreateProject(projectName)
	if err != nil {
		log.Fatal(err.Error())
	}
}
