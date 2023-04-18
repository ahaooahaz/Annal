package webrtc

import (
	"github.com/AHAOAHA/Annal/binaries/cmd/annal/child/webrtc/publish"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "webrtc",
	Short: "webrtc",
	Long:  `webrtc kits.`,
}

func init() {
	Cmd.AddCommand(publish.Cmd)
}
