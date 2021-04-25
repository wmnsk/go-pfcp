// Copyright 2019-2021 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewQueryPacketRateStatusWithinPFCPSessionModificationRequest creates a new QueryPacketRateStatusWithinPFCPSessionModificationRequest IE.
func NewQueryPacketRateStatusWithinPFCPSessionModificationRequest(ies ...*IE) *IE {
	return newGroupedIE(QueryPacketRateStatusWithinPFCPSessionModificationRequest, 0, ies...)
}

// QueryPacketRateStatusWithinPFCPSessionModificationRequest returns the IEs above QueryPacketRateStatusWithinPFCPSessionModificationRequest if the type of IE matches.
func (i *IE) QueryPacketRateStatusWithinPFCPSessionModificationRequest() ([]*IE, error) {
	if i.Type != QueryPacketRateStatusWithinPFCPSessionModificationRequest {
		return nil, &InvalidTypeError{Type: i.Type}
	}

	return ParseMultiIEs(i.Payload)
}
