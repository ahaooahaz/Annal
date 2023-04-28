package words

import (
	"github.com/ahaooahaz/Annal/binaries/cmd/annal/child/words/check"
	"github.com/ahaooahaz/Annal/binaries/cmd/annal/child/words/remb"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "words",
	Short: "remember words",
	Long:  `remember words`,
}

func init() {
	Cmd.AddCommand(check.Cmd)
	Cmd.AddCommand(remb.Cmd)
}
