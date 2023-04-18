package rtc

import (
	"fmt"

	"github.com/pion/webrtc/v3"
	"github.com/sirupsen/logrus"
)

func NewVideoTrackPeerConnection(api *webrtc.API, videoTrack *webrtc.TrackLocalStaticSample) (peerConnection *webrtc.PeerConnection, err error) {
	peerConnection, err = api.NewPeerConnection(webrtc.Configuration{
		ICEServers: []webrtc.ICEServer{
			{
				URLs: []string{"stun:stun.l.google.com:19302"},
			},
		},
		ICETransportPolicy: webrtc.ICETransportPolicyAll,
	})
	if err != nil {
		return
	}

	peerConnection.OnDataChannel(func(d *webrtc.DataChannel) {
		d.OnOpen(func() {
			logrus.Infof("data channel open")
		})
		d.OnMessage(func(msg webrtc.DataChannelMessage) {
			logrus.Infof("data channel message: %v", string(msg.Data))
		})
		d.OnClose(func() {
			logrus.Infof("data channel close")
		})
	})

	peerConnection.OnTrack(func(remoteTrack *webrtc.TrackRemote, receiver *webrtc.RTPReceiver) {
		logrus.Infof("peerConnection OnTrack")
	})

	peerConnection.OnICECandidate(func(candidate *webrtc.ICECandidate) {
	})

	peerConnection.OnICEConnectionStateChange(func(connectionState webrtc.ICEConnectionState) {
		if connectionState == webrtc.ICEConnectionStateDisconnected {
			if err := peerConnection.Close(); err != nil {
				logrus.Errorf("peerConnection.Close() error: %v", err.Error())
			}
		} else if connectionState == webrtc.ICEConnectionStateConnected {
			logrus.Infof("peerConnection connected")
		}
		logrus.Infof("peerConnection state: %v", connectionState.String())
	})

	var rtpSender *webrtc.RTPSender
	if rtpSender, err = peerConnection.AddTrack(videoTrack); err != nil {
		logrus.Errorf("peerConnection.AddTrack() error: %v", err.Error())
		return
	} else {
		go func() {
			for {
				buf := make([]byte, 1024*1024)
				var rtcpErr error
				if _, _, rtcpErr = rtpSender.Read(buf); rtcpErr != nil {
					fmt.Printf("read error: %v", rtcpErr.Error())
					return
				}
			}
		}()
	}

	return
}
