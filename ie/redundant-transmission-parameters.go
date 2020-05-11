// Copyright 2019-2020 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewRedundantTransmissionParameters creates a new RedundantTransmissionParameters IE.
func NewRedundantTransmissionParameters(fteid, ohc, ni *IE) *IE {
	return newGroupedIE(RedundantTransmissionParameters, 0, fteid, ohc, ni)
}

// NewRedundantTransmissionParametersInPDI creates a new RedundantTransmissionParameters IE.
func NewRedundantTransmissionParametersInPDI(fteid, ni *IE) *IE {
	return newGroupedIE(RedundantTransmissionParameters, 0, fteid, ni)
}

// NewRedundantTransmissionParametersInFAR creates a new RedundantTransmissionParameters IE.
func NewRedundantTransmissionParametersInFAR(ohc, ni *IE) *IE {
	return newGroupedIE(RedundantTransmissionParameters, 0, ohc, ni)
}

// RedundantTransmissionParameters returns the IEs above RedundantTransmissionParameters if the type of IE matches.
func (i *IE) RedundantTransmissionParameters() ([]*IE, error) {
	if i.Type != RedundantTransmissionParameters {
		return nil, &InvalidTypeError{Type: i.Type}
	}

	return ParseMultiIEs(i.Payload)
}
