package todo

import "github.com/spf13/cobra"

var Cmd = &cobra.Command{
	Use:   "todo",
	Short: "todo task",
	Long:  `todo task`,
}

func init() {
	Cmd.AddCommand(listCmd)
	Cmd.AddCommand(createCmd)
}
