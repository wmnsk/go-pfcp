// Copyright 2019-2020 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import (
	"io"
)

// NewWeight creates a new Weight IE.
func NewWeight(weight uint8) *IE {
	return newUint8ValIE(Weight, weight)
}

// Weight returns Weight in uint8 if the type of IE matches.
func (i *IE) Weight() (uint8, error) {
	if i.Type != Weight {
		return 0, &InvalidTypeError{Type: i.Type}
	}
	if len(i.Payload) < 1 {
		return 0, io.ErrUnexpectedEOF
	}

	return i.Payload[0], nil
}
