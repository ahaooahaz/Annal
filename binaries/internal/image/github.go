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

	"golang.org/x/oauth2"
)

const (
	_UPLOADTEMPLATE = "https://api.github.com/repos/AHAOAHA/ImageHosting/contents/images/%s"
)

type uploadMessage struct {
	Message string `json:"message"`
	Content string `json:"content"`
	SHA     string `json:"sha"`
}

func GithubPutImage(ctx context.Context, token, path string) (err error) {
	var request *http.Request
	request, err = newUploadRequest(ctx, token, path)
	if err != nil {
		return
	}

	tokenSource := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)

	client := oauth2.NewClient(ctx, tokenSource)
	var response *http.Response
	response, err = client.Do(request)
	if err != nil {
		return
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		err = fmt.Errorf("response status code: %v", response.StatusCode)
		return
	}

	// var b []byte
	// b, err = ioutil.ReadAll(response.Body)
	// if err != nil {
	// 	return
	// }

	return
}

func newUploadRequest(ctx context.Context, token, imagepath string) (request *http.Request, err error) {
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

	p := filepath.Base(imagepath)

	request, err = http.NewRequest(http.MethodPut, fmt.Sprintf(_UPLOADTEMPLATE, p), bytes.NewReader(msgRaw))
	if err != nil {
		return
	}
	return
}
