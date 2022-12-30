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
	_AtTimeFormatString = "15:04 2006-01-02"
)

func CreateAtJob(task *pb.TodoTask) (err error) {
	commandLine := []string{"#!/bin/bash"}
	jobpath := config.ATJOBS + task.GetUUID() + ".sh"
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
	}

	attime := time.Unix(task.GetPlan(), 0).Format(_AtTimeFormatString)
	step3 := []string{"at", "-t", attime, "-f", jobpath}

	cmd := exec.Command("/bin/bash", "-c", strings.Join(step3, " "))
	err = cmd.Run()
	if err != nil {
		logrus.Errorf(err.Error())
		return
	}
	return
}
