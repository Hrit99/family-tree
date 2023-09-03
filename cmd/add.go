/*
Copyright © 2023 Hritik Pihal <hritikpihal51299@gmail.com>

add command is used after family-tree to add a new person and store it. It is also used to add a new relation and store it.
The format goes like family-tree add <sub-Command> <name>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "This command is used to add a new person or new relation.",
	Long: `add command is used after family-tree to add a new person and store it. 
	It is also used to add a new relation and store it. 
	The format goes like family-tree add <sub-Command> <name>`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
