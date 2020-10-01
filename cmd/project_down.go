package cmd

import (
	"healer/storage"
	"os/exec"

	"github.com/spf13/cobra"
)

var downCmd = &cobra.Command{
	Use:   "project-down",
	Short: "down project",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		projectName, _ := cmd.Flags().GetString("project-name")
		project, _ := storage.ReadProject(projectName)
		for _, command := range project.Down.Commands {
			cmd.Printf("execute %s command... \n", command)
			output, _ := exec.Command("/bin/bash", "-c", command).CombinedOutput()
			cmd.Println(string(output))
		}
	},
}

func init() {
	rootCmd.AddCommand(downCmd)
	downCmd.Flags().StringP("project-name", "p", "", "project name (required)")
	_ = downCmd.MarkFlagRequired("project-name")
}
