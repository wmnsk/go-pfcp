// Copyright 2019-2020 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import (
	"encoding/binary"
	"io"
)

// NewPacketDetectionRuleID creates a new PacketDetectionRuleID IE.
func NewPacketDetectionRuleID(id uint16) *IE {
	return newUint16ValIE(PacketDetectionRuleID, id)
}

// PacketDetectionRuleID returns PacketDetectionRuleID in uint16 if the type of IE matches.
func (i *IE) PacketDetectionRuleID() (uint16, error) {
	switch i.Type {
	case PacketDetectionRuleID:
		if len(i.Payload) < 2 {
			return 0, io.ErrUnexpectedEOF
		}
		return binary.BigEndian.Uint16(i.Payload[0:2]), nil
	case ApplicationDetectionInformation:
		ies, err := i.ApplicationDetectionInformation()
		if err != nil {
			return 0, err
		}
		for _, x := range ies {
			if x.Type == PacketDetectionRuleID {
				return x.PacketDetectionRuleID()
			}
		}
		return 0, ErrIENotFound
	case UsageReportIEWithinPFCPSessionReportRequest:
		ies, err := i.UsageReport()
		if err != nil {
			return 0, err
		}
		for _, x := range ies {
			if x.Type == ApplicationDetectionInformation {
				return x.PacketDetectionRuleID()
			}
		}
		return 0, ErrIENotFound
	case DownlinkDataReport:
		ies, err := i.DownlinkDataReport()
		if err != nil {
			return 0, err
		}
		for _, x := range ies {
			if x.Type == PacketDetectionRuleID {
				return x.PacketDetectionRuleID()
			}
		}
		return 0, ErrIENotFound
	case UpdatedPDR:
		ies, err := i.UpdatedPDR()
		if err != nil {
			return 0, err
		}
		for _, x := range ies {
			if x.Type == PacketDetectionRuleID {
				return x.PacketDetectionRuleID()
			}
		}
		return 0, ErrIENotFound
	default:
		return 0, &InvalidTypeError{Type: i.Type}
	}
}
