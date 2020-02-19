// Copyright 2019-2020 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewSxAssociationReleaseRequest creates a new SxAssociationReleaseRequest IE.
func NewSxAssociationReleaseRequest(sarr, urss int) *IE {
	return newUint8ValIE(SxAssociationReleaseRequest, uint8((urss<<1)|(sarr)))
}

// SxAssociationReleaseRequest returns SxAssociationReleaseRequest in uint8 if the type of IE matches.
func (i *IE) SxAssociationReleaseRequest() (uint8, error) {
	if i.Type != SxAssociationReleaseRequest {
		return 0, &InvalidTypeError{Type: i.Type}
	}

	return i.Payload[0], nil
}

// HasURSS reports whether SxAssociationReleaseRequest IE has URSS bit.
func (i *IE) HasURSS() bool {
	v, err := i.SxAssociationReleaseRequest()
	if err != nil {
		return false
	}

	return has2ndBit(v)
}

// HasSARR reports whether SxAssociationReleaseRequest IE has SARR bit.
func (i *IE) HasSARR() bool {
	v, err := i.SxAssociationReleaseRequest()
	if err != nil {
		return false
	}

	return has1stBit(v)
}
