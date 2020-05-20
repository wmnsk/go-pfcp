// Copyright 2019-2020 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewDuplicatingParameters creates a new DuplicatingParameters IE.
func NewDuplicatingParameters(di, ohc, tlm, fp *IE) *IE {
	return newGroupedIE(DuplicatingParameters, 0, di, ohc, tlm, fp)
}

// DuplicatingParameters returns the IEs above DuplicatingParameters if the type of IE matches.
func (i *IE) DuplicatingParameters() ([]*IE, error) {
	if i.Type != DuplicatingParameters {
		return nil, &InvalidTypeError{Type: i.Type}
	}

	return ParseMultiIEs(i.Payload)
}
