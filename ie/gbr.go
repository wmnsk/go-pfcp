// Copyright 2019-2020 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import (
	"encoding/binary"
	"io"
)

// NewGBR creates a new GBR IE.
func NewGBR(ul, dl uint32) *IE {
	i := New(GBR, make([]byte, 8))

	binary.BigEndian.PutUint32(i.Payload[0:4], ul)
	binary.BigEndian.PutUint32(i.Payload[4:8], dl)

	return i
}

// GBR returns GBR in []byte if the type of IE matches.
func (i *IE) GBR() ([]byte, error) {
	if i.Type != GBR {
		return nil, &InvalidTypeError{Type: i.Type}
	}

	return i.Payload, nil
}

// GBRUL returns GBRUL in uint32 if the type of IE matches.
func (i *IE) GBRUL() (uint32, error) {
	if i.Type != GBR {
		return 0, &InvalidTypeError{Type: i.Type}
	}

	if len(i.Payload) < 4 {
		return 0, io.ErrUnexpectedEOF
	}
	return binary.BigEndian.Uint32(i.Payload[0:4]), nil
}

// GBRDL returns GBRDL in uint32 if the type of IE matches.
func (i *IE) GBRDL() (uint32, error) {
	if i.Type != GBR {
		return 0, &InvalidTypeError{Type: i.Type}
	}

	if len(i.Payload) < 8 {
		return 0, io.ErrUnexpectedEOF
	}
	return binary.BigEndian.Uint32(i.Payload[4:8]), nil
}
