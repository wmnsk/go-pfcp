// Copyright 2019-2022 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewPartialFailureInformation creates a new PartialFailureInformation IE.
func NewPartialFailureInformation(ies ...*IE) *IE {
	return newGroupedIE(PartialFailureInformation, 0, ies...)
}

// PartialFailureInformation returns the IEs above PartialFailureInformation if the type of IE matches.
func (i *IE) PartialFailureInformation() ([]*IE, error) {
	if (i.Type != PartialFailureInformation) &&
		(i.Type != PartialFailureInformationSessionEstablishmentResponse) &&
		(i.Type != PartialFailureInformationSessionModificationResponse) {
		return nil, &InvalidTypeError{Type: i.Type}
	}

	return ParseMultiIEs(i.Payload)
}
