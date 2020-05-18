// Copyright 2019-2020 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewAccessAvailabilityControlInformation creates a new AccessAvailabilityControlInformation IE.
func NewAccessAvailabilityControlInformation(info *IE) *IE {
	return newGroupedIE(AccessAvailabilityControlInformation, 0, info)
}

// AccessAvailabilityControlInformation returns the IEs above AccessAvailabilityControlInformation if the type of IE matches.
func (i *IE) AccessAvailabilityControlInformation() ([]*IE, error) {
	if i.Type != AccessAvailabilityControlInformation {
		return nil, &InvalidTypeError{Type: i.Type}
	}

	return ParseMultiIEs(i.Payload)
}
