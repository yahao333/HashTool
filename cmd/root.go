package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "hashtool",
	Short: "A comprehensive hash utility",
	Long: `HashTool is a versatile hash utility that provides both command-line interface (CLI) 
and web service capabilities. This tool supports multiple hash algorithms and offers 
an easy-to-use interface for generating and verifying hash values.`,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}