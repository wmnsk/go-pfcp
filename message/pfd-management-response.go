// Copyright 2019-2020 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package message

import (
	"github.com/wmnsk/go-pfcp/ie"
)

// PFDManagementResponse is a PFDManagementResponse formed PFCP Header and its IEs above.
type PFDManagementResponse struct {
	*Header
	OffendingIE *ie.IE
	IEs         []*ie.IE
}

// NewPFDManagementResponse creates a new PFDManagementResponse.
func NewPFDManagementResponse(ts *ie.IE) *PFDManagementResponse {
	m := &PFDManagementResponse{
		Header: NewHeader(
			1, 0, 0, 0,
			MsgTypePFDManagementResponse, 0, 0, 0,
			nil,
		),
		OffendingIE: ts,
	}
	m.SetLength()

	return m
}

// Marshal returns the byte sequence generated from a PFDManagementResponse.
func (m *PFDManagementResponse) Marshal() ([]byte, error) {
	b := make([]byte, m.MarshalLen())
	if err := m.MarshalTo(b); err != nil {
		return nil, err
	}

	return b, nil
}

// MarshalTo puts the byte sequence in the byte array given as b.
func (m *PFDManagementResponse) MarshalTo(b []byte) error {
	if m.Header.Payload != nil {
		m.Header.Payload = nil
	}
	m.Header.Payload = make([]byte, m.MarshalLen()-m.Header.MarshalLen())

	offset := 0
	if i := m.OffendingIE; i != nil {
		if err := i.MarshalTo(m.Payload[offset:]); err != nil {
			return err
		}
		offset += i.MarshalLen()
	}

	m.Header.SetLength()
	return m.Header.MarshalTo(b)
}

// ParsePFDManagementResponse decodes a given byte sequence as a PFDManagementResponse.
func ParsePFDManagementResponse(b []byte) (*PFDManagementResponse, error) {
	m := &PFDManagementResponse{}
	if err := m.UnmarshalBinary(b); err != nil {
		return nil, err
	}
	return m, nil
}

// UnmarshalBinary decodes a given byte sequence as a PFDManagementResponse.
func (m *PFDManagementResponse) UnmarshalBinary(b []byte) error {
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
		case ie.OffendingIE:
			m.OffendingIE = i
		}
	}

	return nil
}

// MarshalLen returns the serial length of Data.
func (m *PFDManagementResponse) MarshalLen() int {
	l := m.Header.MarshalLen() - len(m.Header.Payload)

	if i := m.OffendingIE; i != nil {
		l += i.MarshalLen()
	}

	return l
}

// SetLength sets the length in Length field.
func (m *PFDManagementResponse) SetLength() {
	l := m.Header.MarshalLen() - len(m.Header.Payload) - 4

	if i := m.OffendingIE; i != nil {
		l += i.MarshalLen()
	}

	m.Header.Length = uint16(l)
}

// MessageTypeName returns the name of protocol.
func (m *PFDManagementResponse) MessageTypeName() string {
	return "PFD Management Response"
}

// SEID returns the SEID in uint64.
func (m *PFDManagementResponse) SEID() uint64 {
	return m.Header.seid()
}
