/*
This command is used to get all the relatives with the specified relation for a person.
The syntax for the command is as follows family-tree get <relation> of <Person>
*/
package cmd

import (
	getqueries "github.com/Hrit99/family-tree/getQueries"
	"github.com/Hrit99/family-tree/queries"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Command to get the relatives",
	Long: `This command is used to get all the relatives with the specified relation for a person.
	The syntax for the command is as follows family-tree get <relation> of <Person>`,
	Run: func(cmd *cobra.Command, args []string) {
		if args[1] == "of" {
			queries.GetRelatives(*getqueries.GetRelation(args[0]), *getqueries.GetPersonWithName(args[2]))
		}
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
