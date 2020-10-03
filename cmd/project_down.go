package cmd

import (
	"healer/storage"
	"log"
	"os/exec"

	"github.com/spf13/cobra"
)

var downCmd = &cobra.Command{
	Use:   "project-down",
	Short: "down project",
	Long:  "",
	Run: runDownCmd,
}

func init() {
	rootCmd.AddCommand(downCmd)
	downCmd.Flags().StringP("project-name", "p", "", "project name (required)")
	_ = downCmd.MarkFlagRequired("project-name")
}

func runDownCmd(cmd *cobra.Command, args []string) {
	projectName, err := cmd.Flags().GetString("project-name")
	if err != nil {
		log.Fatal(err.Error())
	}
	project, err := storage.ReadProject(projectName)
	if err != nil {
		log.Fatal(err.Error())
	}
	for _, command := range project.Down.Commands {
		cmd.Printf("execute %s command... \n", command)
		output, err := exec.Command("/bin/bash", "-c", command).CombinedOutput()
		if err != nil {
			log.Fatal(err.Error())
		}
		cmd.Println(string(output))
	}
}