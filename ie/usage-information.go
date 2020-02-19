// Copyright 2019-2020 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewUsageInformation creates a new UsageInformation IE.
func NewUsageInformation(bef, aft, uae, ube int) *IE {
	return newUint8ValIE(UsageInformation, uint8((ube<<3)|(uae<<2)|(aft<<1)|(bef)))
}

// UsageInformation returns UsageInformation in uint8 if the type of IE matches.
func (i *IE) UsageInformation() (uint8, error) {
	if i.Type != UsageInformation {
		return 0, &InvalidTypeError{Type: i.Type}
	}

	return i.Payload[0], nil
}

// HasUBE reports whether UsageInformation IE has UBE bit.
func (i *IE) HasUBE() bool {
	v, err := i.UsageInformation()
	if err != nil {
		return false
	}

	return has4thBit(v)
}

// HasUAE reports whether UsageInformation IE has UAE bit.
func (i *IE) HasUAE() bool {
	v, err := i.UsageInformation()
	if err != nil {
		return false
	}

	return has3rdBit(v)
}

// HasAFT reports whether UsageInformation IE has AFT bit.
func (i *IE) HasAFT() bool {
	v, err := i.UsageInformation()
	if err != nil {
		return false
	}

	return has2ndBit(v)
}

// HasBEF reports whether UsageInformation IE has BEF bit.
func (i *IE) HasBEF() bool {
	v, err := i.UsageInformation()
	if err != nil {
		return false
	}

	return has1stBit(v)
}
