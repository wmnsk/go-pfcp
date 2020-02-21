// Copyright 2019-2020 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewRQI creates a new RQI IE.
func NewRQI(rqi uint8) *IE {
	return newUint8ValIE(RQI, rqi)
}

// RQI returns RQI in []byte if the type of IE matches.
func (i *IE) RQI() ([]byte, error) {
	if i.Type != RQI {
		return nil, &InvalidTypeError{Type: i.Type}
	}

	return i.Payload, nil
}

// HasRQI reports whether an IE has RQI bit.
func (i *IE) HasRQI() bool {
	if i.Type != RQI {
		return false
	}
	if len(i.Payload) < 1 {
		return false
	}

	return has1stBit(i.Payload[0])
}
