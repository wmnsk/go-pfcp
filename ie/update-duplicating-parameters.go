// Copyright 2019-2020 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewUpdateDuplicatingParameters creates a new UpdateDuplicatingParameters IE.
func NewUpdateDuplicatingParameters(di, ohc, tlm, fp *IE) *IE {
	return newGroupedIE(UpdateDuplicatingParameters, 0, di, ohc, tlm, fp)
}

// UpdateDuplicatingParameters returns the IEs above UpdateDuplicatingParameters if the type of IE matches.
func (i *IE) UpdateDuplicatingParameters() ([]*IE, error) {
	if i.Type != UpdateDuplicatingParameters {
		return nil, &InvalidTypeError{Type: i.Type}
	}

	return ParseMultiIEs(i.Payload)
}
