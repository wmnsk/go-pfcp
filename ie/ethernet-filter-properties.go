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
	switch i.Type {
	case EthernetFilterProperties:
		return i.Payload, nil
	case EthernetPacketFilter:
		ies, err := i.EthernetPacketFilter()
		if err != nil {
			return nil, err
		}
		for _, x := range ies {
			if x.Type == EthernetFilterProperties {
				return x.EthernetFilterProperties()
			}
		}
		return nil, ErrIENotFound
	default:
		return nil, &InvalidTypeError{Type: i.Type}
	}
}

// HasBIDE reports whether an IE has BIDE bit.
func (i *IE) HasBIDE() bool {
	if len(i.Payload) < 1 {
		return false
	}

	switch i.Type {
	case EthernetFilterProperties:
		return has1stBit(i.Payload[0])
	case EthernetPacketFilter:
		ies, err := i.EthernetPacketFilter()
		if err != nil {
			return false
		}
		for _, x := range ies {
			if x.Type == EthernetFilterProperties {
				return x.HasBIDE()
			}
		}
		return false
	default:
		return false
	}
}
