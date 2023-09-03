/*
This command is used to connect two different person with a relation. It uses subcommands such as "as" and "of" to specify relation and direction of the relation.
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// connectCmd represents the connect command
var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "'connect' command is used to initiate a relation.",
	Long: `This command is used to connect two different person with a relation. 
	It uses subcommands such as "as" and "of" to specify relation and direction of the relation.`,
	Run: func(cmd *cobra.Command, args []string) {
		// createqueries.ConnectRelation(getqueries.GetPersonWithName(args[0]))
		if args[1] == "as" {
			asCmd.Run(cmd, args)
		}
	},
}

func init() {
	rootCmd.AddCommand(connectCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// connectCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// connectCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
