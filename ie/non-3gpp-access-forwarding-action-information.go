// Copyright 2019-2020 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewNonTGPPAccessForwardingActionInformation creates a new NonTGPPAccessForwardingActionInformation IE.
func NewNonTGPPAccessForwardingActionInformation(farID, weight, priority, urrID *IE) *IE {
	return newGroupedIE(NonTGPPAccessForwardingActionInformation, 0, farID, weight, priority, urrID)
}

// NonTGPPAccessForwardingActionInformation returns the IEs above NonTGPPAccessForwardingActionInformation if the type of IE matches.
func (i *IE) NonTGPPAccessForwardingActionInformation() ([]*IE, error) {
	if i.Type != NonTGPPAccessForwardingActionInformation {
		return nil, &InvalidTypeError{Type: i.Type}
	}

	return ParseMultiIEs(i.Payload)
}
