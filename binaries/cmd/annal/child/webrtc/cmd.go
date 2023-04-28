package webrtc

import (
	"github.com/ahaooahaz/Annal/binaries/cmd/annal/child/webrtc/publish"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "webrtc",
	Short: "webrtc simple client tool kits",
	Long:  `webrtc simple client tool kits`,
}

func init() {
	Cmd.AddCommand(publish.Cmd)
}
