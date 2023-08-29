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

	createqueries "github.com/Hrit99/family-tree/createQueries"
	getqueries "github.com/Hrit99/family-tree/getQueries"
	"github.com/spf13/cobra"
)

// ofCmd represents the of command
var ofCmd = &cobra.Command{
	Use:   "of",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("of called")
		createqueries.ConnectRelation(getqueries.GetPersonWithName(args[0]), getqueries.GetPersonWithName(args[4]), getqueries.GetRelation(args[2]))
	},
}

func init() {
	asCmd.AddCommand(ofCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ofCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ofCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
