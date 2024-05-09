// Copyright 2019-2024 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewEthertype creates a new Ethertype IE.
func NewEthertype(typ uint16) *IE {
	return newUint16ValIE(Ethertype, typ)
}

// Ethertype returns Ethertype in uint16 if the type of IE matches.
func (i *IE) Ethertype() (uint16, error) {
	switch i.Type {
	case Ethertype:
		return i.ValueAsUint16()
	case PDI:
		ies, err := i.PDI()
		if err != nil {
			return 0, err
		}
		for _, x := range ies {
			if x.Type == Ethertype {
				return x.Ethertype()
			}
		}
		return 0, ErrIENotFound
	case EthernetPacketFilter:
		ies, err := i.EthernetPacketFilter()
		if err != nil {
			return 0, err
		}
		for _, x := range ies {
			if x.Type == Ethertype {
				return x.Ethertype()
			}
		}
		return 0, ErrIENotFound
	default:
		return 0, &InvalidTypeError{Type: i.Type}
	}

}
