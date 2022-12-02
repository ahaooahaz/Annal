package jt

import (
	"fmt"
	"os"
	"syscall"

	"github.com/AHAOAHA/Annal/binaries/internal/global"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "jt",
	Short: "jumpto",
	Long:  `auto jump to remote machine by ssh.`,
	Run: func(cmd *cobra.Command, args []string) {
		command := fmt.Sprintf("%s/scripts/%s", global.ANNALROOT, "jt.sh")
		env := os.Environ()
		err := syscall.Exec(command, args, env)
		if err != nil {
			fmt.Printf("%v", err.Error())
			return
		}
	},
}
