// Copyright 2019-2020 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import (
	"io"
)

// NewUsageReportTrigger creates a new UsageReportTrigger IE.
func NewUsageReportTrigger(triggerOctets ...uint8) *IE {
	return New(UsageReportTrigger, triggerOctets)
}

// UsageReportTrigger returns UsageReportTrigger in []byte if the type of IE matches.
func (i *IE) UsageReportTrigger() ([]byte, error) {
	if i.Type != UsageReportTrigger {
		return nil, &InvalidTypeError{Type: i.Type}
	}
	if len(i.Payload) < 3 {
		return nil, io.ErrUnexpectedEOF
	}

	return i.Payload, nil
}

// HasIMMER reports whether reporting trigger has IMMER bit.
func (i *IE) HasIMMER() bool {
	if len(i.Payload) < 1 {
		return false
	}

	switch i.Type {
	case UsageReportTrigger:
		u8 := uint8(i.Payload[0])
		return has8thBit(u8)
	default:
		return false
	}
}

// HasMONIT reports whether reporting trigger has MONIT bit.
func (i *IE) HasMONIT() bool {
	if len(i.Payload) < 2 {
		return false
	}

	switch i.Type {
	case UsageReportTrigger:
		u8 := uint8(i.Payload[1])
		return has5thBit(u8)
	default:
		return false
	}
}

// HasTERMR reports whether reporting trigger has TERMR bit.
func (i *IE) HasTERMR() bool {
	if len(i.Payload) < 2 {
		return false
	}

	switch i.Type {
	case UsageReportTrigger:
		u8 := uint8(i.Payload[1])
		return has4thBit(u8)
	default:
		return false
	}
}
