// Copyright 2019-2020 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewAdditionalMonitoringTime creates a new AdditionalMonitoringTime IE.
func NewAdditionalMonitoringTime(mTime, svolTh, stimeTh, svolQt, stimeQt, eTh, eQt *IE) *IE {
	return newGroupedIE(AdditionalMonitoringTime, 0, mTime, svolTh, stimeTh, svolQt, stimeQt, eTh, eQt)
}

// AdditionalMonitoringTime returns the IEs above AdditionalMonitoringTime if the type of IE matches.
func (i *IE) AdditionalMonitoringTime() ([]*IE, error) {
	if i.Type != AdditionalMonitoringTime {
		return nil, &InvalidTypeError{Type: i.Type}
	}

	return ParseMultiIEs(i.Payload)
}
