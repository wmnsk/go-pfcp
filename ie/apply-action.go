// Copyright 2019-2022 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import "io"

// NewApplyAction creates a new ApplyAction IE.
func NewApplyAction(flagsOctets ...uint8) *IE {
	return New(ApplyAction, flagsOctets)
}

// ApplyAction returns ApplyAction in []byte if the type of IE matches.
func (i *IE) ApplyAction() ([]byte, error) {
	if len(i.Payload) < 1 {
		return nil, io.ErrUnexpectedEOF
	}

	switch i.Type {
	case ApplyAction:
		return i.Payload, nil
	case CreateFAR:
		ies, err := i.CreateFAR()
		if err != nil {
			return nil, err
		}
		for _, x := range ies {
			if x.Type == ApplyAction {
				return x.ApplyAction()
			}
		}
		return nil, ErrIENotFound
	case UpdateFAR:
		ies, err := i.UpdateFAR()
		if err != nil {
			return nil, err
		}
		for _, x := range ies {
			if x.Type == ApplyAction {
				return x.ApplyAction()
			}
		}
		return nil, ErrIENotFound
	default:
		return nil, &InvalidTypeError{Type: i.Type}
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
	if len(v) < 1 {
		return false
	}
	return has2ndBit(v[0])
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
	if len(v) < 1 {
		return false
	}
	return has4thBit(v[0])
}

// HasDUPL reports whether an IE has DUPL bit.
func (i *IE) HasDUPL() bool {
	v, err := i.ApplyAction()
	if err != nil {
		return false
	}
	if len(v) < 1 {
		return false
	}
	return has5thBit(v[0])
}

// HasIPMA reports wether an IE has IPMA bit.
// This flag has been introduced in release 16.2
func (i *IE) HasIPMA() bool {
	v, err := i.ApplyAction()
	if err != nil {
		return false
	}
	if len(v) < 1 {
		return false
	}
	return has6thBit(v[0])
}

// HasIPMD reports wether an IE has IPMD bit.
// This flag has been introduced in release 16.2
func (i *IE) HasIPMD() bool {
	v, err := i.ApplyAction()
	if err != nil {
		return false
	}
	if len(v) < 1 {
		return false
	}
	return has7thBit(v[0])
}

// HasDFRT reports wether an IE has DFRT bit.
// This flag has been introduced in release 16.3
func (i *IE) HasDFRT() bool {
	v, err := i.ApplyAction()
	if err != nil {
		return false
	}
	if len(v) < 1 {
		return false
	}
	return has8thBit(v[0])
}

// HasEDRT reports wether an IE has EDRT bit.
// This flag has been introduced in release 16.3
func (i *IE) HasEDRT() bool {
	v, err := i.ApplyAction()
	if err != nil {
		return false
	}
	if len(v) < 2 {
		return false
	}
	return has1stBit(v[1])
}

// HasBDPN reports wether an IE has BDPN bit.
// This flag has been introduced in release 16.4
func (i *IE) HasBDPN() bool {
	v, err := i.ApplyAction()
	if err != nil {
		return false
	}
	if len(v) < 2 {
		return false
	}
	return has2ndBit(v[1])
}

// HasDDPN reports wether an IE has DDPN bit.
// This flag has been introduced in release 16.4
func (i *IE) HasDDPN() bool {
	v, err := i.ApplyAction()
	if err != nil {
		return false
	}
	if len(v) < 2 {
		return false
	}
	return has3rdBit(v[1])
}

// HasFSSM reports wether an IE has FSSM bit.
// This flag has been introduced in release 17.2
func (i *IE) HasFSSM() bool {
	v, err := i.ApplyAction()
	if err != nil {
		return false
	}
	if len(v) < 2 {
		return false
	}
	return has4thBit(v[1])
}

// HasMBSU reports wether an IE has MBSU bit.
// This flag has been introduced in release 17.2
func (i *IE) HasMBSU() bool {
	v, err := i.ApplyAction()
	if err != nil {
		return false
	}
	if len(v) < 2 {
		return false
	}
	return has5thBit(v[1])
}
