// Copyright 2019-2020 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewEthernetPDUSessionInformation creates a new EthernetPDUSessionInformation IE.
func NewEthernetPDUSessionInformation(info uint8) *IE {
	return newUint8ValIE(EthernetPDUSessionInformation, info)
}

// EthernetPDUSessionInformation returns EthernetPDUSessionInformation in []byte if the type of IE matches.
func (i *IE) EthernetPDUSessionInformation() ([]byte, error) {
	if i.Type != EthernetPDUSessionInformation {
		return nil, &InvalidTypeError{Type: i.Type}
	}

	return i.Payload, nil
}

// HasETHI reports whether an IE has ETHI bit.
func (i *IE) HasETHI() bool {
	if i.Type != EthernetPDUSessionInformation {
		return false
	}
	if len(i.Payload) < 1 {
		return false
	}

	return has1stBit(i.Payload[0])
}
