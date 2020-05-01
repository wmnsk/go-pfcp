// Copyright 2019-2020 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewAggregatedURRs creates a new AggregatedURRs IE.
func NewAggregatedURRs(id, multiplier *IE) *IE {
	return newGroupedIE(AggregatedURRs, 0, id, multiplier)
}

// AggregatedURRs returns the IEs above AggregatedURRs if the type of IE matches.
func (i *IE) AggregatedURRs() ([]*IE, error) {
	if i.Type != AggregatedURRs {
		return nil, &InvalidTypeError{Type: i.Type}
	}

	return ParseMultiIEs(i.Payload)
}
