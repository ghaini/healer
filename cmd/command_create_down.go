package cmd

import (
	"healer/storage"

	"github.com/spf13/cobra"
)

var addCmdDown = &cobra.Command{
	Use:   "command-create-down",
	Short: "add a down command to project",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		projectName, _ := cmd.Flags().GetString("project-name")
		command, _ := cmd.Flags().GetString("command")
		project, _ := storage.ReadProject(projectName)
		project.Down.Commands = append(project.Down.Commands, command)
		_ = storage.SaveProject(projectName, project)
	},
}

func init() {
	rootCmd.AddCommand(addCmdDown)
	addCmdDown.Flags().StringP("project-name", "p", "", "project name (required)")
	addCmdDown.Flags().StringP("command", "c", "", "command (required)")
	_ = addCmdDown.MarkFlagRequired("project-name")
	_ = addCmdDown.MarkFlagRequired("command")
}
