// Copyright 2019-2020 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewApplicationDetectionInformation creates a new ApplicationDetectionInformation IE.
func NewApplicationDetectionInformation(appID, instID, flowInfo, pdrID *IE) *IE {
	return newGroupedIE(ApplicationDetectionInformation, 0, appID, instID, flowInfo, pdrID)
}

// ApplicationDetectionInformation returns the IEs above ApplicationDetectionInformation if the type of IE matches.
func (i *IE) ApplicationDetectionInformation() ([]*IE, error) {
	if i.Type != ApplicationDetectionInformation {
		return nil, &InvalidTypeError{Type: i.Type}
	}

	return ParseMultiIEs(i.Payload)
}
