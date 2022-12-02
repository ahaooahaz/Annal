package main

import (
	"github.com/AHAOAHA/Annal/binaries/version"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "get version",
	Long:  "get version",
	Run: func(cmd *cobra.Command, args []string) {
		version.PrintFullVersionInfo()
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
