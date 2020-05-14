// Copyright 2019-2020 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import "io"

// NewMTEDTControlInformation creates a new MTEDTControlInformation IE.
func NewMTEDTControlInformation(rdsi uint8) *IE {
	return newUint8ValIE(MTEDTControlInformation, rdsi&0x01)
}

// MTEDTControlInformation returns MTEDTControlInformation in uint8 if the type of IE matches.
func (i *IE) MTEDTControlInformation() (uint8, error) {
	if len(i.Payload) < 1 {
		return 0, io.ErrUnexpectedEOF
	}

	switch i.Type {
	case MTEDTControlInformation:
		return i.Payload[0], nil
	default:
		return 0, &InvalidTypeError{Type: i.Type}
	}
}

// HasRDSI reports whether an IE has MTEDTControlInformation bit.
func (i *IE) HasRDSI() bool {
	v, err := i.MTEDTControlInformation()
	if err != nil {
		return false
	}

	return has1stBit(v)
}
