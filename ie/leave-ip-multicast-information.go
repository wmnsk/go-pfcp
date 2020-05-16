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
	if i.Type != LeaveIPMulticastInformationIEWithinUsageReport {
		return nil, &InvalidTypeError{Type: i.Type}
	}

	return ParseMultiIEs(i.Payload)
}
