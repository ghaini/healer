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
	"os/exec"
)

// downCmd represents the down command
var downCmd = &cobra.Command{
	Use:   "project-down",
	Short: "down project",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		projectName, _ := cmd.Flags().GetString("project-name")
		name, _ := os.UserHomeDir()
		path := name + "/.healer/" + projectName + ".json"
		file, _ := ioutil.ReadFile(path)
		var project Project
		json.Unmarshal(file, &project)
		for _, command := range project.Down.Commands {
			cmd.Printf("execute %s command... \n", command)
			output, _ := exec.Command("/bin/bash","-c", command).CombinedOutput()
			cmd.Println(string(output))
		}
	},
}

func init() {
	rootCmd.AddCommand(downCmd)
	downCmd.Flags().StringP("project-name", "p", "", "project name (required)")
	downCmd.MarkFlagRequired("project-name")

}
