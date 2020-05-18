// Copyright 2019-2020 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import "io"

// NewATSSSLLControlInformation creates a new ATSSSLLControlInformation IE.
func NewATSSSLLControlInformation(lli uint8) *IE {
	return newUint8ValIE(ATSSSLLControlInformation, lli&0x01)
}

// ATSSSLLControlInformation returns ATSSSLLControlInformation in uint8 if the type of IE matches.
func (i *IE) ATSSSLLControlInformation() (uint8, error) {
	if len(i.Payload) < 1 {
		return 0, io.ErrUnexpectedEOF
	}

	switch i.Type {
	case ATSSSLLControlInformation:
		return i.Payload[0], nil
	default:
		return 0, &InvalidTypeError{Type: i.Type}
	}
}

// HasLLI reports whether an IE has LLI bit.
func (i *IE) HasLLI() bool {
	v, err := i.ATSSSLLControlInformation()
	if err != nil {
		return false
	}

	return has1stBit(v)
}
