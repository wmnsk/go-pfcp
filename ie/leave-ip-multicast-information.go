// Copyright 2019-2020 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewLeaveIPMulticastInformationIEWithinUsageReport creates a new LeaveIPMulticastInformationIEWithinUsageReport IE.
func NewLeaveIPMulticastInformationIEWithinUsageReport(multi, source *IE) *IE {
	return newGroupedIE(LeaveIPMulticastInformationIEWithinUsageReport, 0, multi, source)
}

// LeaveIPMulticastInformationIEWithinUsageReport returns the IEs above LeaveIPMulticastInformationIEWithinUsageReport if the type of IE matches.
func (i *IE) LeaveIPMulticastInformationIEWithinUsageReport() ([]*IE, error) {
	switch i.Type {
	case LeaveIPMulticastInformationIEWithinUsageReport:
		return ParseMultiIEs(i.Payload)
	case UsageReportIEWithinPFCPSessionReportRequest:
		ies, err := i.UsageReport()
		if err != nil {
			return nil, err
		}
		for _, x := range ies {
			if x.Type == LeaveIPMulticastInformationIEWithinUsageReport {
				return x.LeaveIPMulticastInformationIEWithinUsageReport()
			}
		}
		return nil, ErrIENotFound
	default:
		return nil, &InvalidTypeError{Type: i.Type}
	}
}
