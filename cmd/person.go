/*
This command is used to create a new person entity. It is the subcommand for add command.
The format for this command id family-tree add person <personName> -g <male/female>
*/
package cmd

import (
	createqueries "github.com/Hrit99/family-tree/createQueries"
	"github.com/spf13/cobra"
)

// personCmd represents the person command
var personCmd = &cobra.Command{
	Use:   "person",
	Short: "Command to create a person",
	Long: `This command is used to create a new person entity. 
	It is the subcommand for add command.
	Format for this command id family-tree add person <personName> -g <male/female>`,
	Run: func(cmd *cobra.Command, args []string) {
		fstatus, _ := cmd.Flags().GetString("gender")
		createqueries.AddPerson(args[0], fstatus)
	},
}

func init() {
	addCmd.AddCommand(personCmd)
	personCmd.Flags().StringP("gender", "g", "Unknown", "Add Floating Numbers")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// personCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// personCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
