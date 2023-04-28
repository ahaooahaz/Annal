package zlmediakit

import (
	"fmt"
	"net"
	"os"

	"github.com/ahaooahaz/Annal/binaries/rtc"
	zlmkit "github.com/ahaooahaz/Annal/binaries/wrapper/ZLMediaKit"
	"github.com/pion/webrtc/v3"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "zlmediakit",
	Short: "zlmediakit WebRTC client",
	Long:  `zlmediakit WebRTC client`,
	Run: func(cmd *cobra.Command, args []string) {
		var api *webrtc.API
		conf := &rtc.RTCEngineConfiguration{}

		if *arg_TCP {
			conf.NetworkType = rtc.RTCEngineNETWORKTYPE_TCP
			conf.TCP = &net.TCPAddr{IP: net.IP{0, 0, 0, 0}, Port: 31922}
		}

		if *arg_UDP {
			conf.NetworkType = rtc.RTCEngineNETWORKTYPE_UDP
			conf.UDP = &net.UDPAddr{IP: net.IP{0, 0, 0, 0}, Port: 31921}
		}

		if *arg_UDP && *arg_TCP {
			conf.NetworkType = rtc.RTCEngineNETWORKTYPE_MIX
		}

		api, err := rtc.NewRTCEngine(conf)
		if err != nil {
			logrus.Error(err)
			fmt.Fprintf(os.Stderr, "failed to create rtc engine: %v", err)
			return
		}

		err = rtc.OfferPublishToAnswer(api, *arg_SOURCE, *arg_FPS, func(offer webrtc.SessionDescription) (answer webrtc.SessionDescription, e error) {
			res, e := zlmkit.WebRTC(*arg_IP, *arg_PORT, *arg_APP, *arg_STREAM, rtc.RTCTYPE_PUBLISH, offer)
			if e != nil {
				fmt.Print(e.Error())
				return
			}
			answer = res
			return
		})

		if err != nil {
			fmt.Print(err.Error())
			return
		}

		select {}
	},
}

var (
	arg_IP, arg_APP, arg_STREAM, arg_SOURCE *string
	arg_PORT                                *uint16
	arg_FPS                                 *int
	arg_TCP, arg_UDP                        *bool
)

func init() {
	arg_IP = Cmd.Flags().StringP("ip", "", "127.0.0.1", "zlmediakit server ip address")
	arg_PORT = Cmd.Flags().Uint16P("port", "p", 80, "zlmediakit http server port")
	arg_FPS = Cmd.Flags().IntP("fps", "", 25, "video source fps")
	arg_APP = Cmd.Flags().StringP("app", "a", "app", "zlmediakit app name")
	arg_STREAM = Cmd.Flags().StringP("stream", "s", "stream", "zlmediakit stream name")
	arg_SOURCE = Cmd.Flags().StringP("source", "", "test.h264", "zlmediakit source file (H264)")
	arg_TCP = Cmd.Flags().BoolP("tcp", "", false, " use tcp to send h264 data")
	arg_UDP = Cmd.Flags().BoolP("udp", "", false, " use udp to send h264 data")
}
