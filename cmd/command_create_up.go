package cmd

import (
	"github.com/spf13/cobra"
	"healer/storage"
	"log"
)

var addCmdUp = &cobra.Command{
	Use:   "command-create-up",
	Short: "add a up command to project",
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
		project.Up.Commands = append(project.Up.Commands, command)
		err = storage.SaveProject(projectName, project)
		if err != nil {
			log.Fatal(err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmdUp)
	addCmdUp.Flags().StringP("project-name", "p", "", "project name (required)")
	addCmdUp.Flags().StringP("command", "c", "", "command (required)")
	_ = addCmdUp.MarkFlagRequired("project-name")
	_ = addCmdUp.MarkFlagRequired("command")
}
