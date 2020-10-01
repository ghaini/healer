package cmd

import (
	"github.com/spf13/cobra"
	"healer/storage"
)

var addCmdUp = &cobra.Command{
	Use:   "command-create-up",
	Short: "add a up command to project",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		projectName, _ := cmd.Flags().GetString("project-name")
		command, _ := cmd.Flags().GetString("command")
		project, _ := storage.ReadProject(projectName)
		project.Up.Commands = append(project.Up.Commands, command)
		_ = storage.SaveProject(projectName, project)
	},
}

func init() {
	rootCmd.AddCommand(addCmdUp)
	addCmdUp.Flags().StringP("project-name", "p", "", "project name (required)")
	addCmdUp.Flags().StringP("command", "c", "", "command (required)")
	_ = addCmdUp.MarkFlagRequired("project-name")
	_ = addCmdUp.MarkFlagRequired("command")
}
