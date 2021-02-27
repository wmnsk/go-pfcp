// Copyright 2019-2021 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

// Command hb-server sends a HeartbeatRequest and checks response.
//
// Heartbeat exchanging feature is planned be included in the go-pfcp package's
// built-in functions in the future.
package main

import (
	"flag"
	"log"
	"net"
	"time"

	"github.com/wmnsk/go-pfcp/ie"
	"github.com/wmnsk/go-pfcp/message"
)

func main() {
	var (
		listen = flag.String("-s", "127.0.0.2:8805", "addr/port to listen on")
	)
	flag.Parse()

	laddr, err := net.ResolveUDPAddr("udp", *listen)
	if err != nil {
		log.Fatal(err)
	}

	conn, err := net.ListenUDP("udp", laddr)
	if err != nil {
		log.Fatal(err)
	}

	buf := make([]byte, 1500)
	for {
		log.Printf("waiting for messages to come on: %s", laddr)
		n, addr, err := conn.ReadFrom(buf)
		if err != nil {
			log.Fatal(err)
		}

		msg, err := message.Parse(buf[:n])
		if err != nil {
			log.Printf("ignored undecodable message: %x, error: %s", buf[:n], err)
			continue
		}

		hbreq, ok := msg.(*message.HeartbeatRequest)
		if !ok {
			log.Printf("got unexpected message: %s, from: %s", msg.MessageTypeName(), addr)
			continue
		}

		ts, err := hbreq.RecoveryTimeStamp.RecoveryTimeStamp()
		if err != nil {
			log.Printf("got Heartbeat Request with invalid TS: %s, from: %s", err, addr)
			continue
		} else {
			log.Printf("got Heartbeat Request with TS: %s, from: %s", ts, addr)
		}

		// Timestamp shouldn't be the time message is sent in the real deployment but anyway :D
		var seq uint32 = 1
		hbres, err := message.NewHeartbeatResponse(seq, ie.NewRecoveryTimeStamp(time.Now())).Marshal()
		if err != nil {
			log.Fatal(err)
		}

		if _, err := conn.WriteTo(hbres, addr); err != nil {
			log.Fatal(err)
		}
		log.Printf("sent Heartbeat Response to: %s", addr)
	}
}
