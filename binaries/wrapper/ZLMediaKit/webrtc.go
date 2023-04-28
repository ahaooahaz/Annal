package zlmediakit

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/ahaooahaz/Annal/binaries/rtc"
	"github.com/pion/webrtc/v3"
	"github.com/sirupsen/logrus"
)

type WebRTCResponse struct {
	ID   string `json:"id"`
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	SDP  string `json:"sdp"`
}

func WebRTC(IP string, HTTPPort uint16, app, stream string, rtcType rtc.RTCTYPE, offer webrtc.SessionDescription) (answer webrtc.SessionDescription, err error) {
	tp := ""
	switch rtcType {
	case rtc.RTCTYPE_PUBLISH:
		tp = "push"
	case rtc.RTCTYPE_PLAY:
		tp = "play"
	default:
		err = fmt.Errorf("invalid rtc type: %v", rtcType)
		return
	}

	url := fmt.Sprintf("http://%s:%d/index/api/webrtc?app=%v&stream=%v&type=%v", IP, HTTPPort, app, stream, tp)
	method := "POST"

	payload := strings.NewReader(string(offer.SDP))
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		logrus.Error(err)
		return
	}

	res, err := client.Do(req)
	if err != nil {
		logrus.Error(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logrus.Error(err)
		return
	}

	resp := make(map[string]interface{})
	err = json.Unmarshal(body, &resp)
	if err != nil {
		logrus.Error(err)
		return
	}

	if resp["code"].(float64) != 0 {
		err = fmt.Errorf("code: %v, msg: %v", resp["code"], resp["msg"])
		return
	}

	sdp := resp["sdp"].(string)

	answer = webrtc.SessionDescription{
		Type: webrtc.SDPTypeAnswer,
		SDP:  sdp,
	}
	return
}
