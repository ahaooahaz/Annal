package version

import (
	"github.com/AHAOAHA/Annal/binaries/version"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "version",
	Short: "get version",
	Long:  "get version",
	Run: func(cmd *cobra.Command, args []string) {
		version.PrintFullVersionInfo()
	},
}
