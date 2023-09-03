/*
This command is used to count the number of relatives with the specified relation of a person.
The format for the command is family-tree count <relation> of <person>
*/
package cmd

import (
	getqueries "github.com/Hrit99/family-tree/getQueries"
	"github.com/Hrit99/family-tree/queries"
	"github.com/spf13/cobra"
)

// countCmd represents the count command
var countCmd = &cobra.Command{
	Use:   "count",
	Short: "Used to count number of specific relatives",
	Long: `This command is used to count the number of relatives with the specified relation of a person.
	The format for the command is family-tree count <relation> of <person>`,
	Run: func(cmd *cobra.Command, args []string) {
		if args[1] == "of" {
			queries.CountQuery(*getqueries.GetPersonWithName(args[2]), *getqueries.GetRelation(args[0]))
		}
	},
}

func init() {
	rootCmd.AddCommand(countCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// countCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// countCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
