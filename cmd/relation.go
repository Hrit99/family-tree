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
	createqueries "github.com/Hrit99/family-tree/createQueries"
	"github.com/spf13/cobra"
)

// relationCmd represents the relation command
var relationCmd = &cobra.Command{
	Use:   "relation",
	Short: "Command to create a relation",
	Long: `This command is used to create a new relation entity. 
	It is the subcommand for add command.
	Format for this command id family-tree add relation <personName> -g <male/female>`,
	Run: func(cmd *cobra.Command, args []string) {
		fstatus, _ := cmd.Flags().GetString("gender")
		createqueries.AddRelation(args[0], fstatus)
	},
}

func init() {
	addCmd.AddCommand(relationCmd)
	relationCmd.Flags().StringP("gender", "g", "Unknown", "Add Floating Numbers")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// relationCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// relationCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
