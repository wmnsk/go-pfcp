// Copyright 2019-2020 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewUpdateForwardingParameters creates a new UpdateForwardingParameters IE.
func NewUpdateForwardingParameters(di, ni, redi, ohc, tlm, fp, he, smflags, ltei, dit, dnai *IE) *IE {
	return newGroupedIE(UpdateForwardingParameters, 0, di, ni, redi, ohc, tlm, fp, he, smflags, ltei, dit, dnai)
}

// UpdateForwardingParameters returns the IEs above UpdateForwardingParameters if the type of IE matches.
func (i *IE) UpdateForwardingParameters() ([]*IE, error) {
	if i.Type != UpdateForwardingParameters {
		return nil, &InvalidTypeError{Type: i.Type}
	}

	return ParseMultiIEs(i.Payload)
}
