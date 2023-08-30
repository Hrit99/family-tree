/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

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
	"os"

	filefunctions "github.com/Hrit99/family-tree/fileFunctions"
	globalvar "github.com/Hrit99/family-tree/globalVar"
	"github.com/Hrit99/family-tree/models"
	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "family-tree",
	Short: "root command for family-tree",
	Long:  `This is the root command for family-tree CLI`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
	Run: func(cmd *cobra.Command, args []string) {
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	initHash()
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.family-tree.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			os.Exit(1)
		}

		// Search config in home directory with name ".family-tree" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".family-tree")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func initHash() {
	globalvar.PersonHashMap = make(map[string][]models.Person)
	persons := filefunctions.ReadFilePerson()
	for _, person := range persons {
		if globalvar.PersonHashMap[person.Name] == nil {
			listOfPerson := make([]models.Person, 0, 1)
			listOfPerson = append(listOfPerson, person)
			globalvar.PersonHashMap[person.Name] = listOfPerson
		} else {
			listOfPerson := globalvar.PersonHashMap[person.Name]
			listOfPerson = append(listOfPerson, person)
			globalvar.PersonHashMap[person.Name] = listOfPerson
		}
	}

	globalvar.RelationTypeLists = make(map[string]models.Relation)
	relations := filefunctions.ReadFileRelation()
	for _, relation := range relations {

		globalvar.RelationTypeLists[relation.Name] = relation

	}
}
