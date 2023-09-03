/*
This command specifies the type of relation between two given person. The syntax goes like <Person1> as <Relation> of <Person2>.
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// asCmd represents the as command
var asCmd = &cobra.Command{
	Use:   "as",
	Short: "'as' command is used to specify the type of relation.",
	Long: `This command specifies the type of relation between two given person. 
	The syntax goes like <Person1> as <Relation> of <Person2>.`,
	Run: func(cmd *cobra.Command, args []string) {
		if args[3] == "of" {
			ofCmd.Run(cmd, args)
		}
	},
}

func init() {
	connectCmd.AddCommand(asCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// asCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// asCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
