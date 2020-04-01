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
		server = flag.String("-s", "127.0.0.1:8805", "server's addr/port")
	)
	flag.Parse()

	raddr, err := net.ResolveUDPAddr("udp", *server)
	if err != nil {
		log.Fatal(err)
	}

	conn, err := net.DialUDP("udp", nil, raddr)
	if err != nil {
		log.Fatal(err)
	}
	//IE required for encapsulation
	nodeID := ie.NewNodeID("172.55.55.102", "", "")
	ts := ie.NewRecoveryTimeStamp(time.Now())
	up := ie.NewUPFunctionFeatures(0x10, 0x00)

	acreq, err := message.NewAssociationSetupRequest(nodeID, ts, up).Marshal()
	if err != nil {
		log.Fatal(err)
	}

	if _, err := conn.Write(acreq); err != nil {
		log.Fatal(err)
	}
	log.Printf("sent AssociationSetup  Request to: %s", raddr)

	if err := conn.SetReadDeadline(time.Now().Add(3 * time.Second)); err != nil {
		log.Fatal(err)
	}

	buf := make([]byte, 1500)
	waiting := true
	for waiting {
		n, addr, err := conn.ReadFrom(buf)
		if err != nil {
			log.Fatal(err)
		}

		msg, err := message.Parse(buf[:n])
		if err != nil {
			log.Printf("ignored undecodable message: %x, error: %s", buf[:n], err)
			continue
		}

		acres, ok := msg.(*message.AssociationSetupResponse)
		if !ok {
			log.Printf("got unexpected message: %s, from: %s", msg.MessageTypeName(), addr)
			continue
		}

		waiting = false
		ts, err := acres.RecoveryTimeStamp.RecoveryTimeStamp()
		if err != nil {
			log.Printf("got Association Setup  Response with invalid TS: %s, from: %s", err, addr)
			break
		}
		log.Printf("got Association Setup  Response with TS: %s, from: %s", ts, addr)
	}
}
