package main

import (
	"os"

	"github.com/ahaooahaz/Annal/binaries/config"
)

func init() {
	err := initEnv()
	if err != nil {
		panic(err.Error())
	}

}

func initEnv() (err error) {
	config.ANNALROOT = os.Getenv("ANNAL_ROOT")
	return
}
