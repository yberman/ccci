// Template for subcommands of ccci
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	someFlag = false
	someCmd  = &cobra.Command{
		Use:   "sample-sub-command",
		Short: "Idunno.",
		Long:  "This is a template for sub commands of ccci using cobra library.",
		Run:   runSubCommand,
	}
)

func init() {
	someCmd.Flags().BoolVarP(&someFlag, "someflag", "s", false, "Some flag for subcommand.")
	rootCmd.AddCommand(someCmd)
}


func runSubCommand(cmd *cobra.Command, args []string) {
	fmt.Println("executing subcommand")
}
