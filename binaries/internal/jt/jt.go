package jt

import (
	"fmt"
	"os"
	"syscall"

	"github.com/AHAOAHA/Annal/binaries/internal/global"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "jt",
	Short: "jumpto",
	Long:  `auto jump to remote machine by ssh.`,
	Run: func(cmd *cobra.Command, args []string) {
		command := fmt.Sprintf("%s/scripts/%s", global.ANNALROOT, "jt.sh")
		logger := logrus.WithFields(logrus.Fields{
			"command": command,
			"args":    args,
		})
		newArgs := []string{"-c", command}
		newArgs = append(newArgs, args...)
		env := os.Environ()
		logger.Infof("running")
		err := syscall.Exec("/bin/bash", newArgs, env)
		if err != nil {
			logger.Errorf("exec failed, err: %v", err.Error())
			fmt.Printf("%v", err.Error())
			return
		}
	},
}
