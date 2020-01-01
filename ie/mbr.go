// Copyright 2019-2020 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import (
	"encoding/binary"
	"io"
)

// NewMBR creates a new MBR IE.
func NewMBR(ul, dl uint32) *IE {
	i := New(MBR, make([]byte, 8))

	binary.BigEndian.PutUint32(i.Payload[0:4], ul)
	binary.BigEndian.PutUint32(i.Payload[4:8], dl)

	return i
}

// MBR returns MBR in []byte if the type of IE matches.
func (i *IE) MBR() ([]byte, error) {
	if i.Type != MBR {
		return nil, &InvalidTypeError{Type: i.Type}
	}

	return i.Payload, nil
}

// MBRUL returns MBRUL in uint32 if the type of IE matches.
func (i *IE) MBRUL() (uint32, error) {
	if i.Type != MBR {
		return 0, &InvalidTypeError{Type: i.Type}
	}

	if len(i.Payload) < 4 {
		return 0, io.ErrUnexpectedEOF
	}
	return binary.BigEndian.Uint32(i.Payload[0:4]), nil
}

// MBRDL returns MBRDL in uint32 if the type of IE matches.
func (i *IE) MBRDL() (uint32, error) {
	if i.Type != MBR {
		return 0, &InvalidTypeError{Type: i.Type}
	}

	if len(i.Payload) < 8 {
		return 0, io.ErrUnexpectedEOF
	}
	return binary.BigEndian.Uint32(i.Payload[4:8]), nil
}
