// Copyright 2019-2020 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewTGPPAccessForwardingActionInformation creates a new TGPPAccessForwardingActionInformation IE.
func NewTGPPAccessForwardingActionInformation(farID, weight, priority, urrID *IE) *IE {
	return newGroupedIE(TGPPAccessForwardingActionInformation, 0, farID, weight, priority, urrID)
}

// TGPPAccessForwardingActionInformation returns the IEs above TGPPAccessForwardingActionInformation if the type of IE matches.
func (i *IE) TGPPAccessForwardingActionInformation() ([]*IE, error) {
	if i.Type != TGPPAccessForwardingActionInformation {
		return nil, &InvalidTypeError{Type: i.Type}
	}

	return ParseMultiIEs(i.Payload)
}
