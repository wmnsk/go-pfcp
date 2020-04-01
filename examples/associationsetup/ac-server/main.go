package main

import (
	"flag"
	"github.com/wmnsk/go-pfcp/ie"
	"github.com/wmnsk/go-pfcp/message"
	"log"
	"net"
	"time"
)

func main() {
	var (
		listen = flag.String("-s", "127.0.0.1:8805", "addr/port to listen on")
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

		acreq, ok := msg.(*message.AssociationSetupRequest)
		if !ok {
			log.Printf("got unexpected message: %s, from: %s", msg.MessageTypeName(), addr)
			continue
		}

		ts, err := acreq.RecoveryTimeStamp.RecoveryTimeStamp()
		if err != nil {
			log.Printf("got Association Setup Request with invalid TS: %s, from: %s", err, addr)
			continue
		} else {
			log.Printf("got Association Setup Request with TS: %s, from: %s", ts, addr)
		}

		//IE required for encapsulation
		nodeID := ie.NewNodeID("172.55.55.101", "", "")
		rets := ie.NewRecoveryTimeStamp(time.Now())
		cause := ie.NewCause(ie.CauseRequestAccepted)
		cp := ie.NewCPFunctionFeatures(0x00)
		// Timestamp shouldn't be the time message is sent in the real deployment but anyway :D
		acres, err := message.NewAssociationSetupResponse(nodeID, cause, rets, cp).Marshal()
		if err != nil {
			log.Fatal(err)
		}

		if _, err := conn.WriteTo(acres, addr); err != nil {
			log.Fatal(err)
		}
		log.Printf("sent Association Setup  Response to: %s", addr)
	}
}
