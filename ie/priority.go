// Copyright 2019-2020 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import "io"

// Priority definitions.
const (
	PriorityActive    uint8 = 0
	PriorityStandby   uint8 = 1
	PriorityNoStandby uint8 = 2
	PriorityHigh      uint8 = 3
	PriorityLow       uint8 = 4
)

// NewPriority creates a new Priority IE.
func NewPriority(priority uint8) *IE {
	return newUint8ValIE(Priority, priority&0x0f)
}

// Priority returns Priority in uint8 if the type of IE matches.
func (i *IE) Priority() (uint8, error) {
	if i.Type != Priority {
		return 0, &InvalidTypeError{Type: i.Type}
	}
	if len(i.Payload) == 0 {
		return 0, io.ErrUnexpectedEOF
	}

	return i.Payload[0], nil
}
