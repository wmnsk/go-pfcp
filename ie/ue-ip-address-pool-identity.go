// Copyright 2019-2020 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import "io"

// NewUEIPAddressPoolIdentity creates a new UEIPAddressPoolIdentity IE.
func NewUEIPAddressPoolIdentity(id string) *IE {
	l := len([]byte(id))
	i := New(UEIPAddressPoolIdentity, make([]byte, 1+l))

	i.Payload[0] = uint8(l)
	copy(i.Payload[1:], []byte(id))

	return i
}

// UEIPAddressPoolIdentity returns UEIPAddressPoolIdentity in []byte if the type of IE matches.
func (i *IE) UEIPAddressPoolIdentity() ([]byte, error) {
	if i.Type != UEIPAddressPoolIdentity {
		return nil, &InvalidTypeError{Type: i.Type}
	}

	return i.Payload, nil
}

// UEIPAddressPoolIdentityIdentifier returns UEIPAddressPoolIdentityIdentifier in string if the type of IE matches.
func (i *IE) UEIPAddressPoolIdentityIdentifier() (string, error) {
	v, err := i.UEIPAddressPoolIdentity()
	if err != nil {
		return "", err
	}

	l := len(v)
	if l < 1 {
		return "", io.ErrUnexpectedEOF
	}

	idlen := int(v[0])
	if l < idlen+1 {
		return "", io.ErrUnexpectedEOF
	}

	return string(v[1:idlen]), nil
}
