// Copyright 2019-2024 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewCreateBridgeInfoForTSC creates a new CreateBridgeInfoForTSC IE.
func NewCreateBridgeInfoForTSC(bii uint8) *IE {
	return newUint8ValIE(CreateBridgeInfoForTSC, bii&0x01)
}

// CreateBridgeInfoForTSC returns CreateBridgeInfoForTSC in uint8 if the type of IE matches.
func (i *IE) CreateBridgeInfoForTSC() (uint8, error) {
	switch i.Type {
	case CreateBridgeInfoForTSC:
		return i.ValueAsUint8()
	default:
		return 0, &InvalidTypeError{Type: i.Type}
	}
}

// HasBII reports whether an IE has BII bit.
func (i *IE) HasBII() bool {
	v, err := i.CreateBridgeInfoForTSC()
	if err != nil {
		return false
	}

	return has1stBit(v)
}
