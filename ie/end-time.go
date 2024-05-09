// Copyright 2019-2024 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import (
	"time"
)

// NewEndTime creates a new EndTime IE.
func NewEndTime(ts time.Time) *IE {
	u64sec := uint64(ts.Sub(time.Date(1900, time.January, 1, 0, 0, 0, 0, time.UTC))) / 1000000000
	return newUint32ValIE(EndTime, uint32(u64sec))
}

// EndTime returns EndTime in time.Time if the type of IE matches.
func (i *IE) EndTime() (time.Time, error) {
	switch i.Type {
	case EndTime:
		return i.valueAs3GPPTimestamp()
	case UsageReportWithinSessionModificationResponse,
		UsageReportWithinSessionDeletionResponse,
		UsageReportWithinSessionReportRequest:
		ies, err := i.UsageReport()
		if err != nil {
			return time.Time{}, err
		}
		for _, x := range ies {
			if x.Type == EndTime {
				return x.EndTime()
			}
		}
		return time.Time{}, ErrIENotFound
	default:
		return time.Time{}, &InvalidTypeError{Type: i.Type}
	}

}
