// Copyright 2019-2020 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewUpdateBAR creates a new UpdateBAR IE.
func NewUpdateBAR(typ uint16, ies ...*IE) *IE {
	return newGroupedIE(typ, 0, ies...)
}

// NewUpdateBARIEWithinPFCPSessionModificationRequest creates a new UpdateBARIEWithinPFCPSessionModificationRequest IE.
func NewUpdateBARIEWithinPFCPSessionModificationRequest(bar, delay, bufCount, mtedt *IE) *IE {
	return NewUpdateBAR(UpdateBARIEWithinPFCPSessionModificationRequest, bar, delay, bufCount, mtedt)
}

// NewUpdateBARIEWithinPCFPSessionReportResponse creates a new UpdateBARIEWithinPCFPSessionReportResponse IE.
func NewUpdateBARIEWithinPCFPSessionReportResponse(bar, delay, duration, dlCount, bufCount *IE) *IE {
	return NewUpdateBAR(UpdateBARIEWithinPCFPSessionReportResponse, bar, delay, duration, dlCount, bufCount)
}

// UpdateBAR returns the IEs above UpdateBAR if the type of IE matches.
func (i *IE) UpdateBAR() ([]*IE, error) {
	switch i.Type {
	case UpdateBARIEWithinPFCPSessionModificationRequest,
		UpdateBARIEWithinPCFPSessionReportResponse:

		return ParseMultiIEs(i.Payload)
	default:
		return nil, &InvalidTypeError{Type: i.Type}
	}
}
