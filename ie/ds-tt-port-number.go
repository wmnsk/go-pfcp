// Copyright 2019-2024 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewDSTTPortNumber creates a new DSTTPortNumber IE.
func NewDSTTPortNumber(port uint32) *IE {
	return newUint32ValIE(DSTTPortNumber, port)
}

// DSTTPortNumber returns DSTTPortNumber in uint32 if the type of IE matches.
func (i *IE) DSTTPortNumber() (uint32, error) {
	switch i.Type {
	case DSTTPortNumber:
		return i.ValueAsUint32()
	case CreatedBridgeInfoForTSC:
		ies, err := i.CreatedBridgeInfoForTSC()
		if err != nil {
			return 0, err
		}
		for _, x := range ies {
			if x.Type == DSTTPortNumber {
				return x.DSTTPortNumber()
			}
		}
		return 0, ErrIENotFound
	default:
		return 0, &InvalidTypeError{Type: i.Type}
	}

}
