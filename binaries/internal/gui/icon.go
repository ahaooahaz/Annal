package gui

import (
	"io/ioutil"

	"fyne.io/fyne/v2"
	"github.com/AHAOAHA/Annal/binaries/internal/config"
)

var (
	icon fyne.Resource
)

func initEnv() (err error) {
	var iconRaw []byte
	iconRaw, err = ioutil.ReadFile(config.ANNALROOT + "/icons/icon.svg")
	if err != nil {
		return
	}

	icon = fyne.NewStaticResource("icon", iconRaw)
	return
}
