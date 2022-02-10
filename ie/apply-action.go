// Copyright 2019-2022 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import "io"

// NewApplyAction creates a new ApplyAction IE.
func NewApplyAction(flags uint16) *IE {
	return newUint16ValIE(ApplyAction, flags)
}

// ApplyAction returns ApplyAction in uint16 if the type of IE matches.
func (i *IE) ApplyAction() (uint16, error) {
	if len(i.Payload) < 1 {
		return 0, io.ErrUnexpectedEOF
	}

	switch i.Type {
	case ApplyAction:
		// If the size of the payload is less than two octets because the original was formatted before
		// 3GPP TS 29.244 V16.3.0, MBSU, FSSM, DDPN, BDPN, and EDRT are set to "0".
		if len(i.Payload) < 2 {
			return (uint16(i.Payload[0]) << 8), nil
		}
		return (uint16(i.Payload[0]) << 8) | uint16(i.Payload[1]), nil
	case CreateFAR:
		ies, err := i.CreateFAR()
		if err != nil {
			return 0, err
		}
		for _, x := range ies {
			if x.Type == ApplyAction {
				return x.ApplyAction()
			}
		}
		return 0, ErrIENotFound
	case UpdateFAR:
		ies, err := i.UpdateFAR()
		if err != nil {
			return 0, err
		}
		for _, x := range ies {
			if x.Type == ApplyAction {
				return x.ApplyAction()
			}
		}
		return 0, ErrIENotFound
	default:
		return 0, &InvalidTypeError{Type: i.Type}
	}
}

// HasDROP reports whether an IE has DROP bit.
func (i *IE) HasDROP() bool {
	if len(i.Payload) < 1 {
		return false
	}

	switch i.Type {
	case ApplyAction, DataStatus:
		return has1stBit(i.Payload[0])
	default:
		return false
	}
}

// HasFORW reports whether an IE has FORW bit.
func (i *IE) HasFORW() bool {
	v, err := i.ApplyAction()
	if err != nil {
		return false
	}

	return has2ndBit(uint8((v & 0xFF00) >> 8))
}

// HasBUFF reports whether an IE has BUFF bit.
func (i *IE) HasBUFF() bool {
	if len(i.Payload) < 1 {
		return false
	}

	switch i.Type {
	case ApplyAction:
		return has3rdBit(i.Payload[0])
	case DataStatus:
		return has2ndBit(i.Payload[0])
	default:
		return false
	}
}

// HasNOCP reports whether an IE has NOCP bit.
func (i *IE) HasNOCP() bool {
	v, err := i.ApplyAction()
	if err != nil {
		return false
	}

	return has4thBit(uint8((v & 0xFF00) >> 8))
}

// HasDUPL reports whether an IE has DUPL bit.
func (i *IE) HasDUPL() bool {
	v, err := i.ApplyAction()
	if err != nil {
		return false
	}

	return has5thBit(uint8((v & 0xFF00) >> 8))
}

// HasIPMA reports wether an IE has IPMA bit.
// This flag has been introduced in release 16.2
func (i *IE) HasIPMA() bool {
	v, err := i.ApplyAction()
	if err != nil {
		return false
	}

	return has6thBit(uint8((v & 0xFF00) >> 8))
}

// HasIPMD reports wether an IE has IPMD bit.
// This flag has been introduced in release 16.2
func (i *IE) HasIPMD() bool {
	v, err := i.ApplyAction()
	if err != nil {
		return false
	}

	return has7thBit(uint8((v & 0xFF00) >> 8))
}

// HasDFRT reports wether an IE has DFRT bit.
// This flag has been introduced in release 16.3
func (i *IE) HasDFRT() bool {
	v, err := i.ApplyAction()
	if err != nil {
		return false
	}

	return has8thBit(uint8((v & 0xFF00) >> 8))
}

// HasEDRT reports wether an IE has EDRT bit.
// This flag has been introduced in release 16.3
func (i *IE) HasEDRT() bool {
	v, err := i.ApplyAction()
	if err != nil {
		return false
	}

	return has1stBit(uint8(v & 0x00FF))
}

// HasBDPN reports wether an IE has BDPN bit.
// This flag has been introduced in release 16.4
func (i *IE) HasBDPN() bool {
	v, err := i.ApplyAction()
	if err != nil {
		return false
	}

	return has2ndBit(uint8(v & 0x00FF))
}

// HasDDPN reports wether an IE has DDPN bit.
// This flag has been introduced in release 16.4
func (i *IE) HasDDPN() bool {
	v, err := i.ApplyAction()
	if err != nil {
		return false
	}

	return has3rdBit(uint8(v & 0x00FF))
}

// HasFSSM reports wether an IE has FSSM bit.
// This flag has been introduced in release 17.2
func (i *IE) HasFSSM() bool {
	v, err := i.ApplyAction()
	if err != nil {
		return false
	}

	return has4thBit(uint8(v & 0x00FF))
}

// HasMBSU reports wether an IE has MBSU bit.
// This flag has been introduced in release 17.2
func (i *IE) HasMBSU() bool {
	v, err := i.ApplyAction()
	if err != nil {
		return false
	}

	return has5thBit(uint8(v & 0x00FF))
}
