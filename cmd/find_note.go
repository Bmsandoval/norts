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
	"nort/config"
	"nort/services"
)

var nortsFindCmd = &cobra.Command{
	Use:   "find",
	Aliases: []string{"f"},
	Short: "Find files and folders in notes location",
	Long: ``,
	Run: FindNote,
}

func FindNote(_ *cobra.Command, _ []string) {
	configs := config.GetConfigFromViper()

	if configs.Nort.Path == "" {
		fmt.Println("whoops, path is empty")
		return
	}

	filePath, err := services.SelectFileInPath(configs.Nort.Path)
	if err != nil {
		fmt.Println(err.Error())
		return
	} else if filePath == "" {
		fmt.Println("no file selected")
		return
	}

	if configs.Nort.Editor == "" {
		fmt.Println("no editor configured. edit nort's config and add an editor")
		return
	}
	// Check for necessary stuff
	if ! services.ExecutableExists(configs.Nort.Editor) {
		fmt.Printf("editor %q not installed", configs.Nort.Editor)
		return
	}

	if err := services.ExecuteOnFile(configs.Nort.Editor, filePath); err != nil {
		fmt.Println(err.Error())
		return
	}
}

func init() {
	rootCmd.AddCommand(nortsFindCmd)
}
