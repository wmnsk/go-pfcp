// Copyright 2019-2020 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewQoSInformationInGTPUPathQoSReport creates a new QoSInformationInGTPUPathQoSReport IE.
func NewQoSInformationInGTPUPathQoSReport(avgDelay, minDelay, maxDelay, dscp *IE) *IE {
	return newGroupedIE(QoSInformationInGTPUPathQoSReport, 0, avgDelay, minDelay, maxDelay, dscp)
}

// QoSInformationInGTPUPathQoSReport returns the IEs above QoSInformationInGTPUPathQoSReport if the type of IE matches.
func (i *IE) QoSInformationInGTPUPathQoSReport() ([]*IE, error) {
	if i.Type != QoSInformationInGTPUPathQoSReport {
		return nil, &InvalidTypeError{Type: i.Type}
	}

	return ParseMultiIEs(i.Payload)
}
