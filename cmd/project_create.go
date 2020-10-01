package cmd

import (
	"github.com/spf13/cobra"
	"healer/storage"
)

var projectCmd = &cobra.Command{
	Use:   "project-create",
	Short: "build a project that has a number of microservices",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		projectName, _ := cmd.Flags().GetString("name")
		_ = storage.CreateProject(projectName)
	},
}

func init() {
	rootCmd.AddCommand(projectCmd)
	projectCmd.Flags().StringP("name", "n", "", "project name (required)")
	_ = projectCmd.MarkFlagRequired("name")
}
