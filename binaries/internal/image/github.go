package image

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"time"

	goimgtype "github.com/shamsher31/goimgtype"
	"github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
)

const (
	_UPLOADTEMPLATE = "https://api.github.com/repos/%s/contents/%s"
	_PROXYTEMPLATE  = `https://cdn.jsdelivr.net/gh/%s@%s/%s`
)

type uploadMessage struct {
	Message string `json:"message"`
	Content string `json:"content"`
	SHA     string `json:"sha"`
}

func GithubPutImage(ctx context.Context, repo, branch, token, imagepath string) (url string, err error) {
	_, err = goimgtype.Get(imagepath)
	if err != nil {
		logrus.Errorf("get image type failed, err: %v", err.Error())
		return
	}

	p := fmt.Sprintf("images/%s/%s", time.Now().Format("20060102"), filepath.Base(imagepath))

	var request *http.Request
	request, err = newUploadRequest(ctx, repo, token, imagepath, p)
	if err != nil {
		logrus.Errorf("new request failed, err: %v", err.Error())
		return
	}

	tokenSource := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)

	if !debug {
		client := oauth2.NewClient(ctx, tokenSource)
		var response *http.Response
		response, err = client.Do(request)
		if err != nil {
			logrus.Errorf("do request failed, err: %v", err.Error())
			return
		}
		defer response.Body.Close()

		if response.StatusCode/100 != 2 {
			err = fmt.Errorf("response status code: %v", response.StatusCode)
			return
		}

		var b []byte
		b, err = ioutil.ReadAll(response.Body)
		if err != nil {
			return
		}
		logrus.Infof("response: %v", string(b))
	}

	url = fmt.Sprintf(_PROXYTEMPLATE, repo, branch, p)

	return
}

func newUploadRequest(ctx context.Context, repo, token, imagepath, hostimagepath string) (request *http.Request, err error) {
	var content []byte
	content, err = ioutil.ReadFile(imagepath)
	if err != nil {
		return
	}

	contentBase64 := make([]byte, base64.StdEncoding.EncodedLen(len(content)))
	base64.StdEncoding.Encode(contentBase64, content)

	msg := &uploadMessage{
		Message: fmt.Sprintf("upload %s", filepath.Base(imagepath)),
		Content: string(contentBase64),
	}

	var msgRaw []byte
	msgRaw, err = json.Marshal(msg)
	if err != nil {
		return
	}

	request, err = http.NewRequest(http.MethodPut, fmt.Sprintf(_UPLOADTEMPLATE, repo, hostimagepath), bytes.NewReader(msgRaw))
	if err != nil {
		return
	}
	return
}
