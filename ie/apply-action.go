// Copyright 2019-2020 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewApplyAction creates a new ApplyAction IE.
func NewApplyAction(flag uint8) *IE {
	return newUint8ValIE(ApplyAction, flag)
}

// ApplyAction returns ApplyAction in []byte if the type of IE matches.
func (i *IE) ApplyAction() ([]byte, error) {
	if i.Type != ApplyAction {
		return nil, &InvalidTypeError{Type: i.Type}
	}

	return i.Payload, nil
}

// HasDROP reports whether apply action has DROP bit.
func (i *IE) HasDROP() bool {
	if i.Type != ApplyAction {
		return false
	}
	if len(i.Payload) < 1 {
		return false
	}

	return has1stBit(i.Payload[0])
}

// HasFORW reports whether apply action has FORW bit.
func (i *IE) HasFORW() bool {
	if i.Type != ApplyAction {
		return false
	}
	if len(i.Payload) < 1 {
		return false
	}

	return has2ndBit(i.Payload[0])
}

// HasBUFF reports whether apply action has BUFF bit.
func (i *IE) HasBUFF() bool {
	if i.Type != ApplyAction {
		return false
	}
	if len(i.Payload) < 1 {
		return false
	}

	return has3rdBit(i.Payload[0])
}

// HasNOCP reports whether apply action has NOCP bit.
func (i *IE) HasNOCP() bool {
	if i.Type != ApplyAction {
		return false
	}
	if len(i.Payload) < 1 {
		return false
	}

	return has4thBit(i.Payload[0])
}

// HasDUPL reports whether apply action has DUPL bit.
func (i *IE) HasDUPL() bool {
	if i.Type != ApplyAction {
		return false
	}
	if len(i.Payload) < 1 {
		return false
	}

	return has5thBit(i.Payload[0])
}
