package rtc

import (
	"net"

	"github.com/pion/interceptor"
	"github.com/pion/logging"
	"github.com/pion/webrtc/v3"
)

var logger = logging.NewDefaultLoggerFactory().NewLogger("zlmediakit")

type RTCEngine_NETWORKTYPE int

const (
	RTCEngineNETWORKTYPE_TCP RTCEngine_NETWORKTYPE = iota + 1
	RTCEngineNETWORKTYPE_UDP
	RTCEngineNETWORKTYPE_MIX
)

type RTCEngineConfiguration struct {
	NetworkType RTCEngine_NETWORKTYPE
	TCP         *net.TCPAddr
	UDP         *net.UDPAddr
}

func init() {

}

func NewRTCEngine(conf *RTCEngineConfiguration) (api *webrtc.API, err error) {
	var t, u bool
	switch conf.NetworkType {
	case RTCEngineNETWORKTYPE_TCP:
		t = true
	case RTCEngineNETWORKTYPE_UDP:
		u = true
	case RTCEngineNETWORKTYPE_MIX:
		t = true
		u = true
	}

	m := &webrtc.MediaEngine{}
	if err = m.RegisterDefaultCodecs(); err != nil {
		return
	}

	i := &interceptor.Registry{}
	if err = webrtc.RegisterDefaultInterceptors(m, i); err != nil {
		return
	}

	settingEngine := webrtc.SettingEngine{}

	if t {
		// Enable support only for TCP ICE candidates.
		settingEngine.SetNetworkTypes([]webrtc.NetworkType{
			webrtc.NetworkTypeTCP4,
			webrtc.NetworkTypeUDP4,
		})
		var tcpListener *net.TCPListener
		tcpListener, err = net.ListenTCP("tcp", conf.TCP)
		if err != nil {
			return
		}

		tcpMux := webrtc.NewICETCPMux(logger, tcpListener, 8)
		settingEngine.SetICETCPMux(tcpMux)
	}
	if u {
		var udpListener *net.UDPConn
		udpListener, err = net.ListenUDP("udp", conf.UDP)
		if err != nil {
			return
		}
		udpMux := webrtc.NewICEUDPMux(logger, udpListener)
		settingEngine.SetICEUDPMux(udpMux)
	}

	api = webrtc.NewAPI(webrtc.WithMediaEngine(m), webrtc.WithInterceptorRegistry(i), webrtc.WithSettingEngine(settingEngine))
	return
}
