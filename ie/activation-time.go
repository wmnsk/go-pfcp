// Copyright 2019-2024 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import (
	"time"
)

// NewActivationTime creates a new ActivationTime IE.
func NewActivationTime(ts time.Time) *IE {
	u64sec := uint64(ts.Sub(time.Date(1900, time.January, 1, 0, 0, 0, 0, time.UTC))) / 1000000000
	return newUint32ValIE(ActivationTime, uint32(u64sec))
}

// ActivationTime returns ActivationTime in time.Time if the type of IE matches.
func (i *IE) ActivationTime() (time.Time, error) {
	switch i.Type {
	case ActivationTime:
		return i.valueAs3GPPTimestamp()
	case CreatePDR:
		ies, err := i.CreatePDR()
		if err != nil {
			return time.Time{}, err
		}
		for _, x := range ies {
			if x.Type == ActivationTime {
				return x.ActivationTime()
			}
		}
		return time.Time{}, ErrIENotFound
	case UpdatePDR:
		ies, err := i.UpdatePDR()
		if err != nil {
			return time.Time{}, err
		}
		for _, x := range ies {
			if x.Type == ActivationTime {
				return x.ActivationTime()
			}
		}
		return time.Time{}, ErrIENotFound
	default:
		return time.Time{}, &InvalidTypeError{Type: i.Type}
	}
}
