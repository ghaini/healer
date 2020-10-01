package cmd

import (
	"healer/storage"
	"os/exec"

	"github.com/spf13/cobra"
)

var upCmd = &cobra.Command{
	Use:   "project-up",
	Short: "up project",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		projectName, _ := cmd.Flags().GetString("project-name")
		project, _ := storage.ReadProject(projectName)
		for _, command := range project.Up.Commands {
			cmd.Printf("execute %s command... \n", command)
			output, _ := exec.Command("/bin/bash", "-c", command).CombinedOutput()
			cmd.Println(string(output))
		}
	},
}

func init() {
	rootCmd.AddCommand(upCmd)
	upCmd.Flags().StringP("project-name", "p", "", "project name (required)")
	_ = upCmd.MarkFlagRequired("project-name")
}
