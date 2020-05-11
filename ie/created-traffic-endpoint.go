// Copyright 2019-2020 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewCreatedTrafficEndpoint creates a new CreatedTrafficEndpoint IE.
func NewCreatedTrafficEndpoint(id, fteid1, fteid2, ueIP *IE) *IE {
	return newGroupedIE(CreatedTrafficEndpoint, 0, id, fteid1, fteid2, ueIP)
}

// CreatedTrafficEndpoint returns the IEs above CreatedTrafficEndpoint if the type of IE matches.
func (i *IE) CreatedTrafficEndpoint() ([]*IE, error) {
	if i.Type != CreatedTrafficEndpoint {
		return nil, &InvalidTypeError{Type: i.Type}
	}

	return ParseMultiIEs(i.Payload)
}
