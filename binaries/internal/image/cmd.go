package image

import "github.com/spf13/cobra"

var Cmd = &cobra.Command{
	Use:   "image",
	Short: "image utils",
	Long:  `image utils.`,
}

func init() {
	Cmd.AddCommand(uploadCmd)
}
