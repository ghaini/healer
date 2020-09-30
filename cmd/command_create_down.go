/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"github.com/spf13/cobra"
	"healer/storage"
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
