// Copyright 2019-2024 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import (
	"time"
)

// NewStartTime creates a new StartTime IE.
func NewStartTime(ts time.Time) *IE {
	u64sec := uint64(ts.Sub(time.Date(1900, time.January, 1, 0, 0, 0, 0, time.UTC))) / 1000000000
	return newUint32ValIE(StartTime, uint32(u64sec))
}

// StartTime returns StartTime in time.Time if the type of IE matches.
func (i *IE) StartTime() (time.Time, error) {
	switch i.Type {
	case StartTime:
		return i.valueAs3GPPTimestamp()
	case UsageReportWithinSessionModificationResponse,
		UsageReportWithinSessionDeletionResponse,
		UsageReportWithinSessionReportRequest:
		ies, err := i.UsageReport()
		if err != nil {
			return time.Time{}, err
		}
		for _, x := range ies {
			if x.Type == StartTime {
				return x.StartTime()
			}
		}
		return time.Time{}, ErrIENotFound
	case GTPUPathQoSReport:
		ies, err := i.GTPUPathQoSReport()
		if err != nil {
			return time.Time{}, err
		}
		for _, x := range ies {
			if x.Type == StartTime {
				return x.StartTime()
			}
		}
		return time.Time{}, ErrIENotFound
	case QoSMonitoringReport:
		ies, err := i.QoSMonitoringReport()
		if err != nil {
			return time.Time{}, err
		}
		for _, x := range ies {
			if x.Type == StartTime {
				return x.StartTime()
			}
		}
		return time.Time{}, ErrIENotFound
	default:
		return time.Time{}, &InvalidTypeError{Type: i.Type}
	}
}
