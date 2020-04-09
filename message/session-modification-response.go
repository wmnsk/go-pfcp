// Copyright 2019-2020 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package message

import (
	"github.com/wmnsk/go-pfcp/ie"
)

// SessionModificationResponse is a SessionModificationResponse formed PFCP Header and its IEs above.
type SessionModificationResponse struct {
	*Header
	PDRID *ie.IE
	IEs   []*ie.IE
}

// NewSessionModificationResponse creates a new SessionModificationResponse.
func NewSessionModificationResponse(ts *ie.IE, ies ...*ie.IE) *SessionModificationResponse {
	m := &SessionModificationResponse{
		Header: NewHeader(
			1, 0, 0, 0,
			MsgTypeSessionModificationResponse, 0, 0, 0,
			nil,
		),
		PDRID: ts,
		IEs:   ies,
	}
	m.SetLength()

	return m
}

// Marshal returns the byte sequence generated from a SessionModificationResponse.
func (m *SessionModificationResponse) Marshal() ([]byte, error) {
	b := make([]byte, m.MarshalLen())
	if err := m.MarshalTo(b); err != nil {
		return nil, err
	}

	return b, nil
}

// MarshalTo puts the byte sequence in the byte array given as b.
func (m *SessionModificationResponse) MarshalTo(b []byte) error {
	if m.Header.Payload != nil {
		m.Header.Payload = nil
	}
	m.Header.Payload = make([]byte, m.MarshalLen()-m.Header.MarshalLen())

	offset := 0
	if i := m.PDRID; i != nil {
		if err := i.MarshalTo(m.Payload[offset:]); err != nil {
			return err
		}
		offset += i.MarshalLen()
	}

	for _, ie := range m.IEs {
		if ie == nil {
			continue
		}
		if err := ie.MarshalTo(m.Header.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}

	m.Header.SetLength()
	return m.Header.MarshalTo(b)
}

// ParseSessionModificationResponse decodes a given byte sequence as a SessionModificationResponse.
func ParseSessionModificationResponse(b []byte) (*SessionModificationResponse, error) {
	m := &SessionModificationResponse{}
	if err := m.UnmarshalBinary(b); err != nil {
		return nil, err
	}
	return m, nil
}

// UnmarshalBinary decodes a given byte sequence as a SessionModificationResponse.
func (m *SessionModificationResponse) UnmarshalBinary(b []byte) error {
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
		//case ie.PDRID:
		//	m.PDRID = i
		default:
			m.IEs = append(m.IEs, i)
		}
	}

	return nil
}

// MarshalLen returns the serial length of Data.
func (m *SessionModificationResponse) MarshalLen() int {
	l := m.Header.MarshalLen() - len(m.Header.Payload)

	if i := m.PDRID; i != nil {
		l += i.MarshalLen()
	}

	for _, ie := range m.IEs {
		if ie == nil {
			continue
		}
		l += ie.MarshalLen()
	}

	return l
}

// SetLength sets the length in Length field.
func (m *SessionModificationResponse) SetLength() {
	l := m.Header.MarshalLen() - len(m.Header.Payload) - 4

	if i := m.PDRID; i != nil {
		l += i.MarshalLen()
	}

	for _, ie := range m.IEs {
		l += ie.MarshalLen()
	}
	m.Header.Length = uint16(l)
}

// MessageTypeName returns the name of protocol.
func (m *SessionModificationResponse) MessageTypeName() string {
	return "Session Modification Response"
}

// SEID returns the SEID in uint64.
func (m *SessionModificationResponse) SEID() uint64 {
	return m.Header.seid()
}
