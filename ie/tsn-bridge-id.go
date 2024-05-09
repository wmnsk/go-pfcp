// Copyright 2019-2024 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import (
	"io"
	"net"
)

// NewTSNBridgeID creates a new TSNBridgeID IE.
func NewTSNBridgeID(mac net.HardwareAddr) *IE {
	if mac == nil {
		return New(TSNBridgeID, []byte{0x00})
	}
	l := len(mac) + 1
	b := make([]byte, l)
	b[0] = 0x01
	copy(b[1:l], mac)
	return New(TSNBridgeID, b)
}

// HasMAC reports whether an IE has MAC bit.
func (i *IE) HasMAC() bool {
	switch i.Type {
	case TSNBridgeID:
		return has1stBit(i.Payload[0])
	default:
		return false
	}
}

// TSNBridgeID returns TSNBridgeID in net.HardwareAddr if the type of IE matches.
func (i *IE) TSNBridgeID() (net.HardwareAddr, error) {
	if len(i.Payload) < 1 {
		return nil, io.ErrUnexpectedEOF
	}

	switch i.Type {
	case TSNBridgeID:
		if !has1stBit(i.Payload[0]) {
			return nil, nil
		}
		if (len(i.Payload) != 7) && (len(i.Payload) < 9) {
			return nil, io.ErrUnexpectedEOF
		}
		if len(i.Payload) == 7 {
			return net.HardwareAddr(i.Payload[1:7]), nil
		}
		return net.HardwareAddr(i.Payload[1:9]), nil
	case CreatedBridgeInfoForTSC:
		ies, err := i.CreatedBridgeInfoForTSC()
		if err != nil {
			return nil, err
		}
		for _, x := range ies {
			if x.Type == TSNBridgeID {
				return x.TSNBridgeID()
			}
		}
		return nil, ErrIENotFound
	default:
		return nil, &InvalidTypeError{Type: i.Type}
	}
}
