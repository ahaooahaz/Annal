package main

import (
	"github.com/AHAOAHA/Annal/internal/video"
	"github.com/spf13/cobra"
)

var videoCmd = &cobra.Command{
	Use:   "video",
	Short: "video tools",
	Long:  `video tools`,
}

func init() {
	rootCmd.AddCommand(videoCmd)
	videoCmd.AddCommand(video.GenCmd)
}
