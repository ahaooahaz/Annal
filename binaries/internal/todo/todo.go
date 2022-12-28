package todo

import "github.com/spf13/cobra"

var Cmd = &cobra.Command{
	Use:   "todo",
	Short: "todo task",
	Long:  `todo task`,
}

var (
	_TimeFormatString = "2006-01-02 15:04:05"
)

func init() {
	Cmd.AddCommand(listCmd)
	Cmd.AddCommand(createCmd)
	Cmd.AddCommand(doneCmd)
	Cmd.AddCommand(pruneCmd)
}
