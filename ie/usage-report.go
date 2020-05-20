// Copyright 2019-2020 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewUsageReport creates a new UsageReport IE.
func NewUsageReport(typ uint16, ies ...*IE) *IE {
	return newGroupedIE(typ, 0, ies...)
}

// NewUsageReportIEWithinPFCPSessionModificationResponse creates a new UsageReportIEWithinPFCPSessionModificationResponse IE.
func NewUsageReportIEWithinPFCPSessionModificationResponse(urr, seq, trigger, start, end, vol, dur, firstPkt, lastPkt, usage, query, eth *IE) *IE {
	return NewUsageReport(UsageReportIEWithinPFCPSessionModificationResponse, urr, seq, trigger, start, end, vol, dur, firstPkt, lastPkt, usage, query, eth)
}

// NewUsageReportIEWithinPFCPSessionDeletionResponse creates a new UsageReportIEWithinPFCPSessionDeletionResponse IE.
func NewUsageReportIEWithinPFCPSessionDeletionResponse(urr, seq, trigger, start, end, vol, dur, firstPkt, lastPkt, usage, eth *IE) *IE {
	return NewUsageReport(UsageReportIEWithinPFCPSessionDeletionResponse, urr, seq, trigger, start, end, vol, dur, firstPkt, lastPkt, usage, eth)
}

// NewUsageReportIEWithinPFCPSessionReportRequest creates a new UsageReportIEWithinPFCPSessionReportRequest IE.
func NewUsageReportIEWithinPFCPSessionReportRequest(urr, seq, trigger, start, end, vol, dur, app, ip, firstPkt, lastPkt, usage, query, ts, eth, join, leave *IE) *IE {
	return NewUsageReport(UsageReportIEWithinPFCPSessionReportRequest, urr, seq, trigger, start, end, vol, dur, app, ip, firstPkt, lastPkt, usage, query, ts, eth, join, leave)
}

// UsageReport returns the IEs above UsageReport if the type of IE matches.
func (i *IE) UsageReport() ([]*IE, error) {
	switch i.Type {
	case UsageReportIEWithinPFCPSessionModificationResponse,
		UsageReportIEWithinPFCPSessionDeletionResponse,
		UsageReportIEWithinPFCPSessionReportRequest:

		return ParseMultiIEs(i.Payload)
	default:
		return nil, &InvalidTypeError{Type: i.Type}
	}
}
