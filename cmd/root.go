package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "majora",
	Short: "Majora is a CLI auto grading tool",
	// Run: func(cmd *cobra.Command, args []string) {
	// 	// cmd here
	// },
}
