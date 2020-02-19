// Copyright 2019-2020 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import (
	"encoding/binary"
	"io"
)

// NewURSEQN creates a new URSEQN IE.
func NewURSEQN(seq uint32) *IE {
	return newUint32ValIE(URSEQN, seq)
}

// URSEQN returns URSEQN in uint32 if the type of IE matches.
func (i *IE) URSEQN() (uint32, error) {
	if i.Type != URSEQN {
		return 0, &InvalidTypeError{Type: i.Type}
	}
	if len(i.Payload) < 4 {
		return 0, io.ErrUnexpectedEOF
	}

	return binary.BigEndian.Uint32(i.Payload[0:4]), nil
}
