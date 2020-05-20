// Copyright 2019-2020 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewIPMuliticastAddressingInfoWithinSessionEstablishmentRequest creates a new IPMuliticastAddressingInfoWithinSessionEstablishmentRequest IE.
func NewIPMuliticastAddressingInfoWithinSessionEstablishmentRequest(multi, source *IE) *IE {
	return newGroupedIE(IPMuliticastAddressingInfoWithinSessionEstablishmentRequest, 0, multi, source)
}

// IPMuliticastAddressingInfoWithinSessionEstablishmentRequest returns the IEs above IPMuliticastAddressingInfoWithinSessionEstablishmentRequest if the type of IE matches.
func (i *IE) IPMuliticastAddressingInfoWithinSessionEstablishmentRequest() ([]*IE, error) {
	if i.Type != IPMuliticastAddressingInfoWithinSessionEstablishmentRequest {
		return nil, &InvalidTypeError{Type: i.Type}
	}

	return ParseMultiIEs(i.Payload)
}
