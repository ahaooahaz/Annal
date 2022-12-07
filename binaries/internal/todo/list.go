package todo

import "github.com/spf13/cobra"

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l"},
	Short:   "list tasks",
	Long:    `list tasks`,
	Run: func(cmd *cobra.Command, args []string) {
		print("TODO")
	},
}
