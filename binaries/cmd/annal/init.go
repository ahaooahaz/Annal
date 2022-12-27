package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/AHAOAHA/Annal/binaries/internal/config"
	"github.com/AHAOAHA/Annal/binaries/internal/gui"
	"github.com/AHAOAHA/Annal/binaries/internal/jt"
	"github.com/AHAOAHA/Annal/binaries/internal/rtmp"
	"github.com/AHAOAHA/Annal/binaries/internal/storage"
	"github.com/AHAOAHA/Annal/binaries/internal/todo"
	"github.com/AHAOAHA/Annal/binaries/internal/version"
	"github.com/AHAOAHA/encapsutils"
	"github.com/sirupsen/logrus"
)

func init() {
	err := initEnv()
	if err != nil {
		panic(err.Error())
	}

	rootCmd.AddCommand(gui.Cmd)
	rootCmd.AddCommand(jt.Cmd)
	rootCmd.AddCommand(rtmp.ServeRTMPCmd)
	rootCmd.AddCommand(todo.Cmd)
	rootCmd.AddCommand(version.Cmd)
}

func initEnv() (err error) {
	config.ANNALROOT = os.Getenv("ANNAL_ROOT")
	now := time.Now()
	config.LOGFILE = fmt.Sprintf("%s/log/%s", config.ANNALROOT, fmt.Sprintf("%04d-%02d-%02d.log", now.Year(), now.Month(), now.Day()))

	err = encapsutils.MustCreateFile(config.LOGFILE)
	if err != nil {
		return
	}

	var f *os.File
	f, err = os.OpenFile(config.LOGFILE, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return
	}

	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetReportCaller(true)
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	logrus.SetOutput(f)

	migrations := config.ANNALROOT + "/migrations/sqlite3"
	storageFile := config.ANNALROOT + "/storage/annal.db"
	config.DBPATH = storageFile
	err = encapsutils.MustCreateDir(filepath.Dir(storageFile), os.ModePerm)
	if err != nil {
		return
	}
	log := logrus.WithFields(logrus.Fields{
		"migrations":   migrations,
		"storage_file": storageFile,
	})
	err = storage.Sqlite3Migrate(migrations, storageFile)
	if err != nil {
		log.Errorf("migrate failed, err: %v", err.Error())
		return
	}
	log.Info("init done")
	return
}