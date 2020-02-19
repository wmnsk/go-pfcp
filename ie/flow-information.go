// Copyright 2019-2020 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import (
	"encoding/binary"
	"io"
)

const (
	FlowDirectionUnspecified   uint8 = 0
	FlowDirectionDownlink      uint8 = 1
	FlowDirectionUplink        uint8 = 2
	FlowDirectionBidirectional uint8 = 3
)

// NewFlowInformation creates a new FlowInformation IE.
func NewFlowInformation(dir uint8, desc string) *IE {
	d := []byte(desc)
	l := len(d)

	i := New(FlowInformation, make([]byte, 3+l))
	i.Payload[0] = dir
	binary.BigEndian.PutUint16(i.Payload[1:3], uint16(l))
	copy(i.Payload[3:], d)

	return i
}

// FlowInformation returns FlowInformation in []byte if the type of IE matches.
func (i *IE) FlowInformation() ([]byte, error) {
	if i.Type != FlowInformation {
		return nil, &InvalidTypeError{Type: i.Type}
	}

	return i.Payload, nil
}

// FlowDirection returns FlowDirection in uint8 if the type of IE matches.
func (i *IE) FlowDirection() (uint8, error) {
	if i.Type != FlowInformation {
		return 0, &InvalidTypeError{Type: i.Type}
	}
	if len(i.Payload) < 1 {
		return 0, io.ErrUnexpectedEOF
	}

	return i.Payload[0] & 0x07, nil
}

// FlowDescription returns FlowDescription in string if the type of IE matches.
func (i *IE) FlowDescription() (string, error) {
	if i.Type != FlowInformation {
		return "", &InvalidTypeError{Type: i.Type}
	}
	if len(i.Payload) < 3 {
		return "", io.ErrUnexpectedEOF
	}
	l := binary.BigEndian.Uint16(i.Payload[1:3])

	if len(i.Payload) < int(l) {
		return "", io.ErrUnexpectedEOF
	}

	return string(i.Payload[4:l]), nil
}
