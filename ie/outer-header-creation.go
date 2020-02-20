// Copyright 2019-2020 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import (
	"encoding/binary"
	"io"
	"net"
)

// NewOuterHeaderCreation creates a new OuterHeaderCreation IE.
func NewOuterHeaderCreation(flags uint16, teid uint32, v4, v6 string, port uint16, ctag, stag uint32) *IE {
	fields := NewOuterHeaderCreationFields(flags, teid, v4, v6, port, ctag, stag)
	b, err := fields.Marshal()
	if err != nil {
		return nil
	}

	return New(OuterHeaderCreation, b)
}

// OuterHeaderCreation returns OuterHeaderCreation in *OuterHeaderCreationFields if the type of IE matches.
func (i *IE) OuterHeaderCreation() (*OuterHeaderCreationFields, error) {
	if i.Type != OuterHeaderCreation {
		return nil, &InvalidTypeError{Type: i.Type}
	}

	f, err := ParseOuterHeaderCreationFields(i.Payload)
	if err != nil {
		return nil, err
	}
	return f, nil
}

// HasTEID reports whether OuterHeaderCreation IE has DLVOL bit.
func (i *IE) HasTEID() bool {
	if i.Type != OuterHeaderCreation {
		return false
	}
	if len(i.Payload) < 1 {
		return false
	}

	return has1stBit(i.Payload[0]) || has2ndBit(i.Payload[0])
}

// HasIPv4 reports whether OuterHeaderCreation IE has DLVOL bit.
func (i *IE) HasIPv4() bool {
	switch i.Type {
	case OuterHeaderCreation:
		if len(i.Payload) < 1 {
			return false
		}

		return has1stBit(i.Payload[0]) || has3rdBit(i.Payload[0]) || has5thBit(i.Payload[0])
	case UEIPAddress:
		if len(i.Payload) < 1 {
			return false
		}

		return has2ndBit(i.Payload[0]) && !has5thBit(i.Payload[0])
	default:
		return false
	}
}

// HasIPv6 reports whether OuterHeaderCreation IE has DLVOL bit.
func (i *IE) HasIPv6() bool {
	switch i.Type {
	case OuterHeaderCreation:
		if len(i.Payload) < 1 {
			return false
		}

		return has2ndBit(i.Payload[0]) || has4thBit(i.Payload[0]) || has6thBit(i.Payload[0])
	case UEIPAddress:
		if len(i.Payload) < 1 {
			return false
		}

		return has1stBit(i.Payload[0]) && !has5thBit(i.Payload[0])
	default:
		return false
	}
}

// HasCTag reports whether OuterHeaderCreation IE has DLVOL bit.
func (i *IE) HasCTag() bool {
	if i.Type != OuterHeaderCreation {
		return false
	}
	if len(i.Payload) < 1 {
		return false
	}

	return has7thBit(i.Payload[0])
}

// HasSTag reports whether OuterHeaderCreation IE has DLVOL bit.
func (i *IE) HasSTag() bool {
	if i.Type != OuterHeaderCreation {
		return false
	}
	if len(i.Payload) < 1 {
		return false
	}

	return has8thBit(i.Payload[0])
}

// IsN19 reports whether OuterHeaderCreation IE has DLVOL bit.
func (i *IE) IsN19() bool {
	if i.Type != OuterHeaderCreation {
		return false
	}
	if len(i.Payload) < 2 {
		return false
	}

	return has1stBit(i.Payload[1])
}

// IsN6 reports whether OuterHeaderCreation IE has DLVOL bit.
func (i *IE) IsN6() bool {
	if i.Type != OuterHeaderCreation {
		return false
	}
	if len(i.Payload) < 2 {
		return false
	}

	return has2ndBit(i.Payload[1])
}

// OuterHeaderCreationFields represents a fields contained in OuterHeaderCreation IE.
type OuterHeaderCreationFields struct {
	Flags       uint16
	TEID        uint32
	IPv4Address net.IP
	IPv6Address net.IP
	PortNumber  uint16
	CTag        uint32
	STag        uint32
}

func NewOuterHeaderCreationFields(flags uint16, teid uint32, v4, v6 string, port uint16, ctag, stag uint32) *OuterHeaderCreationFields {
	f := &OuterHeaderCreationFields{Flags: flags}

	if has1stBit(uint8(flags&0xff)) || has2ndBit(uint8(flags&0xff)) {
		f.TEID = teid
	}

	if has1stBit(uint8(flags&0xff)) || has3rdBit(uint8(flags&0xff)) || has5thBit(uint8(flags&0xff)) {
		f.IPv4Address = net.ParseIP(v4).To4()
	}

	if has2ndBit(uint8(flags&0xff)) || has4thBit(uint8(flags&0xff)) || has6thBit(uint8(flags&0xff)) {
		f.IPv6Address = net.ParseIP(v6).To16()
	}

	if has3rdBit(uint8(flags&0xff)) || has4thBit(uint8(flags&0xff)) {
		f.PortNumber = port
	}

	if has7thBit(uint8(flags & 0xff)) {
		f.CTag = ctag
	}

	if has8thBit(uint8(flags & 0xff)) {
		f.STag = stag
	}

	return f
}

// ParseOuterHeaderCreationFields parses b into OuterHeaderCreationFields.
func ParseOuterHeaderCreationFields(b []byte) (*OuterHeaderCreationFields, error) {
	f := &OuterHeaderCreationFields{}
	if err := f.UnmarshalBinary(b); err != nil {
		return nil, err
	}
	return f, nil
}

// UnmarshalBinary parses b into IE.
func (f *OuterHeaderCreationFields) UnmarshalBinary(b []byte) error {
	l := len(b)
	if l < 2 {
		return io.ErrUnexpectedEOF
	}

	f.Flags = binary.BigEndian.Uint16(b[0:2])
	offset := 3

	if has1stBit(uint8(f.Flags&0xff)) || has2ndBit(uint8(f.Flags&0xff)) {
		if l < offset+4 {
			return io.ErrUnexpectedEOF
		}
		f.TEID = binary.BigEndian.Uint32(b[offset : offset+4])
		offset += 4
	}

	if has1stBit(uint8(f.Flags&0xff)) || has3rdBit(uint8(f.Flags&0xff)) || has5thBit(uint8(f.Flags&0xff)) {
		if l < offset+4 {
			return io.ErrUnexpectedEOF
		}
		f.IPv4Address = net.IP(b[offset : offset+4]).To4()
		offset += 4
	}

	if has2ndBit(uint8(f.Flags&0xff)) || has4thBit(uint8(f.Flags&0xff)) || has6thBit(uint8(f.Flags&0xff)) {
		if l < offset+16 {
			return io.ErrUnexpectedEOF
		}
		f.IPv6Address = net.IP(b[offset : offset+16]).To16()
		offset += 16
	}

	if has3rdBit(uint8(f.Flags&0xff)) || has4thBit(uint8(f.Flags&0xff)) {
		if l < offset+2 {
			return io.ErrUnexpectedEOF
		}
		f.PortNumber = binary.BigEndian.Uint16(b[offset : offset+2])
		offset += 2
	}

	if has7thBit(uint8(f.Flags & 0xff)) {
		if l < offset+3 {
			return io.ErrUnexpectedEOF
		}
		f.CTag = binary.BigEndian.Uint32(b[offset : offset+3])
		offset += 3
	}

	if has8thBit(uint8(f.Flags & 0xff)) {
		if l < offset+3 {
			return io.ErrUnexpectedEOF
		}
		f.STag = binary.BigEndian.Uint32(b[offset : offset+3])
	}

	return nil
}

// Marshal returns the serialized bytes of OuterHeaderCreationFields.
func (f *OuterHeaderCreationFields) Marshal() ([]byte, error) {
	b := make([]byte, f.MarshalLen())
	if err := f.MarshalTo(b); err != nil {
		return nil, err
	}
	return b, nil
}

// MarshalTo puts the byte sequence in the byte array given as b.
func (f *OuterHeaderCreationFields) MarshalTo(b []byte) error {
	l := len(b)
	if l < 2 {
		return io.ErrUnexpectedEOF
	}

	binary.BigEndian.PutUint16(b[0:2], f.Flags)
	offset := 2

	if has1stBit(uint8(f.Flags&0xff)) || has2ndBit(uint8(f.Flags&0xff)) {
		binary.BigEndian.PutUint32(b[offset:offset+4], f.TEID)
		offset += 4
	}

	if has1stBit(uint8(f.Flags&0xff)) || has3rdBit(uint8(f.Flags&0xff)) || has5thBit(uint8(f.Flags&0xff)) {
		copy(b[offset:offset+4], f.IPv4Address)
		offset += 4
	}

	if has2ndBit(uint8(f.Flags&0xff)) || has4thBit(uint8(f.Flags&0xff)) || has6thBit(uint8(f.Flags&0xff)) {
		copy(b[offset:offset+16], f.IPv6Address)
		offset += 16
	}

	if has3rdBit(uint8(f.Flags&0xff)) || has4thBit(uint8(f.Flags&0xff)) {
		binary.BigEndian.PutUint16(b[offset:offset+2], f.PortNumber)
	}

	if has7thBit(uint8(f.Flags & 0xff)) {
		p := make([]byte, 4)
		binary.BigEndian.PutUint32(p, f.CTag)
		copy(b[offset:offset+3], p[1:4])
		offset += 3
	}

	if has8thBit(uint8(f.Flags & 0xff)) {
		p := make([]byte, 4)
		binary.BigEndian.PutUint32(p, f.STag)
		copy(b[offset:offset+3], p[1:4])
	}

	return nil
}

// MarshalLen returns field length in integer.
func (f *OuterHeaderCreationFields) MarshalLen() int {
	l := 2
	if has1stBit(uint8(f.Flags&0xff)) || has2ndBit(uint8(f.Flags&0xff)) {
		l += 4
	}
	if has1stBit(uint8(f.Flags&0xff)) || has3rdBit(uint8(f.Flags&0xff)) || has5thBit(uint8(f.Flags&0xff)) {
		l += 4
	}
	if has2ndBit(uint8(f.Flags&0xff)) || has4thBit(uint8(f.Flags&0xff)) || has6thBit(uint8(f.Flags&0xff)) {
		l += 16
	}
	if has3rdBit(uint8(f.Flags&0xff)) || has4thBit(uint8(f.Flags&0xff)) {
		l += 2
	}
	if has7thBit(uint8(f.Flags & 0xff)) {
		l += 3
	}
	if has8thBit(uint8(f.Flags & 0xff)) {
		l += 3
	}

	return l
}
