/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/ChikaKakazu/code-forge/internal/initialize"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a setup for code-forge",
	Long:  `Initialize a setup for code-forge. This will create a .code-forge directory in the current directory.`,
	Run: func(cmd *cobra.Command, args []string) {
		initialize.Initialize()

	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
