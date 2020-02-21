// Copyright 2019-2020 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewEthernetFilterProperties creates a new EthernetFilterProperties IE.
func NewEthernetFilterProperties(props uint8) *IE {
	return newUint8ValIE(EthernetFilterProperties, props)
}

// EthernetFilterProperties returns EthernetFilterProperties in []byte if the type of IE matches.
func (i *IE) EthernetFilterProperties() ([]byte, error) {
	if i.Type != EthernetFilterProperties {
		return nil, &InvalidTypeError{Type: i.Type}
	}

	return i.Payload, nil
}

// HasBIDE reports whether an IE has BIDE bit.
func (i *IE) HasBIDE() bool {
	if i.Type != EthernetFilterProperties {
		return false
	}
	if len(i.Payload) < 1 {
		return false
	}

	return has1stBit(i.Payload[0])
}
