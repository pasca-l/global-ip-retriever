package network

import (
	"github.com/pion/webrtc/v3"
)

func configureWebRtc() webrtc.Configuration {
	config := webrtc.Configuration{
		ICEServers: []webrtc.ICEServer{
			{
				URLs: []string{"stun:stun.l.google.com:19302"},
			},
		},
	}
	return config
}

func newWebRtcConnection() (*webrtc.PeerConnection, error) {
	config := configureWebRtc()
	conn, err := webrtc.NewPeerConnection(config)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func handleOnICECandidate(
	ips *Ips,
	retriever chan struct{},
) func(*webrtc.ICECandidate) {
	return func(c *webrtc.ICECandidate) {
		// close channel, when there is no more callback
		if c == nil {
			close(retriever)
			return
		}
		switch c.Typ {
		case webrtc.ICECandidateTypeHost:
			ips.LocalIp = c.Address
		case webrtc.ICECandidateTypeSrflx:
			ips.PublicIp = c.Address
		}
	}
}
