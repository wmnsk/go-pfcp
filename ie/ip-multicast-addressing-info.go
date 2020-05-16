// Copyright 2019-2020 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewIPMuliticastAddressingInfoWithinPFCPSessionEstablishmentRequest creates a new IPMuliticastAddressingInfoWithinPFCPSessionEstablishmentRequest IE.
func NewIPMuliticastAddressingInfoWithinPFCPSessionEstablishmentRequest(multi, source *IE) *IE {
	return newGroupedIE(IPMuliticastAddressingInfoWithinPFCPSessionEstablishmentRequest, 0, multi, source)
}

// IPMuliticastAddressingInfoWithinPFCPSessionEstablishmentRequest returns the IEs above IPMuliticastAddressingInfoWithinPFCPSessionEstablishmentRequest if the type of IE matches.
func (i *IE) IPMuliticastAddressingInfoWithinPFCPSessionEstablishmentRequest() ([]*IE, error) {
	if i.Type != IPMuliticastAddressingInfoWithinPFCPSessionEstablishmentRequest {
		return nil, &InvalidTypeError{Type: i.Type}
	}

	return ParseMultiIEs(i.Payload)
}
