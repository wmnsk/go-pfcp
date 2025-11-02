// Copyright 2019-2024 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewOffendingIE creates a new OffendingIE IE.
func NewOffendingIE(itype IEType) *IE {
	return newUint16ValIE(OffendingIE, uint16(itype))
}

// OffendingIE returns OffendingIE in IEType if the type of IE matches.
func (i *IE) OffendingIE() (IEType, error) {
	if i.Type != OffendingIE {
		return 0, &InvalidTypeError{Type: i.Type}
	}

	n, err := i.ValueAsUint16()
	if err != nil {
		return 0, err
	}
	return IEType(n), nil
}
