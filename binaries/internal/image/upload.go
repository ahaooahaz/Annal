package image

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var uploadCmd = &cobra.Command{
	Use:     "upload",
	Aliases: []string{"u"},
	Short:   "upload image to imagehosting",
	Long:    `upload image to imagehosting`,
	Run:     uploadImageToImageHostinig,
}

func init() {
	uploadCmd.Flags().StringP("path", "p", "", "image path")
	uploadCmd.Flags().StringP("token", "t", "", "github access token")

	uploadCmd.MarkFlagRequired("path")
	uploadCmd.MarkFlagRequired("token")
}

func uploadImageToImageHostinig(cmd *cobra.Command, args []string) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var imagepath, token string
	var err error
	imagepath, err = cmd.Flags().GetString("path")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err.Error())
		return
	}

	token, err = cmd.Flags().GetString("token")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err.Error())
		return
	}

	err = GithubPutImage(ctx, token, imagepath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		return
	}
}
