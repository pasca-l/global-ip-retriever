package network

import (
	"errors"
	"time"
)

type Ips struct {
	LocalIp  string
	PublicIp string
}

func RetrieveIps() (Ips, error) {
	ips := Ips{}
	err := ips.retrieveIps()
	if err != nil {
		return Ips{}, err
	}

	return ips, nil
}

func (ips *Ips) retrieveIps() error {
	conn, err := newWebRtcConnection()
	if err != nil {
		return err
	}
	defer conn.Close()

	// using channel as signaling
	retriever := make(chan struct{})
	conn.OnICECandidate(handleOnICECandidate(ips, retriever))

	if _, err = conn.CreateDataChannel("", nil); err != nil {
		return err
	}

	offer, err := conn.CreateOffer(nil)
	if err != nil {
		return err
	}
	if err = conn.SetLocalDescription(offer); err != nil {
		return err
	}

	select {
	case <-retriever:
	case <-time.After(10 * time.Second):
		return errors.New("time out on retrieving data")
	}

	return nil
}
