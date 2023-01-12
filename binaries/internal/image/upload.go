package image

import (
	"context"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/AHAOAHA/encapsutils"
	"github.com/spf13/cobra"
)

var (
	debug = false
)

var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "upload image to imagehosting",
	Long:  `upload image to imagehosting`,
	Run:   uploadImageToImageHosting,
}

func init() {
	uploadCmd.Flags().StringP("path", "p", "", "image path or dir, support recursive upload with dir")
	uploadCmd.Flags().StringP("token", "t", "", "github access token")
	uploadCmd.Flags().StringP("repo", "r", "", "image hosting repo, eg: AHAOAHA/ImageHosting")
	uploadCmd.Flags().StringP("branch", "b", "main", "image hosting repo branch")
	uploadCmd.Flags().BoolP("debug", "d", false, "debug switch")

	uploadCmd.MarkFlagRequired("path")
	uploadCmd.MarkFlagRequired("token")
	uploadCmd.MarkFlagRequired("hosting")
}

func uploadImageToImageHosting(cmd *cobra.Command, args []string) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var imagepath, token, branch, repo string
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

	branch, err = cmd.Flags().GetString("branch")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err.Error())
		return
	}

	repo, err = cmd.Flags().GetString("repo")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err.Error())
		return
	}

	debug, err = cmd.Flags().GetBool("debug")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err.Error())
		return
	}

	if encapsutils.IsFile(imagepath) {
		var URL string
		URL, err = GithubPutImage(ctx, repo, branch, token, imagepath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s@%s\n", filepath.Base(imagepath), err.Error())
			return
		}

		fmt.Printf("%s@%s", filepath.Base(imagepath), URL)
		return
	} else if encapsutils.IsDir(imagepath) {
		err = filepath.Walk(imagepath, func(path string, info fs.FileInfo, ein error) (ineo error) {
			var ine error
			if ein != nil {
				return
			}
			if !info.IsDir() {
				var URL string
				URL, ine = GithubPutImage(ctx, repo, branch, token, path)
				if ine != nil {
					fmt.Fprintf(os.Stderr, "%s@%s\n", filepath.Base(path), ine.Error())
					return
				}

				fmt.Printf("%s@%s\n", filepath.Base(path), URL)
			}
			return
		})
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		}
	}

}
