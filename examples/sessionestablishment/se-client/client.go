// Copyright 2019-2020 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

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
		server = flag.String("-s", "127.0.0.2:8805", "server's addr/port")
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
	nid := ie.NewNodeID("172.55.55.101", "", "")
	fseid := ie.NewFSEID(0xbebbebebebbebebf, net.IPv4(172, 55, 55, 101), nil, nil)

	pdrid := ie.NewPacketDetectionRuleID(1)
	precedence := ie.NewPrecedence(0)

	sourceInterface := ie.NewSourceInterface(ie.SrcInterfaceAccess)
	fteid := ie.NewFTEID(0, nil, nil, nil)
	ueIPAddress := ie.NewUEIPAddress(0x02, "10.0.0.3", "", 0)
	pdi := ie.NewPDI(0, sourceInterface, fteid, ueIPAddress)
	outerHeaderRemoval := ie.NewOuterHeaderRemoval(0, 0)
	farid := ie.NewFARID(1)
	createPDR := ie.NewCreatePDR(0, pdrid, precedence, pdi, outerHeaderRemoval, farid)

	applyAction := ie.NewApplyAction(0x02)
	destinationInterface := ie.NewDestinationInterface(ie.DstInterfaceCore)
	outerHeaderCreation := ie.NewOuterHeaderCreation(0x0100, 0x00000001, "192.168.20.136", "", 0, 0, 0)
	forwardingParameters := ie.NewForwardingParameters(0, destinationInterface, outerHeaderCreation)
	createFAR := ie.NewCreateFAR(0, farid, applyAction, forwardingParameters)
	hbreq, err := message.NewSessionEstablishmentRequset(nid, fseid, createPDR, createFAR).Marshal()
	if err != nil {
		log.Fatal(err)
	}

	if _, err := conn.Write(hbreq); err != nil {
		log.Fatal(err)
	}
	log.Printf("sent Heartbeat Request to: %s", raddr)

	if err := conn.SetReadDeadline(time.Now().Add(3 * time.Second)); err != nil {
		log.Fatal(err)
	}
}
