package cmd

import (
	"healer/storage"
	"log"

	"github.com/spf13/cobra"
)

var addCmdDown = &cobra.Command{
	Use:   "command-create-down",
	Short: "add a down command to project",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		projectName, err := cmd.Flags().GetString("project-name")
		if err != nil {
			log.Fatal(err.Error())
		}
		command, err := cmd.Flags().GetString("command")
		if err != nil {
			log.Fatal(err.Error())
		}
		project, err := storage.ReadProject(projectName)
		if err != nil {
			log.Fatal(err.Error())
		}
		project.Down.Commands = append(project.Down.Commands, command)
		err = storage.SaveProject(projectName, project)
		if err != nil {
			log.Fatal(err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmdDown)
	addCmdDown.Flags().StringP("project-name", "p", "", "project name (required)")
	addCmdDown.Flags().StringP("command", "c", "", "command (required)")
	_ = addCmdDown.MarkFlagRequired("project-name")
	_ = addCmdDown.MarkFlagRequired("command")
}
