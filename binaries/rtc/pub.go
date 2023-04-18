package rtc

import (
	"github.com/pion/webrtc/v3"
	"github.com/sirupsen/logrus"
)

func OfferPublishToAnswer(api *webrtc.API, videoFile string, videoFps int, generateAnswerFunc func(webrtc.SessionDescription) (webrtc.SessionDescription, error)) (err error) {
	videoTrack, err := NewH264LocalStaticSampleVideoTrack(videoFile, videoFps)
	if err != nil {
		return
	}

	peerConnection, err := NewVideoTrackPeerConnection(api, videoTrack)
	if err != nil {
		return
	}

	var offer1, offer2, answer webrtc.SessionDescription
	offer1, err = peerConnection.CreateOffer(nil)
	if err != nil {
		return
	}

	err = peerConnection.SetLocalDescription(offer1)
	if err != nil {
		return
	}
	<-webrtc.GatheringCompletePromise(peerConnection)

	offer2 = *peerConnection.LocalDescription()
	logrus.Debug("offer2: ", offer2.SDP)

	answer, err = generateAnswerFunc(offer2)
	if err != nil {
		return
	}

	err = peerConnection.SetRemoteDescription(answer)
	if err != nil {
		return
	}

	return
}
