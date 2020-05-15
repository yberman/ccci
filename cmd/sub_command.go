// Template for subcommands
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	someCmd.Flags().BoolVarP(&useBeer, "someflag", "s", false, "Some flag for subcommand.")
	rootCmd.AddCommand(someCmd)
}

var (
	someFlag = false
	someCmd  = &cobra.Command{
		Use:   "sample-sub-command",
		Short: "Idunno.",
		Long:  "You be you.",
		Run:   runSubCommand,
	}
)

func runSubCommand(cmd *cobra.Command, args []string) {
	fmt.Println("execute subcommand")
}
