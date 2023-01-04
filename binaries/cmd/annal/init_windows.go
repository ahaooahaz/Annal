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

	// rootCmd.AddCommand(serve.Cmd) // TODO:
}

func initEnv() (err error) {
	config.ANNALROOT = os.Getenv("ANNAL_ROOT")
	return
}
