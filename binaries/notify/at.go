package notify

import (
	"context"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

var (
	_AtTimeFormatString = "200601021504"
)

func CreateOnTimeJob(ctx context.Context, scriptFile string, expTime time.Time) (ID uint64, err error) {
	var scriptAbs string
	scriptAbs, err = filepath.Abs(scriptFile)
	if err != nil {
		return
	}

	attime := expTime.Format(_AtTimeFormatString)
	step3 := []string{"at", "-t", attime, "-f", scriptAbs}
	command3 := strings.Join(step3, " ")
	cmd := exec.Command("/bin/bash", "-c", command3)

	var out []byte
	out, err = cmd.CombinedOutput()
	if err != nil {
		logrus.Errorf("command3: %s, error: %s", command3, err.Error())
		return
	}

	compileRegex := regexp.MustCompile("job ([0-9]+) ")
	matchArr := compileRegex.FindStringSubmatch(string(out))
	ID, err = strconv.ParseUint(matchArr[len(matchArr)-1], 10, 64)
	if err != nil {
		return
	}
	return
}
