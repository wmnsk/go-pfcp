// Copyright 2019-2020 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewPortManagementInformationForTSC creates a new PortManagementInformationForTSC IE.
func NewPortManagementInformationForTSC(typ uint16, info *IE) *IE {
	return newGroupedIE(typ, 0, info)
}

// NewPortManagementInformationForTSCIEWithinPFCPSessionModificationRequest creates a new PortManagementInformationForTSCIEWithinPFCPSessionModificationRequest IE.
func NewPortManagementInformationForTSCIEWithinPFCPSessionModificationRequest(info *IE) *IE {
	return newGroupedIE(PortManagementInformationForTSCIEWithinPFCPSessionModificationRequest, 0, info)
}

// NewPortManagementInformationForTSCIEWithinPFCPSessionModificationResponse creates a new PortManagementInformationForTSCIEWithinPFCPSessionModificationResponse IE.
func NewPortManagementInformationForTSCIEWithinPFCPSessionModificationResponse(info *IE) *IE {
	return newGroupedIE(PortManagementInformationForTSCIEWithinPFCPSessionModificationResponse, 0, info)
}

// NewPortManagementInformationForTSCIEWithinPFCPSessionReportRequest creates a new PortManagementInformationForTSCIEWithinPFCPSessionReportRequest IE.
func NewPortManagementInformationForTSCIEWithinPFCPSessionReportRequest(info *IE) *IE {
	return newGroupedIE(PortManagementInformationForTSCIEWithinPFCPSessionReportRequest, 0, info)
}

// PortManagementInformationForTSC returns the IEs above PortManagementInformationForTSC if the type of IE matches.
func (i *IE) PortManagementInformationForTSC() ([]*IE, error) {
	switch i.Type {
	case PortManagementInformationForTSCIEWithinPFCPSessionModificationRequest,
		PortManagementInformationForTSCIEWithinPFCPSessionModificationResponse,
		PortManagementInformationForTSCIEWithinPFCPSessionReportRequest:

		return ParseMultiIEs(i.Payload)
	default:
		return nil, &InvalidTypeError{Type: i.Type}
	}
}
