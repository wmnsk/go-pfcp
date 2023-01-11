// Copyright 2019-2022 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import "io"

// NewPMFControlInformation creates a new PMFControlInformation IE.
func NewPMFControlInformation(pmfi uint8) *IE {
	return newUint8ValIE(PMFControlInformation, pmfi&0x01)
}

// PMFControlInformation returns PMFControlInformation in uint8 if the type of IE matches.
func (i *IE) PMFControlInformation() (uint8, error) {
	if len(i.Payload) < 1 {
		return 0, io.ErrUnexpectedEOF
	}

	switch i.Type {
	case PMFControlInformation:
		return i.Payload[0], nil
	case ProvideATSSSControlInformation:
		ies, err := i.ProvideATSSSControlInformation()
		if err != nil {
			return 0, err
		}
		for _, x := range ies {
			if x.Type == PMFControlInformation {
				return x.PMFControlInformation()
			}
		}
		return 0, ErrIENotFound
	default:
		return 0, &InvalidTypeError{Type: i.Type}
	}
}

// HasPMFI reports whether an IE has PMFI bit.
func (i *IE) HasPMFI() bool {
	v, err := i.PMFControlInformation()
	if err != nil {
		return false
	}

	return has1stBit(v)
}

// HasDRTTI reports whether an IE has DRTTI bit.
func (i *IE) HasDRTTI() bool {
	v, err := i.PMFControlInformation()
	if err != nil {
		return false
	}

	return has2ndBit(v)
}

// HasPQPM reports whether an IE has PQPM bit.
func (i *IE) HasPQPM() bool {
	v, err := i.PMFControlInformation()
	if err != nil {
		return false
	}

	return has3rdBit(v)
}
