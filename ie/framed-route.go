// Copyright 2019-2020 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewFramedRoute creates a new FramedRoute IE.
func NewFramedRoute(route string) *IE {
	return newStringIE(FramedRoute, route)
}

// FramedRoute returns FramedRoute in string if the type of IE matches.
func (i *IE) FramedRoute() (string, error) {
	if i.Type != FramedRoute {
		return "", &InvalidTypeError{Type: i.Type}
	}

	return string(i.Payload), nil
}
