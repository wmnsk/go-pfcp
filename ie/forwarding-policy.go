// Copyright 2019-2020 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import "io"

// NewForwardingPolicy creates a new ForwardingPolicy IE.
func NewForwardingPolicy(id string) *IE {
	l := len([]byte(id))
	i := New(ForwardingPolicy, make([]byte, 1+l))

	i.Payload[0] = uint8(l)
	copy(i.Payload[1:], []byte(id))

	return i
}

// ForwardingPolicy returns ForwardingPolicy in []byte if the type of IE matches.
func (i *IE) ForwardingPolicy() ([]byte, error) {
	if i.Type != ForwardingPolicy {
		return nil, &InvalidTypeError{Type: i.Type}
	}

	return i.Payload, nil
}

// ForwardingPolicyIdentifier returns ForwardingPolicyIdentifier in string if the type of IE matches.
func (i *IE) ForwardingPolicyIdentifier() (string, error) {
	v, err := i.ForwardingPolicy()
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
