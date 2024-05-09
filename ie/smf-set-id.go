// Copyright 2019-2024 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewSMFSetID creates a new SMFSetID IE.
func NewSMFSetID(id string) *IE {
	return newFQDNIE(SMFSetID, id)
}

// SMFSetID returns SMFSetID in string if the type of IE matches.
func (i *IE) SMFSetID() (string, error) {
	if i.Type != SMFSetID {
		return "", &InvalidTypeError{Type: i.Type}
	}

	return i.ValueAsFQDN()
}
