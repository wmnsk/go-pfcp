// Copyright 2019-2024 go-pfcp authors. All rights reserved.
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
	switch i.Type {
	case ApplyAction:
		if len(i.Payload) < 1 {
			return nil, io.ErrUnexpectedEOF
		}
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

// ValidateApplyAction can be used to facilitate the detection of some inconsistencies in Apply Action flags.
// Its use is optional because validation could also be done on upper layers, or completely skipped for testing purposes.
func (i *IE) ValidateApplyAction() error {
	if i.Type != ApplyAction {
		return &InvalidTypeError{Type: i.Type}
	}
	// One and only one of the DROP, FORW, BUFF, IPMA and IPMD flags shall be set to "1".
	flags := []bool{i.HasDROP(), i.HasFORW(), i.HasBUFF(), i.HasIPMA(), i.HasIPMD()}
	counter := 0
	for _, v := range flags {
		if v {
			counter++
		}
	}
	if counter != 1 {
		return ErrMalformed
	}
	// The NOCP flag and BDPN flag may only be set if the BUFF flag is set.
	if (i.HasNOCP() || i.HasBDPN()) && !i.HasBUFF() {
		return ErrMalformed
	}
	// The DUPL flag may be set with any of the DROP, FORW, BUFF and NOCP flags.
	if i.HasDUPL() && !(i.HasDROP() || i.HasFORW() || i.HasBUFF() || i.HasNOCP()) {
		return ErrMalformed
	}
	// The DFRT flag may only be set if the FORW flag is set.
	// Note: in TS 29.244 V18.0.1 (most recent as of writing), there is a typo and DFRN is stated instead of DFRT
	if i.HasDFRT() && !i.HasFORW() {
		return ErrMalformed
	}
	// The DDPN flag may be set with any of the DROP and BUFF flags.
	if i.HasDDPN() && !(i.HasDROP() || i.HasBUFF()) {
		return ErrMalformed
	}

	// Note: The following is also stated in TS 29.244 section 8.2.26 V18.0.1,
	// but since "may" is used and not "may only"
	// it cannot be used to check for inconsistent IEs:
	// - The EDRT flag may be set if the FORW flag is set.
	// - Both the MBSU flag and the FSSM flag may be set […]
	return nil
}

// HasDROP reports whether an IE has DROP bit.
func (i *IE) HasDROP() bool {
	switch i.Type {
	case CreateFAR, UpdateFAR, ApplyAction:
		v, err := i.ApplyAction()
		if err != nil {
			return false
		}
		return has1stBit(v[0])
	case DownlinkDataReport, DataStatus:
		v, err := i.DataStatus()
		if err != nil {
			return false
		}
		return has1stBit(v)
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
	return has2ndBit(v[0])
}

// HasBUFF reports whether an IE has BUFF bit.
func (i *IE) HasBUFF() bool {
	switch i.Type {
	case CreateFAR, UpdateFAR, ApplyAction:
		v, err := i.ApplyAction()
		if err != nil {
			return false
		}
		return has3rdBit(v[0])
	case DownlinkDataReport, DataStatus:
		v, err := i.DataStatus()
		if err != nil {
			return false
		}
		return has2ndBit(v)
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
	return has4thBit(v[0])
}

// HasDUPL reports whether an IE has DUPL bit.
func (i *IE) HasDUPL() bool {
	v, err := i.ApplyAction()
	if err != nil {
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
	return has6thBit(v[0])
}

// HasIPMD reports wether an IE has IPMD bit.
// This flag has been introduced in release 16.2
func (i *IE) HasIPMD() bool {
	v, err := i.ApplyAction()
	if err != nil {
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
