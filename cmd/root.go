package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "ccci",
		Short: "Cracking the Coding Interview",
		Long: `Solutions to some of the problems in Cracking the Coding Interview.
					The extra "c" is for quality.`,
		Run: func(cmd *cobra.Command, args []string) {
		},
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
