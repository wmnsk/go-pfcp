// Copyright 2019-2022 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewOffendingIEInformation creates a new OffendingIEInformation IE.
func NewOffendingIEInformation(offendingIE *IE) *IE {
	return newGroupedIE(OffendingIEInformation, 0, offendingIE)
}

// OffendingIEInformation returns the IE above OffendingIEInformation if the type of IE matches.
func (i *IE) OffendingIEInformation() (*IE, error) {
	if i.Type != OffendingIEInformation {
		return nil, &InvalidTypeError{Type: i.Type}
	}
	return Parse(i.Payload)
}
