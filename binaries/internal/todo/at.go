package todo

import (
	"bufio"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/AHAOAHA/Annal/binaries/internal/config"
	pb "github.com/AHAOAHA/Annal/binaries/pb/gen"
	"github.com/AHAOAHA/encapsutils"
	"github.com/sirupsen/logrus"
)

var (
	_AtTimeFormatString = "200601021504"
)

func notify(task *pb.TodoTask) (err error) {
	commandLine := []string{"#!/bin/bash"}
	jobpath := config.ATJOBS + "/" + task.GetUUID() + ".sh"
	step1 := []string{config.NOTIFYSENDSH, "-ti", task.GetTitle(), "-d", task.GetDescription(), "-t", "3"}
	step2 := []string{"rm", "-rf", jobpath}
	strings.Join(step1, " ")
	commandLine = append(commandLine, strings.Join(step1, " "))
	commandLine = append(commandLine, strings.Join(step2, " "))

	var f *os.File
	f, err = encapsutils.CreateFile(jobpath)
	if err != nil {
		logrus.Errorf(err.Error())
		return
	}
	defer f.Close()

	write := bufio.NewWriter(f)
	for _, c := range commandLine {
		_, err = write.WriteString(c)
		if err != nil {
			logrus.Errorf(err.Error())
			return
		}
		_, err = write.WriteString("\n")
		if err != nil {
			logrus.Errorf(err.Error())
			return
		}
	}
	write.Flush()

	attime := time.Unix(task.GetPlan(), 0).Format(_AtTimeFormatString)
	step3 := []string{"at", "-t", attime, "-f", jobpath}
	command3 := strings.Join(step3, " ")
	cmd := exec.Command("/bin/bash", "-c", command3)
	err = cmd.Run()
	out, _ := cmd.Output()
	if err != nil {
		logrus.Errorf("command3: %s, error: %d, output: %v", command3, err.Error(), string(out))
		return
	}
	return
}
