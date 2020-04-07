// Copyright 2019-2020 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package message

import (
	"github.com/wmnsk/go-pfcp/ie"
)

// PFDManagementRequest is a PFDManagementRequest formed PFCP Header and its IEs above.
type PFDManagementRequest struct {
	*Header
	ApplicationIDsPFDs *ie.IE
}

// NewPFDManagementRequest creates a new PFDManagementRequest.
func NewPFDManagementRequest(ts *ie.IE, ies ...*ie.IE) *PFDManagementRequest {
	m := &PFDManagementRequest{
		Header: NewHeader(
			1, 0, 0, 0,
			MsgTypePFDManagementRequest, 0, 0, 0,
			nil,
		),
		ApplicationIDsPFDs: ts,
	}
	m.SetLength()

	return m
}

// Marshal returns the byte sequence generated from a PFDManagementRequest.
func (m *PFDManagementRequest) Marshal() ([]byte, error) {
	b := make([]byte, m.MarshalLen())
	if err := m.MarshalTo(b); err != nil {
		return nil, err
	}

	return b, nil
}

// MarshalTo puts the byte sequence in the byte array given as b.
func (m *PFDManagementRequest) MarshalTo(b []byte) error {
	if m.Header.Payload != nil {
		m.Header.Payload = nil
	}
	m.Header.Payload = make([]byte, m.MarshalLen()-m.Header.MarshalLen())

	offset := 0
	if i := m.ApplicationIDsPFDs; i != nil {
		if err := i.MarshalTo(m.Payload[offset:]); err != nil {
			return err
		}
		offset += i.MarshalLen()
	}

	m.Header.SetLength()
	return m.Header.MarshalTo(b)
}

// ParsePFDManagementRequest decodes a given byte sequence as a PFDManagementRequest.
func ParsePFDManagementRequest(b []byte) (*PFDManagementRequest, error) {
	m := &PFDManagementRequest{}
	if err := m.UnmarshalBinary(b); err != nil {
		return nil, err
	}
	return m, nil
}

// UnmarshalBinary decodes a given byte sequence as a PFDManagementRequest.
func (m *PFDManagementRequest) UnmarshalBinary(b []byte) error {
	var err error
	m.Header, err = ParseHeader(b)
	if err != nil {
		return err
	}
	if len(m.Header.Payload) < 2 {
		return nil
	}

	ies, err := ie.ParseMultiIEs(m.Header.Payload)
	if err != nil {
		return err
	}

	for _, i := range ies {
		switch i.Type {
		case ie.ApplicationIDsPFDs:
			m.ApplicationIDsPFDs = i
		}
	}

	return nil
}

// MarshalLen returns the serial length of Data.
func (m *PFDManagementRequest) MarshalLen() int {
	l := m.Header.MarshalLen() - len(m.Header.Payload)

	if i := m.ApplicationIDsPFDs; i != nil {
		l += i.MarshalLen()
	}

	return l
}

// SetLength sets the length in Length field.
func (m *PFDManagementRequest) SetLength() {
	l := m.Header.MarshalLen() - len(m.Header.Payload) - 4

	if i := m.ApplicationIDsPFDs; i != nil {
		l += i.MarshalLen()
	}

	m.Header.Length = uint16(l)
}

// MessageTypeName returns the name of protocol.
func (m *PFDManagementRequest) MessageTypeName() string {
	return "PFD Management Request"
}

// SEID returns the SEID in uint64.
func (m *PFDManagementRequest) SEID() uint64 {
	return m.Header.seid()
}
