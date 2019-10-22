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
  tm "github.com/buger/goterm"
  "github.com/eiannone/keyboard"
  "github.com/mitchellh/go-homedir"
  "github.com/spf13/cobra"
  "github.com/ttacon/chalk"
  "nort/config"
  "os"
)


var cfgFile string


// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
  Use:   "nort",
  Short: "",
  Long: ``,
  Run: NortsMenu,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
  if err := rootCmd.Execute(); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}

func init() {
  cobra.OnInitialize(initConfig)

  // Here you will define your flags and configuration settings.
  // Cobra supports persistent flags, which, if defined here,
  // will be global for your application.

  rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.nort.yaml)")


  // Cobra also supports local flags, which will only run
  // when this action is called directly.
  rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


// initConfig reads in config file and ENV variables if set.
func initConfig() {
  // Get the config file path
  if cfgFile == "" {
    // if not provided, use the home directory
    home, err := homedir.Dir()
    if err != nil {
      fmt.Println(err)
      os.Exit(1)
    }

    cfgFile = home + "/.nort.yaml"
  }

  // If the config file doesn't exist, create it
  if _, err := os.Stat(cfgFile); os.IsNotExist(err) {
    fmt.Println("Required config not found at:", cfgFile)
    fmt.Println("Creating initial file:", cfgFile)
  	if err := config.GenerateSampleConfigYaml(cfgFile); err != nil {
      fmt.Println(err.Error())
      panic("error creating the missing config file")
    }
  }

  // Once we know the config file exists, load it in
  var conf *config.Configurations
  conf, err := config.LoadConfigYaml(cfgFile)
  if err != nil {
    fmt.Println(err.Error())
    panic("error opening file")
  }

  // Store the config in viper
  config.StoreConfigInViper(conf)
}

func PrintMainMenu() {
  tm.Clear()
  box := tm.NewBox(30|tm.PCT, 20, 0)
  table := tm.NewTable(0, 10, 5, ' ', 0)
  fmt.Fprintf(table, "Enter command to begin \n\n")
  fmt.Fprintf(table, "F -- Find note (fuzzy search) \n")
  fmt.Fprintf(table, "A -- Add new note \n")
  fmt.Fprintf(table, "D -- Delete old note \n")
  fmt.Fprintf(table, "E -- Edit existing note \n")
  fmt.Fprintf(table, "Q -- Quit \n")
  //tm.Println(table)
  fmt.Fprint(box, table)
  tm.Print(tm.MoveTo(box.String(), 5|tm.PCT, 5|tm.PCT))
  //tm.Print(box)
  tm.Flush()
  fmt.Print("\n\n\n")
  PrintPrompt()
}

func PrintPrompt() {
  fmt.Print("> ")
}

func NortsMenu(_ *cobra.Command, _ []string) {
  PrintMainMenu()

  for {
    err := keyboard.Open()
    if err != nil {
      panic(err)
    }
    char, key, err := keyboard.GetKey()
    keyboard.Close()
    fmt.Printf("%c\n", char)


    if err != nil {
      panic(err)
    } else if key == keyboard.KeyEsc {
      tm.Clear()
      tm.Flush()
      break
    } else if char == 'F' || char == 'f' {
      fmt.Println("Find Note")
      FindNorts(nil, []string{})
    } else if char == 'A' || char == 'a' {
      fmt.Println("Add Note")
      fmt.Println(chalk.Red.String() + "Not Implemented" + chalk.Reset.String())
    } else if char == 'D' || char == 'd' {
      fmt.Println("Delete Note")
      fmt.Println(chalk.Red.String() + "Not Implemented" + chalk.Reset.String())
    } else if char == 'E' || char == 'e' {
      fmt.Println("Edit Note")
      fmt.Println(chalk.Red.String() + "Not Implemented" + chalk.Reset.String())
    } else if char == 'Q' || char == 'q' {
      fmt.Println("Quitting...")
      tm.Clear()
      tm.Flush()
      break
    } else {
      fmt.Printf(chalk.Red.String() + "Invalid selection: %q\n" + chalk.Reset.String(), char)
    }
    PrintPrompt()
  }
}
