// Copyright 2019-2020 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewJoinIPMulticastInformationIEWithinUsageReport creates a new JoinIPMulticastInformationIEWithinUsageReport IE.
func NewJoinIPMulticastInformationIEWithinUsageReport(multi, source *IE) *IE {
	return newGroupedIE(JoinIPMulticastInformationIEWithinUsageReport, 0, multi, source)
}

// JoinIPMulticastInformationIEWithinUsageReport returns the IEs above JoinIPMulticastInformationIEWithinUsageReport if the type of IE matches.
func (i *IE) JoinIPMulticastInformationIEWithinUsageReport() ([]*IE, error) {
	if i.Type != JoinIPMulticastInformationIEWithinUsageReport {
		return nil, &InvalidTypeError{Type: i.Type}
	}

	return ParseMultiIEs(i.Payload)
}
