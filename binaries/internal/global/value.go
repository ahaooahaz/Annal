package global

import "os"

var (
	ANNALROOT = ""
)

func init() {
	ANNALROOT = os.Getenv("ANNAL_ROOT")
}
