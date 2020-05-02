// Copyright 2019-2020 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewPagingPolicyIndicator creates a new PagingPolicyIndicator IE.
func NewPagingPolicyIndicator(indicator uint8) *IE {
	return newUint8ValIE(PagingPolicyIndicator, indicator)
}

// PagingPolicyIndicator returns PagingPolicyIndicator in uint8 if the type of IE matches.
func (i *IE) PagingPolicyIndicator() (uint8, error) {
	if i.Type != PagingPolicyIndicator {
		return 0, &InvalidTypeError{Type: i.Type}
	}

	return i.Payload[0] & 0x07, nil
}
