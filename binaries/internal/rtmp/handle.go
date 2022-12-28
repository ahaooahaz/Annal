package rtmp

import (
	"fmt"
	"time"

	"github.com/nareix/joy4/av/avutil"
	"github.com/nareix/joy4/av/pktque"
	"github.com/nareix/joy4/av/pubsub"
	"github.com/nareix/joy4/format/rtmp"
	"github.com/sirupsen/logrus"
)

func (s *serveRTMP) handlePlay(conn *rtmp.Conn) {
	s.rwmu.RLock()
	ch := s.channels[conn.URL.Path]
	s.rwmu.RUnlock()

	if ch != nil {
		cursor := ch.que.Latest()
		query := conn.URL.Query()

		if q := query.Get("delaygop"); q != "" {
			n := 0
			fmt.Sscanf(q, "%d", &n)
			cursor = ch.que.DelayedGopCount(n)
		} else if q := query.Get("delaytime"); q != "" {
			dur, _ := time.ParseDuration(q)
			cursor = ch.que.DelayedTime(dur)
		}

		filters := pktque.Filters{}

		if q := query.Get("waitkey"); q != "" {
			filters = append(filters, &pktque.WaitKeyFrame{})
		}

		filters = append(filters, &pktque.FixTime{StartFromZero: true, MakeIncrement: true})

		if q := query.Get("framedrop"); q != "" {
			n := 0
			fmt.Sscanf(q, "%d", &n)
			filters = append(filters, &FrameDropper{Interval: n})
		}

		if q := query.Get("delayskip"); q != "" {
			dur, _ := time.ParseDuration(q)
			skipper := &FrameDropper{DelaySkip: dur}
			if q := query.Get("skipinterval"); q != "" {
				n := 0
				fmt.Sscanf(q, "%d", &n)
				skipper.SkipInterval = n
			}
			filters = append(filters, skipper)
		}

		demuxer := &pktque.FilterDemuxer{
			Filter:  filters,
			Demuxer: cursor,
		}

		avutil.CopyFile(conn, demuxer)
	}
}

func (s *serveRTMP) handlePublish(conn *rtmp.Conn) {
	s.rwmu.Lock()
	ch := s.channels[conn.URL.Path]
	logrus.Infof("rtmp publish path: %v", conn.URL.Path)
	if ch == nil {
		ch = &Channel{}
		ch.que = pubsub.NewQueue()
		query := conn.URL.Query()
		if q := query.Get("cachegop"); q != "" {
			var n int
			fmt.Sscanf(q, "%d", &n)
			ch.que.SetMaxGopCount(n)
		}
		s.channels[conn.URL.Path] = ch
	} else {
		ch = nil
	}
	s.rwmu.Unlock()
	if ch == nil {
		return
	}

	avutil.CopyFile(ch.que, conn)

	s.rwmu.Lock()
	delete(s.channels, conn.URL.Path)
	s.rwmu.Unlock()
	ch.que.Close()
}
