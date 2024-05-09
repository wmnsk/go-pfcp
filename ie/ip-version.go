// Copyright 2019-2024 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewIPVersion creates a new IPVersion IE.
func NewIPVersion(v4, v6 bool) *IE {
	i := New(IPVersion, make([]byte, 1))
	if v4 {
		i.Payload[0] |= 0x01
	}
	if v6 {
		i.Payload[0] |= 0x02
	}
	return i
}

// IPVersion returns IPVersion in uint8 if the type of IE matches.
func (i *IE) IPVersion() (uint8, error) {
	switch i.Type {
	case IPVersion:
		return i.ValueAsUint8()
	case UEIPAddressPoolInformation:
		ies, err := i.UEIPAddressPoolInformation()
		if err != nil {
			return 0, err
		}
		for _, x := range ies {
			if x.Type == IPVersion {
				return x.IPVersion()
			}
		}
		return 0, ErrIENotFound
	default:
		return 0, &InvalidTypeError{Type: i.Type}
	}
}
