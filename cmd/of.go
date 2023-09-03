/*
This command is used to declare the direction of the relation defined between two people. It is subcommand for commands like connect and get.
The syntax for the command is family-tree <pre-command> ... <relation> of <Person>
*/
package cmd

import (
	createqueries "github.com/Hrit99/family-tree/createQueries"
	getqueries "github.com/Hrit99/family-tree/getQueries"
	"github.com/spf13/cobra"
)

// ofCmd represents the of command
var ofCmd = &cobra.Command{
	Use:   "of",
	Short: "Command to specify the direction of relation.",
	Long: `This command is used to declare the direction of the relation defined between two people.
	It is subcommand for commands like connect and get.
	The syntax for the command is family-tree <pre-command> ... <relation> of <Person>`,
	Run: func(cmd *cobra.Command, args []string) {
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
