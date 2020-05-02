// Copyright 2019-2020 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import (
	"encoding/binary"
	"io"
)

// Framed-Routing definitions.
//
// Ref: https://tools.ietf.org/html/rfc2865#section-5.10
const (
	FramedRoutingNone                    uint32 = 0
	FramedRoutingSendRoutingPackets      uint32 = 1
	FramedRoutingListenForRoutingPackets uint32 = 2
	FramedRoutingSendAndListen           uint32 = 3
)

// NewFramedRouting creates a new FramedRouting IE.
func NewFramedRouting(routing uint32) *IE {
	return newUint32ValIE(FramedRouting, routing)
}

// FramedRouting returns FramedRouting in uint32 if the type of IE matches.
func (i *IE) FramedRouting() (uint32, error) {
	if i.Type != FramedRouting {
		return 0, &InvalidTypeError{Type: i.Type}
	}
	if len(i.Payload) < 4 {
		return 0, io.ErrUnexpectedEOF
	}

	return binary.BigEndian.Uint32(i.Payload[0:4]), nil
}
