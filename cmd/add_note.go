/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

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
	"fmt"
	"github.com/spf13/cobra"
	"norts/config"
	"norts/services"
)

var nortsAddCmd = &cobra.Command{
	Use:   "add",
	Aliases: []string{"a"},
	Short: "Add new notes within notes location",
	Long: ``,
	Run: AddNote,
}

func AddNote(_ *cobra.Command, _ []string) {
	configs := config.GetConfigFromViper()

	if configs.Norts.Path == "" {
		fmt.Println("whoops, path is empty")
		return
	}

	filePath, err := services.SelectFolderInPath(configs.Norts.Path)
	if err != nil {
		fmt.Println(err.Error())
		return
	} else if filePath == "" {
		fmt.Println("no folder selected")
		return
	}

}

func init() {
	rootCmd.AddCommand(nortsAddCmd)
}
