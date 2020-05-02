// Copyright 2019-2020 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewFramedIPv6Route creates a new FramedIPv6Route IE.
func NewFramedIPv6Route(name string) *IE {
	return newStringIE(FramedIPv6Route, name)
}

// FramedIPv6Route returns FramedIPv6Route in string if the type of IE matches.
func (i *IE) FramedIPv6Route() (string, error) {
	if i.Type != FramedIPv6Route {
		return "", &InvalidTypeError{Type: i.Type}
	}

	return string(i.Payload), nil
}
