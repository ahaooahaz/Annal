package rtmp

import (
	"context"
	"sync"
	"time"

	"github.com/nareix/joy4/av"
	"github.com/nareix/joy4/av/pubsub"
	"github.com/nareix/joy4/format/rtmp"
	"github.com/spf13/cobra"
)

var ServeRTMPCmd = &cobra.Command{
	Use:   "servertmp",
	Short: "simple rtmp server.",
	Long:  `simple rtmp server.`,
	Run:   doServeRTMP,
}

type Channel struct {
	que *pubsub.Queue
}

type serveRTMP struct {
	rwmu     sync.RWMutex
	channels map[string]*Channel
	server   *rtmp.Server
}

func doServeRTMP(cmd *cobra.Command, args []string) {
	addr := ""
	if len(args) != 0 {
		addr = args[0]
	}

	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()
	ServeRTMP(ctx, addr)
}

func ServeRTMP(ctx context.Context, addr string) {
	s := &serveRTMP{
		channels: make(map[string]*Channel),
	}

	server := &rtmp.Server{
		Addr: addr,
	}

	server.HandlePlay = func(conn *rtmp.Conn) {
		s.handlePlay(conn)
	}

	server.HandlePublish = func(conn *rtmp.Conn) {
		s.handlePublish(conn)
	}
	s.server = server

	s.server.ListenAndServe()
}

type FrameDropper struct {
	Interval     int
	n            int
	skipping     bool
	DelaySkip    time.Duration
	lasttime     time.Time
	lastpkttime  time.Duration
	delay        time.Duration
	SkipInterval int
}

func (f *FrameDropper) ModifyPacket(pkt *av.Packet, streams []av.CodecData, videoidx int, audioidx int) (drop bool, err error) {
	if f.DelaySkip != 0 && pkt.Idx == int8(videoidx) {
		now := time.Now()
		if !f.lasttime.IsZero() {
			realdiff := now.Sub(f.lasttime)
			pktdiff := pkt.Time - f.lastpkttime
			f.delay += realdiff - pktdiff
		}
		f.lasttime = time.Now()
		f.lastpkttime = pkt.Time

		if !f.skipping {
			if f.delay > f.DelaySkip {
				f.skipping = true
				f.delay = 0
			}
		} else {
			if pkt.IsKeyFrame {
				f.skipping = false
			}
		}
		if f.skipping {
			drop = true
		}

		if f.SkipInterval != 0 && pkt.IsKeyFrame {
			if f.n == f.SkipInterval {
				f.n = 0
				f.skipping = true
			}
			f.n++
		}
	}

	if f.Interval != 0 {
		if f.n >= f.Interval && pkt.Idx == int8(videoidx) && !pkt.IsKeyFrame {
			drop = true
			f.n = 0
		}
		f.n++
	}

	return
}
