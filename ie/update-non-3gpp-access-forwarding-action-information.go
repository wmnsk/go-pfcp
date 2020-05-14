// Copyright 2019-2020 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewUpdateNonTGPPAccessForwardingActionInformation creates a new UpdateNonTGPPAccessForwardingActionInformation IE.
func NewUpdateNonTGPPAccessForwardingActionInformation(farID, weight, priority, urrID *IE) *IE {
	return newGroupedIE(UpdateNonTGPPAccessForwardingActionInformation, 0, farID, weight, priority, urrID)
}

// UpdateNonTGPPAccessForwardingActionInformation returns the IEs above UpdateNonTGPPAccessForwardingActionInformation if the type of IE matches.
func (i *IE) UpdateNonTGPPAccessForwardingActionInformation() ([]*IE, error) {
	if i.Type != UpdateNonTGPPAccessForwardingActionInformation {
		return nil, &InvalidTypeError{Type: i.Type}
	}

	return ParseMultiIEs(i.Payload)
}
