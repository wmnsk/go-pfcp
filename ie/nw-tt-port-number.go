// Copyright 2019-2024 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewNWTTPortNumber creates a new NWTTPortNumber IE.
func NewNWTTPortNumber(port uint32) *IE {
	return newUint32ValIE(NWTTPortNumber, port)
}

// NWTTPortNumber returns NWTTPortNumber in uint32 if the type of IE matches.
func (i *IE) NWTTPortNumber() (uint32, error) {
	switch i.Type {
	case NWTTPortNumber:
		return i.ValueAsUint32()
	case CreatedBridgeInfoForTSC:
		ies, err := i.CreatedBridgeInfoForTSC()
		if err != nil {
			return 0, err
		}
		for _, x := range ies {
			if x.Type == NWTTPortNumber {
				return x.NWTTPortNumber()
			}
		}
		return 0, ErrIENotFound
	default:
		return 0, &InvalidTypeError{Type: i.Type}
	}

}
