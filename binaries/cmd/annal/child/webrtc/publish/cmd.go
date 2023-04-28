package publish

import (
	"github.com/ahaooahaz/Annal/binaries/cmd/annal/child/webrtc/publish/zlmediakit"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "publish",
	Short: "publish video to remote webrtc server",
	Long:  `publish video to remote webrtc server`,
}

func init() {
	Cmd.AddCommand(zlmediakit.Cmd)
}
