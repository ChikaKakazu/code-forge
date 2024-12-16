/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/ChikaKakazu/code-forge/internal/generate"
	"github.com/spf13/cobra"
)

var fileName string

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a new file from template file",
	Long:  `Generate a new file from template file`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// 引数がない場合はエラーを返す
		n := args[0]

		generate.GenerateFile(n)
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

}
