package rtc

type RTCTYPE int

const (
	RTCTYPE_PUBLISH RTCTYPE = iota + 1
	RTCTYPE_PLAY
)

type PROTOCOL int

const (
	TCP PROTOCOL = iota + 1
	UDP
)
