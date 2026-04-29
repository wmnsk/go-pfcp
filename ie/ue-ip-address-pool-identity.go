// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import (
	"encoding/binary"
	"io"
)

// NewUEIPAddressPoolIdentity creates a new UEIPAddressPoolIdentity IE.
func NewUEIPAddressPoolIdentity(id string) *IE {
	l := len([]byte(id))
	i := New(UEIPAddressPoolIdentity, make([]byte, 2+l))

	binary.BigEndian.PutUint16(i.Payload[0:2], uint16(l))
	copy(i.Payload[2:], id)

	return i
}

// UEIPAddressPoolIdentity returns UEIPAddressPoolIdentity in []byte if the type of IE matches.
func (i *IE) UEIPAddressPoolIdentity() ([]byte, error) {
	if len(i.Payload) < 1 {
		return nil, io.ErrUnexpectedEOF
	}

	switch i.Type {
	case UEIPAddressPoolIdentity:
		return i.Payload, nil
	case CreatePDR:
		ies, err := i.CreatePDR()
		if err != nil {
			return nil, err
		}
		for _, x := range ies {
			if x.Type == UEIPAddressPoolIdentity {
				return x.UEIPAddressPoolIdentity()
			}
		}
		return nil, ErrIENotFound
	case UEIPAddressPoolInformation:
		ies, err := i.UEIPAddressPoolInformation()
		if err != nil {
			return nil, err
		}
		for _, x := range ies {
			if x.Type == UEIPAddressPoolIdentity {
				return x.UEIPAddressPoolIdentity()
			}
		}
		return nil, ErrIENotFound
	default:
		return nil, &InvalidTypeError{Type: i.Type}
	}
}

// UEIPAddressPoolIdentityString returns UEIPAddressPoolIdentity in string if the type of IE matches.
func (i *IE) UEIPAddressPoolIdentityString() (string, error) {
	v, err := i.UEIPAddressPoolIdentity()
	if err != nil {
		return "", err
	}

	if len(v) < 2 {
		return "", io.ErrUnexpectedEOF
	}

	idlen := int(binary.BigEndian.Uint16(v[0:2]))
	if len(v) < idlen+2 {
		return "", io.ErrUnexpectedEOF
	}

	return string(v[2 : idlen+2]), nil
}
