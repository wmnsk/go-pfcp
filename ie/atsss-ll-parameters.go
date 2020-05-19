// Copyright 2019-2020 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewATSSSLLParameters creates a new ATSSSLLParameters IE.
func NewATSSSLLParameters(info *IE) *IE {
	return newGroupedIE(ATSSSLLParameters, 0, info)
}

// ATSSSLLParameters returns the IEs above ATSSSLLParameters if the type of IE matches.
func (i *IE) ATSSSLLParameters() ([]*IE, error) {
	if i.Type != ATSSSLLParameters {
		return nil, &InvalidTypeError{Type: i.Type}
	}

	return ParseMultiIEs(i.Payload)
}
