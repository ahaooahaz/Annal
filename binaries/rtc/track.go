package rtc

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/google/uuid"
	webrtc "github.com/pion/webrtc/v3"
	"github.com/pion/webrtc/v3/pkg/media"
	"github.com/pion/webrtc/v3/pkg/media/h264reader"
	"github.com/sirupsen/logrus"
)

// loop
func NewH264LocalStaticSampleVideoTrack(videoFile string, fps int) (videoTrack *webrtc.TrackLocalStaticSample, err error) {
	if videoFile == "" {
		err = fmt.Errorf("videoFile is empty")
		return
	}

	var videoFileAbs string
	videoFileAbs, err = filepath.Abs(videoFile)
	if err != nil {
		return
	}

	_, err = os.Stat(videoFileAbs)
	if err != nil {
		return
	}

	videoTrack, err = webrtc.NewTrackLocalStaticSample(webrtc.RTPCodecCapability{
		MimeType: webrtc.MimeTypeH264,
	}, "video", fmt.Sprintf("rtckit/%s", uuid.New().String()))
	if err != nil {
		return
	}

	go h264VideoFileConsumer(videoTrack, videoFileAbs, time.Duration(1000/fps)*time.Millisecond)

	return
}

func h264VideoFileConsumer(videoTrack *webrtc.TrackLocalStaticSample, videoFile string, videoFps time.Duration) {
	for {
		file, h264Err := os.Open(videoFile)
		if h264Err != nil {
			logrus.Errorf("open file failed, err: %v", h264Err.Error())
			continue
		}

		h264, h264Err := h264reader.NewReader(file)
		if h264Err != nil {
			logrus.Errorf("new reader failed, err: %v", h264Err.Error())
			continue
		}

		buf := make(chan []byte, 1024*1024) // 1MB?

		var wg sync.WaitGroup

		wg.Add(2)
		go func() {
			defer wg.Done()
			for data := range buf {
				sample := media.Sample{Data: data, Duration: videoFps, Timestamp: time.Now(), PacketTimestamp: uint32(time.Now().Unix())}

				if h264Err1 := videoTrack.WriteSample(sample); h264Err1 != nil {
					logrus.Errorf("write sample failed, err: %v", h264Err1.Error())
					continue
				}
			}
		}()

		go func() {
			defer wg.Done()
			ticker := time.NewTicker(videoFps)
			for ; true; <-ticker.C {
				nal, h264Err := h264.NextNAL()
				if h264Err == io.EOF {
					logrus.Warnf("all video frames parsed and sent, loop playback")
					break
				}
				if h264Err != nil {
					logrus.Errorf("next nal failed, err: %v", h264Err.Error())
					break
				}

				buf <- nal.Data
			}

			close(buf)
		}()

		wg.Wait()
		_ = file.Close()
	}
}
