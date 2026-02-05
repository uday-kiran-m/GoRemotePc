package main

import (
	"fmt"

	"github.com/pion/webrtc/v4"
)

func InitWebRTC() {
	config := webrtc.Configuration{
		ICEServers: []webrtc.ICEServer{
			{
				URLs: []string{"stun:stun.l.google.com:19302"},
			},
		},
	}

	// Create PeerConnection
	peerConnection, err := webrtc.NewPeerConnection(config)
	if err != nil {
		panic(err)
	}

	// Handle ICE candidates
	peerConnection.OnICECandidate(func(c *webrtc.ICECandidate) {
		if c != nil {
			fmt.Println("New ICE candidate:", c.ToJSON())
		}
	})

	// Data channel
	dataChannel, err := peerConnection.CreateDataChannel("chat", nil)
	if err != nil {
		panic(err)
	}

	dataChannel.OnOpen(func() {
		fmt.Println("Data channel open")
		dataChannel.SendText("hello from Go")
	})

	dataChannel.OnMessage(func(msg webrtc.DataChannelMessage) {
		fmt.Printf("Message: %s\n", string(msg.Data))
	})

	// Create offer
	offer, err := peerConnection.CreateOffer(nil)
	if err != nil {
		panic(err)
	}

	// Set local description
	err = peerConnection.SetLocalDescription(offer)
	if err != nil {
		panic(err)
	}

	// Output SDP (send this via signaling)
	fmt.Println("SDP OFFER:")
	fmt.Println(peerConnection.LocalDescription().SDP)
}
