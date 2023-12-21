package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/ahaooahaz/Annal/binaries/cmd/annal/child/version"
	"github.com/ahaooahaz/Annal/binaries/cmd/annal/child/webrtc"

	"github.com/ahaooahaz/Annal/binaries/config"

	"github.com/ahaooahaz/encapsutils"
	"github.com/sirupsen/logrus"
)

func init() {
	err := initEnv()
	if err != nil {
		panic(err.Error())
	}

	rootCmd.AddCommand(version.Cmd)
	rootCmd.AddCommand(webrtc.Cmd)
}

func initEnv() (err error) {
	config.ANNALROOT = os.Getenv("ANNAL_ROOT")
	now := time.Now()
	config.LOGFILE = fmt.Sprintf("%s/log/%s", config.ANNALROOT, fmt.Sprintf("%04d-%02d-%02d.log", now.Year(), now.Month(), now.Day()))
	config.ICONPATH = config.ANNALROOT + "/icons/icon.svg"
	config.NOTIFYSENDSH = config.ANNALROOT + "/scripts/notify-send.sh"
	config.ATJOBS = config.ANNALROOT + "/.at.jobs"

	err = encapsutils.CreateDir(filepath.Dir(config.LOGFILE), os.ModePerm)
	if err != nil {
		return
	}

	var f *os.File
	f, err = os.OpenFile(config.LOGFILE, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return
	}

	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetReportCaller(true)
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	logrus.SetOutput(f)

	logrus.Info("init done")
	return
}
