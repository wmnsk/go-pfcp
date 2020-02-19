// Copyright 2019-2020 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import (
	"io"
)

// NewBARID creates a new BARID IE.
func NewBARID(id uint8) *IE {
	return newUint8ValIE(BARID, id)
}

// BARID returns BARID in uint8 if the type of IE matches.
func (i *IE) BARID() (uint8, error) {
	if i.Type != BARID {
		return 0, &InvalidTypeError{Type: i.Type}
	}
	if len(i.Payload) < 1 {
		return 0, io.ErrUnexpectedEOF
	}

	return i.Payload[0], nil
}
