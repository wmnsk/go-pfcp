// Copyright 2019-2020 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewEthernetTrafficInformation creates a new EthernetTrafficInformation IE.
func NewEthernetTrafficInformation(detected, removed *IE) *IE {
	return newGroupedIE(EthernetTrafficInformation, 0, detected, removed)
}

// EthernetTrafficInformation returns the IEs above EthernetTrafficInformation if the type of IE matches.
func (i *IE) EthernetTrafficInformation() ([]*IE, error) {
	if i.Type != EthernetTrafficInformation {
		return nil, &InvalidTypeError{Type: i.Type}
	}

	return ParseMultiIEs(i.Payload)
}
