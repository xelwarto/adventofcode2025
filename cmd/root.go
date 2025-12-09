package cmd

import (
	"code/cmd/days"
	"log"

	"github.com/spf13/cobra"
)

// Root Command
var rootCmd = &cobra.Command{
	Use:   "main.go",
	Short: "Advent of Code 2025",
}

func init() {
	// Load event day commands
	rootCmd.AddCommand(days.DayCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Error: %v", err)
	}
}
