// Copyright 2019-2020 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewNetworkInstance creates a new NetworkInstance IE.
func NewNetworkInstance(instance string) *IE {
	return newStringIE(NetworkInstance, instance)
}

// NetworkInstance returns NetworkInstance in string if the type of IE matches.
func (i *IE) NetworkInstance() (string, error) {
	if i.Type != NetworkInstance {
		return "", &InvalidTypeError{Type: i.Type}
	}

	return string(i.Payload), nil
}
