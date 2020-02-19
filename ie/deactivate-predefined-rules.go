// Copyright 2019-2020 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewDeactivatePredefinedRules creates a new DeactivatePredefinedRules IE.
func NewDeactivatePredefinedRules(name string) *IE {
	return newStringIE(DeactivatePredefinedRules, name)
}

// DeactivatePredefinedRules returns DeactivatePredefinedRules in string if the type of IE matches.
func (i *IE) DeactivatePredefinedRules() (string, error) {
	if i.Type != DeactivatePredefinedRules {
		return "", &InvalidTypeError{Type: i.Type}
	}

	return string(i.Payload), nil
}
