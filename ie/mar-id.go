// Copyright 2019-2020 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import (
	"encoding/binary"
)

// NewMARID creates a new MARID IE.
func NewMARID(id uint16) *IE {
	return newUint16ValIE(MARID, id)
}

// MARID returns MARID in uint16 if the type of IE matches.
func (i *IE) MARID() (uint16, error) {
	if i.Type != MARID {
		return 0, &InvalidTypeError{Type: i.Type}
	}

	if len(i.Payload) < 2 {
		return 0, &InvalidTypeError{Type: i.Type}
	}
	return binary.BigEndian.Uint16(i.Payload[0:2]), nil
}
