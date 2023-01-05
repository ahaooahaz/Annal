package main

import (
	"os"

	"github.com/AHAOAHA/Annal/binaries/internal/config"
)

func init() {
	err := initEnv()
	if err != nil {
		panic(err.Error())
	}

	rootCmd.AddCommand(image.Cmd)
}

func initEnv() (err error) {
	config.ANNALROOT = os.Getenv("ANNAL_ROOT")
	return
}
