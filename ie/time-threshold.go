// Copyright 2019-2024 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import (
	"encoding/binary"
	"io"
	"time"
)

// NewTimeThreshold creates a new TimeThreshold IE.
func NewTimeThreshold(t time.Duration) *IE {
	return newUint32ValIE(TimeThreshold, uint32(t.Seconds()))
}

// TimeThreshold returns TimeThreshold in time.Duration if the type of IE matches.
func (i *IE) TimeThreshold() (time.Duration, error) {
	if len(i.Payload) < 4 {
		return 0, io.ErrUnexpectedEOF
	}

	switch i.Type {
	case TimeThreshold:
		t := binary.BigEndian.Uint32(i.Payload[0:4])
		return time.Duration(t) * time.Second, nil
	case CreateURR:
		ies, err := i.CreateURR()
		if err != nil {
			return 0, err
		}
		for _, x := range ies {
			if x.Type == TimeThreshold {
				return x.TimeThreshold()
			}
		}
		return 0, ErrIENotFound
	case UpdateURR:
		ies, err := i.UpdateURR()
		if err != nil {
			return 0, err
		}
		for _, x := range ies {
			if x.Type == TimeThreshold {
				return x.TimeThreshold()
			}
		}
		return 0, ErrIENotFound
	default:
		return 0, &InvalidTypeError{Type: i.Type}
	}
}
