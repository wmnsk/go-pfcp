// Copyright 2019-2020 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewApplicationInstanceID creates a new ApplicationInstanceID IE.
func NewApplicationInstanceID(id string) *IE {
	return newStringIE(ApplicationInstanceID, id)
}

// ApplicationInstanceID returns ApplicationInstanceID in string if the type of IE matches.
func (i *IE) ApplicationInstanceID() (string, error) {
	if i.Type != ApplicationInstanceID {
		return "", &InvalidTypeError{Type: i.Type}
	}

	return string(i.Payload), nil
}
