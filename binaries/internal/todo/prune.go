package todo

import "github.com/spf13/cobra"

var pruneCmd = &cobra.Command{
	Use:   "prune",
	Short: "prune task",
	Long:  `prune task`,
	Run:   pruneTodoTask,
}

func pruneTodoTask(cmd *cobra.Command, args []string) {

}
