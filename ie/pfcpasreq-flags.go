// Copyright 2019-2024 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewPFCPASReqFlags creates a new PFCPASReqFlags IE.
func NewPFCPASReqFlags(flag uint8) *IE {
	return newUint8ValIE(PFCPASReqFlags, flag)
}

// PFCPASReqFlags returns PFCPASReqFlags in uint8 if the type of IE matches.
func (i *IE) PFCPASReqFlags() (uint8, error) {
	if i.Type != PFCPASReqFlags {
		return 0, &InvalidTypeError{Type: i.Type}
	}

	return i.ValueAsUint8()
}

// HasUUPSI reports whether an IE has UUPSI bit.
func (i *IE) HasUUPSI() bool {
	v, err := i.PFCPASReqFlags()
	if err != nil {
		return false
	}

	return has1stBit(v)
}
