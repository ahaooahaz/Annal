package remb

import (
	"fmt"

	"github.com/ahaooahaz/Annal/binaries/config"
	"github.com/ahaooahaz/encapsutils"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "remb",
	Short: "remember words",
	Long:  `remember words`,
	Run: func(cmd *cobra.Command, args []string) {
		text, err := encapsutils.RandomLineFromFile(fmt.Sprintf("%s%s", config.ANNALROOT, "/configs/CET4.csv"))
		if err != nil {
			fmt.Print(err.Error())
			return
		}

		fmt.Println(text)
	},
}
