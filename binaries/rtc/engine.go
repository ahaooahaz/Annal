package rtc

import (
	"net"

	"github.com/pion/interceptor"
	"github.com/pion/webrtc/v3"
)

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

func NewRTCEngine(conf *RTCEngineConfiguration) (api *webrtc.API, err error) {
	if conf == nil {
		conf = &RTCEngineConfiguration{
			NetworkType: RTCEngineNETWORKTYPE_UDP,
			UDP:         &net.UDPAddr{IP: net.IP{0, 0, 0, 0}, Port: 31938},
		}
	}
	var t, u bool
	switch conf.NetworkType {
	case RTCEngineNETWORKTYPE_TCP:
		t = true
	case RTCEngineNETWORKTYPE_UDP:
		u = true
	case RTCEngineNETWORKTYPE_MIX:
		t = true
		u = true
	default:
		panic("invalid network type")
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
		tcpListener, err = net.ListenTCP("tcp", &net.TCPAddr{
			IP:   net.IP{0, 0, 0, 0},
			Port: 31937,
		})
		if err != nil {
			return
		}

		tcpMux := webrtc.NewICETCPMux(nil, tcpListener, 8)
		settingEngine.SetICETCPMux(tcpMux)
	}
	if u {
		var udpListener *net.UDPConn
		udpListener, err = net.ListenUDP("udp", &net.UDPAddr{
			IP:   net.IP{0, 0, 0, 0},
			Port: 31938,
		})
		if err != nil {
			return
		}
		udpMux := webrtc.NewICEUDPMux(nil, udpListener)
		settingEngine.SetICEUDPMux(udpMux)
	}

	api = webrtc.NewAPI(webrtc.WithMediaEngine(m), webrtc.WithInterceptorRegistry(i), webrtc.WithSettingEngine(settingEngine))
	return
}
