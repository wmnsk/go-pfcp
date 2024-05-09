// Copyright 2019-2024 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import (
	"io"
)

// NewReportingTriggers creates a new ReportingTriggers IE.
func NewReportingTriggers(triggersOctets ...uint8) *IE {
	return New(ReportingTriggers, triggersOctets)
}

// ReportingTriggers returns ReportingTriggers in []byte if the type of IE matches.
func (i *IE) ReportingTriggers() ([]byte, error) {
	if len(i.Payload) < 2 {
		return nil, io.ErrUnexpectedEOF
	}

	switch i.Type {
	case ReportingTriggers:
		return i.Payload, nil
	case CreateURR:
		ies, err := i.CreateURR()
		if err != nil {
			return nil, err
		}
		for _, x := range ies {
			if x.Type == ReportingTriggers {
				return x.ReportingTriggers()
			}
		}
		return nil, ErrIENotFound
	case UpdateURR:
		ies, err := i.UpdateURR()
		if err != nil {
			return nil, err
		}
		for _, x := range ies {
			if x.Type == ReportingTriggers {
				return x.ReportingTriggers()
			}
		}
		return nil, ErrIENotFound
	default:
		return nil, &InvalidTypeError{Type: i.Type}
	}
}

// HasLIUSA reports whether an IE has LIUSA bit.
func (i *IE) HasLIUSA() bool {
	switch i.Type {
	case CreateURR, UpdateURR, ReportingTriggers:
		v, err := i.ReportingTriggers()
		if err != nil {
			return false
		}
		return has8thBit(v[0])
	case UsageReportWithinSessionModificationResponse,
		UsageReportWithinSessionDeletionResponse,
		UsageReportWithinSessionReportRequest,
		UsageReportTrigger:
		v, err := i.UsageReportTrigger()
		if err != nil {
			return false
		}
		return has3rdBit(v[1])
	default:
		return false
	}
}

// HasDROTH reports whether an IE has DROTH bit.
func (i *IE) HasDROTH() bool {
	switch i.Type {
	case CreateURR, UpdateURR, ReportingTriggers:
		v, err := i.ReportingTriggers()
		if err != nil {
			return false
		}
		return has7thBit(v[0])
	case UsageReportWithinSessionModificationResponse,
		UsageReportWithinSessionDeletionResponse,
		UsageReportWithinSessionReportRequest,
		UsageReportTrigger:
		v, err := i.UsageReportTrigger()
		if err != nil {
			return false
		}
		return has7thBit(v[0])
	default:
		return false
	}
}

// HasSTOPT reports whether an IE has STOPT bit.
func (i *IE) HasSTOPT() bool {
	switch i.Type {
	case CreateURR, UpdateURR, ReportingTriggers:
		v, err := i.ReportingTriggers()
		if err != nil {
			return false
		}
		return has6thBit(v[0])
	case UsageReportWithinSessionModificationResponse,
		UsageReportWithinSessionDeletionResponse,
		UsageReportWithinSessionReportRequest,
		UsageReportTrigger:
		v, err := i.UsageReportTrigger()
		if err != nil {
			return false
		}
		return has6thBit(v[0])
	default:
		return false
	}
}

// HasSTART reports whether an IE has START bit.
func (i *IE) HasSTART() bool {
	switch i.Type {
	case CreateURR, UpdateURR, ReportingTriggers:
		v, err := i.ReportingTriggers()
		if err != nil {
			return false
		}
		return has5thBit(v[0])
	case UsageReportWithinSessionModificationResponse,
		UsageReportWithinSessionDeletionResponse,
		UsageReportWithinSessionReportRequest,
		UsageReportTrigger:
		v, err := i.UsageReportTrigger()
		if err != nil {
			return false
		}
		return has5thBit(v[0])
	default:
		return false
	}
}

// HasQUHTI reports whether an IE has QUHTI bit.
func (i *IE) HasQUHTI() bool {
	switch i.Type {
	case CreateURR, UpdateURR, ReportingTriggers:
		v, err := i.ReportingTriggers()
		if err != nil {
			return false
		}
		return has4thBit(v[0])
	case UsageReportWithinSessionModificationResponse,
		UsageReportWithinSessionDeletionResponse,
		UsageReportWithinSessionReportRequest,
		UsageReportTrigger:
		v, err := i.UsageReportTrigger()
		if err != nil {
			return false
		}
		return has4thBit(v[0])
	default:
		return false
	}
}

// HasTIMTH reports whether an IE has TIMTH bit.
func (i *IE) HasTIMTH() bool {
	switch i.Type {
	case CreateURR, UpdateURR, ReportingTriggers:
		v, err := i.ReportingTriggers()
		if err != nil {
			return false
		}
		return has3rdBit(v[0])
	case UsageReportWithinSessionModificationResponse,
		UsageReportWithinSessionDeletionResponse,
		UsageReportWithinSessionReportRequest,
		UsageReportTrigger:
		v, err := i.UsageReportTrigger()
		if err != nil {
			return false
		}
		return has3rdBit(v[0])
	default:
		return false
	}
}

// HasVOLTH reports whether an IE has VOLTH bit.
func (i *IE) HasVOLTH() bool {
	switch i.Type {
	case CreateURR, UpdateURR, ReportingTriggers:
		v, err := i.ReportingTriggers()
		if err != nil {
			return false
		}
		return has2ndBit(v[0])
	case UsageReportWithinSessionModificationResponse,
		UsageReportWithinSessionDeletionResponse,
		UsageReportWithinSessionReportRequest,
		UsageReportTrigger:
		v, err := i.UsageReportTrigger()
		if err != nil {
			return false
		}
		return has2ndBit(v[0])
	default:
		return false
	}
}

// HasPERIO reports whether an IE has PERIO bit.
func (i *IE) HasPERIO() bool {
	switch i.Type {
	case CreateURR, UpdateURR, ReportingTriggers:
		v, err := i.ReportingTriggers()
		if err != nil {
			return false
		}
		return has1stBit(v[0])
	case UsageReportWithinSessionModificationResponse,
		UsageReportWithinSessionDeletionResponse,
		UsageReportWithinSessionReportRequest,
		UsageReportTrigger:
		v, err := i.UsageReportTrigger()
		if err != nil {
			return false
		}
		return has1stBit(v[0])
	case QoSMonitoringPerQoSFlowControlInformation,
		ReportingFrequency:
		v, err := i.ReportingFrequency()
		if err != nil {
			return false
		}
		return has2ndBit(v)
	default:
		return false
	}
}

// HasQUVTI reports whether an IE has QUVTI bit.
func (i *IE) HasQUVTI() bool {
	switch i.Type {
	case CreateURR, UpdateURR, ReportingTriggers:
		v, err := i.ReportingTriggers()
		if err != nil {
			return false
		}
		return has8thBit(v[1])
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
		return has4thBit(v[2])
	default:
		return false
	}
}

// HasIPMJL reports whether an IE has IPMJL bit.
func (i *IE) HasIPMJL() bool {
	switch i.Type {
	case CreateURR, UpdateURR, ReportingTriggers:
		v, err := i.ReportingTriggers()
		if err != nil {
			return false
		}
		return has7thBit(v[1])
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
		return has3rdBit(v[2])
	default:
		return false
	}
}

// HasEVEQU reports whether an IE has EVEQU bit.
func (i *IE) HasEVEQU() bool {
	switch i.Type {
	case CreateURR, UpdateURR, ReportingTriggers:
		v, err := i.ReportingTriggers()
		if err != nil {
			return false
		}
		return has6thBit(v[1])
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
		return has1stBit(v[2])
	default:
		return false
	}
}

// HasEVETH reports whether an IE has EVETH bit.
func (i *IE) HasEVETH() bool {
	switch i.Type {
	case CreateURR, UpdateURR, ReportingTriggers:
		v, err := i.ReportingTriggers()
		if err != nil {
			return false
		}
		return has5thBit(v[1])
	case UsageReportWithinSessionModificationResponse,
		UsageReportWithinSessionDeletionResponse,
		UsageReportWithinSessionReportRequest,
		UsageReportTrigger:
		v, err := i.UsageReportTrigger()
		if err != nil {
			return false
		}
		return has8thBit(v[1])
	default:
		return false
	}
}

// HasMACAR reports whether an IE has MACAR bit.
func (i *IE) HasMACAR() bool {
	switch i.Type {
	case CreateURR, UpdateURR, ReportingTriggers:
		v, err := i.ReportingTriggers()
		if err != nil {
			return false
		}
		return has4thBit(v[1])
	case UsageReportWithinSessionModificationResponse,
		UsageReportWithinSessionDeletionResponse,
		UsageReportWithinSessionReportRequest,
		UsageReportTrigger:
		v, err := i.UsageReportTrigger()
		if err != nil {
			return false
		}
		return has7thBit(v[1])
	default:
		return false
	}
}

// HasENVCL reports whether an IE has ENVCL bit.
func (i *IE) HasENVCL() bool {
	switch i.Type {
	case CreateURR, UpdateURR, ReportingTriggers:
		v, err := i.ReportingTriggers()
		if err != nil {
			return false
		}
		return has3rdBit(v[1])
	case UsageReportWithinSessionModificationResponse,
		UsageReportWithinSessionDeletionResponse,
		UsageReportWithinSessionReportRequest,
		UsageReportTrigger:
		v, err := i.UsageReportTrigger()
		if err != nil {
			return false
		}
		return has6thBit(v[1])
	default:
		return false
	}
}

// HasTIMQU reports whether an IE has TIMQU bit.
func (i *IE) HasTIMQU() bool {
	switch i.Type {
	case CreateURR, UpdateURR, ReportingTriggers:
		v, err := i.ReportingTriggers()
		if err != nil {
			return false
		}
		return has2ndBit(v[1])
	case UsageReportWithinSessionModificationResponse,
		UsageReportWithinSessionDeletionResponse,
		UsageReportWithinSessionReportRequest,
		UsageReportTrigger:
		v, err := i.UsageReportTrigger()
		if err != nil {
			return false
		}
		return has2ndBit(v[1])
	default:
		return false
	}
}

// HasVOLQU reports whether an IE has VOLQU bit.
func (i *IE) HasVOLQU() bool {
	switch i.Type {
	case CreateURR, UpdateURR, ReportingTriggers:
		v, err := i.ReportingTriggers()
		if err != nil {
			return false
		}
		return has1stBit(v[1])
	case UsageReportWithinSessionModificationResponse,
		UsageReportWithinSessionDeletionResponse,
		UsageReportWithinSessionReportRequest,
		UsageReportTrigger:
		v, err := i.UsageReportTrigger()
		if err != nil {
			return false
		}
		return has1stBit(v[1])
	default:
		return false
	}
}

// HasUPINT reports whether an IE has UPINT bit.
func (i *IE) HasUPINT() bool {
	switch i.Type {
	case CreateURR, UpdateURR, ReportingTriggers:
		v, err := i.ReportingTriggers()
		if err != nil {
			return false
		}
		if len(v) < 3 {
			// The 3rd byte only appears in R16
			// This is for backward-compatibility with R15
			return false
		}
		return has2ndBit(v[2])
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
		return has6thBit(v[2])
	default:
		return false
	}
}

// HasREEMR reports whether an IE has REEMR bit.
func (i *IE) HasREEMR() bool {
	switch i.Type {
	case CreateURR, UpdateURR, ReportingTriggers:
		v, err := i.ReportingTriggers()
		if err != nil {
			return false
		}
		if len(v) < 3 {
			// The 3rd byte only appears in R16
			// This is for backward-compatibility with R15
			return false
		}
		return has1stBit(v[2])
	default:
		return false
	}
}
