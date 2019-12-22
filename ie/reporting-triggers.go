// Copyright 2019 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import (
	"encoding/binary"
	"io"
)

// NewReportingTriggers creates a new ReportingTriggers IE.
func NewReportingTriggers(triggers uint16) *IE {
	return newUint16ValIE(ReportingTriggers, triggers)
}

// ReportingTriggers returns ReportingTriggers in uint16 if the type of IE matches.
func (i *IE) ReportingTriggers() (uint16, error) {
	if i.Type != ReportingTriggers {
		return 0, &InvalidTypeError{Type: i.Type}
	}

	if len(i.Payload) < 2 {
		return 0, io.ErrUnexpectedEOF
	}
	return binary.BigEndian.Uint16(i.Payload[0:2]), nil
}

// HasLIUSA reports whether reporting trigger has LIUSA bit.
func (i *IE) HasLIUSA() bool {
	v, err := i.ReportingTriggers()
	if err != nil {
		return false
	}

	u8 := uint8(v >> 8)
	return has8thBit(u8)
}

// HasDROTH reports whether reporting trigger has DROTH bit.
func (i *IE) HasDROTH() bool {
	v, err := i.ReportingTriggers()
	if err != nil {
		return false
	}

	u8 := uint8(v >> 8)
	return has7thBit(u8)
}

// HasSTOPT reports whether reporting trigger has STOPT bit.
func (i *IE) HasSTOPT() bool {
	v, err := i.ReportingTriggers()
	if err != nil {
		return false
	}

	u8 := uint8(v >> 8)
	return has6thBit(u8)
}

// HasSTART reports whether reporting trigger has START bit.
func (i *IE) HasSTART() bool {
	v, err := i.ReportingTriggers()
	if err != nil {
		return false
	}

	u8 := uint8(v >> 8)
	return has5thBit(u8)
}

// HasQUHTI reports whether reporting trigger has QUHTI bit.
func (i *IE) HasQUHTI() bool {
	v, err := i.ReportingTriggers()
	if err != nil {
		return false
	}

	u8 := uint8(v >> 8)
	return has4thBit(u8)
}

// HasTIMTH reports whether reporting trigger has TIMTH bit.
func (i *IE) HasTIMTH() bool {
	v, err := i.ReportingTriggers()
	if err != nil {
		return false
	}

	u8 := uint8(v >> 8)
	return has3rdBit(u8)
}

// HasVOLTH reports whether reporting trigger has VOLTH bit.
func (i *IE) HasVOLTH() bool {
	v, err := i.ReportingTriggers()
	if err != nil {
		return false
	}

	u8 := uint8(v >> 8)
	return has2ndBit(u8)
}

// HasPERIO reports whether reporting trigger has PERIO bit.
func (i *IE) HasPERIO() bool {
	v, err := i.ReportingTriggers()
	if err != nil {
		return false
	}

	u8 := uint8(v >> 8)
	return has1stBit(u8)
}

// HasEVEQU reports whether reporting trigger has EVEQU bit.
func (i *IE) HasEVEQU() bool {
	v, err := i.ReportingTriggers()
	if err != nil {
		return false
	}

	u8 := uint8(v & 0xff)
	return has6thBit(u8)
}

// HasEVETH reports whether reporting trigger has EVETH bit.
func (i *IE) HasEVETH() bool {
	v, err := i.ReportingTriggers()
	if err != nil {
		return false
	}

	u8 := uint8(v & 0xff)
	return has5thBit(u8)
}

// HasMACAR reports whether reporting trigger has MACAR bit.
func (i *IE) HasMACAR() bool {
	v, err := i.ReportingTriggers()
	if err != nil {
		return false
	}

	u8 := uint8(v & 0xff)
	return has4thBit(u8)
}

// HasENVCL reports whether reporting trigger has ENVCL bit.
func (i *IE) HasENVCL() bool {
	v, err := i.ReportingTriggers()
	if err != nil {
		return false
	}

	u8 := uint8(v & 0xff)
	return has3rdBit(u8)
}

// HasTIMQU reports whether reporting trigger has TIMQU bit.
func (i *IE) HasTIMQU() bool {
	v, err := i.ReportingTriggers()
	if err != nil {
		return false
	}

	u8 := uint8(v & 0xff)
	return has2ndBit(u8)
}

// HasVOLQU reports whether reporting trigger has VOLQU bit.
func (i *IE) HasVOLQU() bool {
	v, err := i.ReportingTriggers()
	if err != nil {
		return false
	}

	u8 := uint8(v & 0xff)
	return has1stBit(u8)
}
