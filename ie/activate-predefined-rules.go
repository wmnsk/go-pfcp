// Copyright 2019-2020 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewActivatePredefinedRules creates a new ActivatePredefinedRules IE.
func NewActivatePredefinedRules(name string) *IE {
	return newStringIE(ActivatePredefinedRules, name)
}

// ActivatePredefinedRules returns ActivatePredefinedRules in string if the type of IE matches.
func (i *IE) ActivatePredefinedRules() (string, error) {
	if i.Type != ActivatePredefinedRules {
		return "", &InvalidTypeError{Type: i.Type}
	}

	return string(i.Payload), nil
}
