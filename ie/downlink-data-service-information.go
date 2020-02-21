// Copyright 2019-2020 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import (
	"io"
)

// NewDownlinkDataServiceInformation creates a new DownlinkDataServiceInformation IE.
func NewDownlinkDataServiceInformation(hasPPI, hasQFI bool, ppi, qfi uint8) *IE {
	payload := make([]byte, 1)
	if hasPPI {
		payload[0] |= 0x01
		payload = append(payload, ppi)
	}
	if hasQFI {
		payload[0] |= 0x02
		payload = append(payload, qfi)
	}

	return New(DownlinkDataServiceInformation, payload)
}

// DownlinkDataServiceInformation returns DownlinkDataServiceInformation in []byte if the type of IE matches.
func (i *IE) DownlinkDataServiceInformation() ([]byte, error) {
	if i.Type != DownlinkDataServiceInformation {
		return nil, &InvalidTypeError{Type: i.Type}
	}

	return i.Payload, nil
}

// HasPPI reports whether downlink data service notification has PPI bit.
func (i *IE) HasPPI() bool {
	if i.Type != DownlinkDataServiceInformation {
		return false
	}
	if len(i.Payload) < 1 {
		return false
	}

	return has1stBit(i.Payload[0])
}

// HasQFI reports whether downlink data service notification has QFI bit.
func (i *IE) HasQFI() bool {
	if i.Type != DownlinkDataServiceInformation {
		return false
	}
	if len(i.Payload) < 1 {
		return false
	}

	return has2ndBit(i.Payload[0])
}

// PPI returns PPI in uint8 if the type of IE matches.
func (i *IE) PPI() (uint8, error) {
	if i.Type != DownlinkDataServiceInformation {
		return 0, &InvalidTypeError{Type: i.Type}
	}
	if len(i.Payload) < 2 {
		return 0, io.ErrUnexpectedEOF
	}

	return i.Payload[1], nil
}

// QFI returns QFI in uint8 if the type of IE matches.
func (i *IE) QFI() (uint8, error) {
	switch i.Type {
	case DownlinkDataServiceInformation:
		if len(i.Payload) < 2 {
			return 0, io.ErrUnexpectedEOF
		}

		return i.Payload[2], nil
	case QFI:
		if len(i.Payload) < 1 {
			return 0, io.ErrUnexpectedEOF
		}

		return i.Payload[0], nil
	default:
		return 0, &InvalidTypeError{Type: i.Type}
	}
}
