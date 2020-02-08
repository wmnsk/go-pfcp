// Copyright 2019-2020 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import (
	"encoding/binary"
	"io"
	"time"
)

// NewDurationMeasurement creates a new DurationMeasurement IE.
//
// The period should be within the range of uint32, otherwise it overflows.
func NewDurationMeasurement(duration time.Duration) *IE {
	return newUint32ValIE(DurationMeasurement, uint32(duration.Seconds()))
}

// DurationMeasurement returns DurationMeasurement in time.Duration if the type of IE matches.
func (i *IE) DurationMeasurement() (time.Duration, error) {
	if i.Type != DurationMeasurement {
		return 0, &InvalidTypeError{Type: i.Type}
	}
	if len(i.Payload) < 4 {
		return 0, io.ErrUnexpectedEOF
	}

	return time.Duration(binary.BigEndian.Uint32(i.Payload[0:4])) * time.Second, nil
}
