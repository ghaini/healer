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
	"encoding/json"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
)

type Project struct {
	Commands []string `json:"commands"`
}

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add commands to run the project",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		projectName, _ := cmd.Flags().GetString("project-name")
		command, _ := cmd.Flags().GetString("command")
		name, _ := os.UserHomeDir()
		path := name + "/.healer/" + projectName + ".json"
		file, _ := ioutil.ReadFile(path)
		var project Project
		json.Unmarshal(file, &project)
		project.Commands = append(project.Commands, command)
		jsonString, _ := json.Marshal(project)
		_ = ioutil.WriteFile(path, jsonString, 0777)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringP("project-name", "p", "", "project name (required)")
	addCmd.Flags().StringP("command", "c", "", "command (required)")
	addCmd.MarkFlagRequired("project-name")
	addCmd.MarkFlagRequired("command")
}
