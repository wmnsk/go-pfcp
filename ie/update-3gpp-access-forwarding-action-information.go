// Copyright 2019-2020 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewUpdateTGPPAccessForwardingActionInformation creates a new UpdateTGPPAccessForwardingActionInformation IE.
func NewUpdateTGPPAccessForwardingActionInformation(farID, weight, priority, urrID *IE) *IE {
	return newGroupedIE(UpdateTGPPAccessForwardingActionInformation, 0, farID, weight, priority, urrID)
}

// UpdateTGPPAccessForwardingActionInformation returns the IEs above UpdateTGPPAccessForwardingActionInformation if the type of IE matches.
func (i *IE) UpdateTGPPAccessForwardingActionInformation() ([]*IE, error) {
	if i.Type != UpdateTGPPAccessForwardingActionInformation {
		return nil, &InvalidTypeError{Type: i.Type}
	}

	return ParseMultiIEs(i.Payload)
}
