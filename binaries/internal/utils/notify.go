package utils

import (
	"fmt"
	"os/exec"

	"github.com/AHAOAHA/Annal/binaries/internal/config"
)

type Notification struct {
	Title       string
	Description string
}

func Notify(notice *Notification) {
	if notice == nil {
		return
	}

	args := []string{
		"-u",
		"critical",
		"-t",
		"0",
		"-a",
		config.PROJECT,
		"-i",
		config.ICONPATH,
		"title", // 8
		"desp",  // 9
	}

	args[8] = notice.Title
	args[9] = notice.Description

	cmd := exec.Command(config.NOTIFYSENDSH, args...)
	cmd.Run()
	out, _ := cmd.Output()
	fmt.Printf("%v\n", string(out))
}
