package check

import "github.com/spf13/cobra"

var Cmd = &cobra.Command{
	Use:   "check",
	Short: "check words",
	Long:  `check words`,
	Run:   func(cmd *cobra.Command, args []string) {},
}
