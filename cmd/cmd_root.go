package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "root",
	Short: "root.. short",
	Long:  "root.. long",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("you know how to execute")
	},
}

// Execute is the root cmd entry point
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
