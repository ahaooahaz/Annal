package global

import (
	"fmt"
	"os"
	"time"

	"github.com/AHAOAHA/encapsutils"
	"github.com/sirupsen/logrus"
)

var (
	ANNALROOT = ""
	LOGFILE   = ""
)

func init() {
	ANNALROOT = os.Getenv("ANNAL_ROOT")
	now := time.Now()
	LOGFILE = fmt.Sprintf("%s/log/%s", ANNALROOT, fmt.Sprintf("%04d-%02d-%02d.log", now.Year(), now.Month(), now.Day()))

	err := encapsutils.MustCreateFile(LOGFILE)
	if err != nil {
		panic(err.Error())
	}

	var f *os.File
	f, err = os.OpenFile(LOGFILE, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err.Error())
	}

	logrus.SetOutput(f)
}
