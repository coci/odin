package cmd

import (
	"github.com/coci/odin/odin"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "initialize app",
	Long:  `initialize required files for creating blog.`,
	Run: func(cmd *cobra.Command, args []string) {
		odin.Init()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
