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
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
)

// projectCmd represents the project command
var projectCmd = &cobra.Command{
	Use:   "project",
	Short: "build a project that has a number of microservices",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		projectName, _ := cmd.Flags().GetString("name")
		name, _ := os.UserHomeDir()
		path := name + "/.healer/" + projectName + ".json"
		ioutil.WriteFile(path, []byte("{}"), 0777)
	},
}

func init() {
	rootCmd.AddCommand(projectCmd)
	projectCmd.Flags().StringP("name", "n", "", "project name (required)")
	projectCmd.MarkFlagRequired("name")
}
