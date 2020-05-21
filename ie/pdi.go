// Copyright 2019-2020 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewPDI creates a new PDI IE.
func NewPDI(srcIF, fteid, ni, rtp, ip, endpoint, sdffilter, appID, ethInfo, ethFilter *IE) *IE {
	return newGroupedIE(PDI, 0, srcIF, fteid, ni, rtp, ip, endpoint, sdffilter, appID, ethInfo, ethFilter)
}

// PDI returns the IEs above PDI if the type of IE matches.
func (i *IE) PDI() ([]*IE, error) {
	if i.Type != PDI {
		return nil, &InvalidTypeError{Type: i.Type}
	}

	return ParseMultiIEs(i.Payload)
}
