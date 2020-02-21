// Copyright 2019-2020 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import (
	"io"
)

// NewTrafficEndpointID creates a new TrafficEndpointID IE.
func NewTrafficEndpointID(id uint8) *IE {
	return newUint8ValIE(TrafficEndpointID, id)
}

// TrafficEndpointID returns TrafficEndpointID in uint8 if the type of IE matches.
func (i *IE) TrafficEndpointID() (uint8, error) {
	if i.Type != TrafficEndpointID {
		return 0, &InvalidTypeError{Type: i.Type}
	}
	if len(i.Payload) < 1 {
		return 0, io.ErrUnexpectedEOF
	}

	return i.Payload[0], nil
}
