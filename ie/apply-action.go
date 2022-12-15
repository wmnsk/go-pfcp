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
	if len(i.Payload) < 1 {
		return false
	}

	switch i.Type {
	case ApplyAction:
		return has2ndBit(i.Payload[0])
	default:
		return false
	}
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
	if len(i.Payload) < 1 {
		return false
	}

	switch i.Type {
	case ApplyAction:
		return has4thBit(i.Payload[0])
	default:
		return false
	}
}

// HasDUPL reports whether an IE has DUPL bit.
func (i *IE) HasDUPL() bool {
	if len(i.Payload) < 1 {
		return false
	}

	switch i.Type {
	case ApplyAction:
		return has5thBit(i.Payload[0])
	default:
		return false
	}
}

// HasIPMA reports wether an IE has IPMA bit.
// This flag has been introduced in release 16.2
func (i *IE) HasIPMA() bool {
	if len(i.Payload) < 1 {
		return false
	}

	switch i.Type {
	case ApplyAction:
		return has6thBit(i.Payload[0])
	default:
		return false
	}
}

// HasIPMD reports wether an IE has IPMD bit.
// This flag has been introduced in release 16.2
func (i *IE) HasIPMD() bool {
	if len(i.Payload) < 1 {
		return false
	}

	switch i.Type {
	case ApplyAction:
		return has7thBit(i.Payload[0])
	default:
		return false
	}
}

// HasDFRT reports wether an IE has DFRT bit.
// This flag has been introduced in release 16.3
func (i *IE) HasDFRT() bool {
	if len(i.Payload) < 1 {
		return false
	}

	switch i.Type {
	case ApplyAction:
		return has8thBit(i.Payload[0])
	default:
		return false
	}
}

// HasEDRT reports wether an IE has EDRT bit.
// This flag has been introduced in release 16.3
func (i *IE) HasEDRT() bool {
	if len(i.Payload) < 2 {
		return false
	}

	switch i.Type {
	case ApplyAction:
		return has1stBit(i.Payload[1])
	default:
		return false
	}
}

// HasBDPN reports wether an IE has BDPN bit.
// This flag has been introduced in release 16.4
func (i *IE) HasBDPN() bool {
	if len(i.Payload) < 2 {
		return false
	}

	switch i.Type {
	case ApplyAction:
		return has2ndBit(i.Payload[1])
	default:
		return false
	}
}

// HasDDPN reports wether an IE has DDPN bit.
// This flag has been introduced in release 16.4
func (i *IE) HasDDPN() bool {
	if len(i.Payload) < 2 {
		return false
	}

	switch i.Type {
	case ApplyAction:
		return has3rdBit(i.Payload[1])
	default:
		return false
	}
}

// HasFSSM reports wether an IE has FSSM bit.
// This flag has been introduced in release 17.2
func (i *IE) HasFSSM() bool {
	if len(i.Payload) < 2 {
		return false
	}

	switch i.Type {
	case ApplyAction:
		return has4thBit(i.Payload[1])
	default:
		return false
	}
}

// HasMBSU reports wether an IE has MBSU bit.
// This flag has been introduced in release 17.2
func (i *IE) HasMBSU() bool {
	if len(i.Payload) < 2 {
		return false
	}

	switch i.Type {
	case ApplyAction:
		return has5thBit(i.Payload[1])
	default:
		return false
	}
}
