// Copyright 2019-2024 go-pfcp authors. All rights reserved.
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
	if len(i.Payload) < 2 {
		return nil, io.ErrUnexpectedEOF
	}

	switch i.Type {
	case UsageReportTrigger:
		return i.Payload, nil
	case UsageReportWithinSessionModificationResponse,
		UsageReportWithinSessionDeletionResponse,
		UsageReportWithinSessionReportRequest:
		ies, err := i.UsageReport()
		if err != nil {
			return nil, err
		}
		for _, x := range ies {
			if x.Type == UsageReportTrigger {
				return x.UsageReportTrigger()
			}
		}
		return nil, ErrIENotFound
	default:
		return nil, &InvalidTypeError{Type: i.Type}
	}
}

// HasIMMER reports whether an IE has IMMER bit.
func (i *IE) HasIMMER() bool {
	switch i.Type {
	case UsageReportWithinSessionModificationResponse,
		UsageReportWithinSessionDeletionResponse,
		UsageReportWithinSessionReportRequest,
		UsageReportTrigger:
		v, err := i.UsageReportTrigger()
		if err != nil {
			return false
		}
		return has8thBit(v[0])
	default:
		return false
	}
}

// HasMONIT reports whether an IE has MONIT bit.
func (i *IE) HasMONIT() bool {
	switch i.Type {
	case UsageReportWithinSessionModificationResponse,
		UsageReportWithinSessionDeletionResponse,
		UsageReportWithinSessionReportRequest,
		UsageReportTrigger:
		v, err := i.UsageReportTrigger()
		if err != nil {
			return false
		}
		return has5thBit(v[1])
	default:
		return false
	}
}

// HasTERMR reports whether an IE has TERMR bit.
func (i *IE) HasTERMR() bool {
	switch i.Type {
	case UsageReportWithinSessionModificationResponse,
		UsageReportWithinSessionDeletionResponse,
		UsageReportWithinSessionReportRequest,
		UsageReportTrigger:
		v, err := i.UsageReportTrigger()
		if err != nil {
			return false
		}
		return has4thBit(v[1])
	default:
		return false
	}
}

// HasEMRRE reports whether an IE has EMRRE bit.
func (i *IE) HasEMRRE() bool {
	switch i.Type {
	case UsageReportWithinSessionModificationResponse,
		UsageReportWithinSessionDeletionResponse,
		UsageReportWithinSessionReportRequest,
		UsageReportTrigger:
		v, err := i.UsageReportTrigger()
		if err != nil {
			return false
		}
		if len(v) < 3 {
			// The 3rd byte only appears in R16 or newer R15
			// This is for backward-compatibility with older R15
			return false
		}
		return has5thBit(v[2])
	default:
		return false
	}
}

// HasTEBUR reports whether an IE has TEBUR bit.
func (i *IE) HasTEBUR() bool {
	switch i.Type {
	case UsageReportWithinSessionModificationResponse,
		UsageReportWithinSessionDeletionResponse,
		UsageReportWithinSessionReportRequest,
		UsageReportTrigger:
		v, err := i.UsageReportTrigger()
		if err != nil {
			return false
		}
		if len(v) < 3 {
			// The 3rd byte only appears in R16 or newer R15
			// This is for backward-compatibility with older R15
			return false
		}
		return has2ndBit(v[2])
	default:
		return false
	}
}
