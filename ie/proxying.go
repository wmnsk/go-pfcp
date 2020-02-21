// Copyright 2019-2020 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewProxying creates a new Proxying IE.
func NewProxying(ins, arp uint8) *IE {
	return newUint8ValIE(Proxying, (ins<<1)|arp)
}

// Proxying returns Proxying in uint8 if the type of IE matches.
func (i *IE) Proxying() (uint8, error) {
	if i.Type != Proxying {
		return 0, &InvalidTypeError{Type: i.Type}
	}

	return i.Payload[0], nil
}

// HasINS reports whether an IE has INS bit.
func (i *IE) HasINS() bool {
	if i.Type != Proxying {
		return false
	}
	if len(i.Payload) < 1 {
		return false
	}

	return has2ndBit(i.Payload[0])
}

// HasARP reports whether an IE has ARP bit.
func (i *IE) HasARP() bool {
	if i.Type != Proxying {
		return false
	}
	if len(i.Payload) < 1 {
		return false
	}

	return has1stBit(i.Payload[0])
}
