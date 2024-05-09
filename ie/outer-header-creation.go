// Copyright 2019-2024 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import (
	"encoding/binary"
	"io"
	"net"
)

// NewOuterHeaderCreation creates a new OuterHeaderCreation IE.
func NewOuterHeaderCreation(desc uint16, teid uint32, v4, v6 string, port uint16, ctag, stag uint32) *IE {
	fields := NewOuterHeaderCreationFields(desc, teid, v4, v6, port, ctag, stag)
	b, err := fields.Marshal()
	if err != nil {
		return nil
	}

	return New(OuterHeaderCreation, b)
}

// OuterHeaderCreation returns OuterHeaderCreation in *OuterHeaderCreationFields if the type of IE matches.
func (i *IE) OuterHeaderCreation() (*OuterHeaderCreationFields, error) {
	switch i.Type {
	case OuterHeaderCreation:
		f, err := ParseOuterHeaderCreationFields(i.Payload)
		if err != nil {
			return nil, err
		}
		return f, nil
	case ForwardingParameters:
		ies, err := i.ForwardingParameters()
		if err != nil {
			return nil, err
		}
		for _, x := range ies {
			if x.Type == OuterHeaderCreation {
				return x.OuterHeaderCreation()
			}
		}
		return nil, ErrIENotFound
	case UpdateForwardingParameters:
		ies, err := i.UpdateForwardingParameters()
		if err != nil {
			return nil, err
		}
		for _, x := range ies {
			if x.Type == OuterHeaderCreation {
				return x.OuterHeaderCreation()
			}
		}
		return nil, ErrIENotFound
	case DuplicatingParameters:
		ies, err := i.DuplicatingParameters()
		if err != nil {
			return nil, err
		}
		for _, x := range ies {
			if x.Type == OuterHeaderCreation {
				return x.OuterHeaderCreation()
			}
		}
		return nil, ErrIENotFound
	case UpdateDuplicatingParameters:
		ies, err := i.UpdateDuplicatingParameters()
		if err != nil {
			return nil, err
		}
		for _, x := range ies {
			if x.Type == OuterHeaderCreation {
				return x.OuterHeaderCreation()
			}
		}
		return nil, ErrIENotFound
	case RedundantTransmissionParameters:
		ies, err := i.RedundantTransmissionParameters()
		if err != nil {
			return nil, err
		}
		for _, x := range ies {
			if x.Type == OuterHeaderCreation {
				return x.OuterHeaderCreation()
			}
		}
		return nil, ErrIENotFound
	default:
		return nil, &InvalidTypeError{Type: i.Type}
	}
}

// HasTEID reports whether an IE has TEID bit.
func (i *IE) HasTEID() bool {
	if i.Type != OuterHeaderCreation {
		return false
	}
	if len(i.Payload) < 1 {
		return false
	}
	v, err := i.OuterHeaderCreation()
	if err != nil {
		return false
	}

	return v.HasTEID()
}

// HasPortNumber reports wether an IE has Port Number bit.
func (i *IE) HasPortNumber() bool {
	if i.Type != OuterHeaderCreation {
		return false
	}
	if len(i.Payload) < 1 {
		return false
	}
	v, err := i.OuterHeaderCreation()
	if err != nil {
		return false
	}

	return v.HasPortNumber()
}

// HasIPv4 reports whether an IE has IPv4 bit.
func (i *IE) HasIPv4() bool {
	switch i.Type {
	case OuterHeaderCreation:
		if len(i.Payload) < 1 {
			return false
		}
		v, err := i.OuterHeaderCreation()
		if err != nil {
			return false
		}

		return v.HasIPv4()
	case UEIPAddress:
		if len(i.Payload) < 1 {
			return false
		}

		return has2ndBit(i.Payload[0]) && !has5thBit(i.Payload[0])
	case UserPlaneIPResourceInformation:
		if len(i.Payload) < 1 {
			return false
		}

		return has1stBit(i.Payload[0])
	case FSEID:
		v, err := i.FSEID()
		if err != nil {
			return false
		}
		return v.HasIPv4()
	case FTEID:
		v, err := i.FTEID()
		if err != nil {
			return false
		}
		return v.HasIPv4()
	case IPVersion:
		v, err := i.IPVersion()
		if err != nil {
			return false
		}
		return has2ndBit(v)
	default:
		return false
	}
}

// HasIPv6 reports whether an IE has IPv6 bit.
func (i *IE) HasIPv6() bool {
	switch i.Type {
	case OuterHeaderCreation:
		if len(i.Payload) < 1 {
			return false
		}
		v, err := i.OuterHeaderCreation()
		if err != nil {
			return false
		}

		return v.HasIPv6()
	case UEIPAddress:
		if len(i.Payload) < 1 {
			return false
		}

		return has1stBit(i.Payload[0]) && !has5thBit(i.Payload[0])
	case UserPlaneIPResourceInformation:
		if len(i.Payload) < 1 {
			return false
		}

		return has2ndBit(i.Payload[0])
	case FSEID:
		v, err := i.FSEID()
		if err != nil {
			return false
		}
		return v.HasIPv6()
	case FTEID:
		v, err := i.FTEID()
		if err != nil {
			return false
		}
		return v.HasIPv6()
	case IPVersion:
		v, err := i.IPVersion()
		if err != nil {
			return false
		}
		return has1stBit(v)
	default:
		return false
	}
}

// HasCTag reports whether an IE has CTAG bit.
func (i *IE) HasCTag() bool {
	if i.Type != OuterHeaderCreation {
		return false
	}
	if len(i.Payload) < 1 {
		return false
	}
	v, err := i.OuterHeaderCreation()
	if err != nil {
		return false
	}

	return v.HasCTag()
}

// HasSTag reports whether an IE has STAG bit.
func (i *IE) HasSTag() bool {
	if i.Type != OuterHeaderCreation {
		return false
	}
	if len(i.Payload) < 1 {
		return false
	}
	v, err := i.OuterHeaderCreation()
	if err != nil {
		return false
	}

	return v.HasSTag()
}

// IsN19 reports whether an IE has N19 bit.
func (i *IE) IsN19() bool {
	if i.Type != OuterHeaderCreation {
		return false
	}
	if len(i.Payload) < 2 {
		return false
	}
	v, err := i.OuterHeaderCreation()
	if err != nil {
		return false
	}

	return v.IsN19()
}

// IsN6 reports whether an IE has N6 bit.
func (i *IE) IsN6() bool {
	if i.Type != OuterHeaderCreation {
		return false
	}
	if len(i.Payload) < 2 {
		return false
	}
	v, err := i.OuterHeaderCreation()
	if err != nil {
		return false
	}

	return v.IsN6()
}

// IsLLSSMCTEID reports wether an IE has Low Layer SSM and C-TEID bit.
func (i *IE) IsLLSSMCTEID() bool {
	if i.Type != OuterHeaderCreation {
		return false
	}
	if len(i.Payload) < 2 {
		return false
	}
	v, err := i.OuterHeaderCreation()
	if err != nil {
		return false
	}

	return v.IsLLSSMCTEID()
}

// OuterHeaderCreationFields represents a fields contained in OuterHeaderCreation IE.
type OuterHeaderCreationFields struct {
	OuterHeaderCreationDescription uint16
	TEID                           uint32
	IPv4Address                    net.IP
	IPv6Address                    net.IP
	PortNumber                     uint16
	CTag                           uint32
	STag                           uint32
}

// NewOuterHeaderCreationFields creates a new OuterHeaderCreationFields.
func NewOuterHeaderCreationFields(desc uint16, teid uint32, v4, v6 string, port uint16, ctag, stag uint32) *OuterHeaderCreationFields {
	f := &OuterHeaderCreationFields{OuterHeaderCreationDescription: desc}

	oct5 := uint8((desc & 0xff00) >> 8)

	if has1stBit(oct5) || has2ndBit(oct5) {
		f.TEID = teid
	}

	if has1stBit(oct5) || has3rdBit(oct5) || has5thBit(oct5) {
		f.IPv4Address = net.ParseIP(v4).To4()
	}

	if has2ndBit(oct5) || has4thBit(oct5) || has6thBit(oct5) {
		f.IPv6Address = net.ParseIP(v6).To16()
	}

	if has3rdBit(oct5) || has4thBit(oct5) {
		f.PortNumber = port
	}

	if has7thBit(oct5) {
		f.CTag = ctag
	}

	if has8thBit(oct5) {
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

	f.OuterHeaderCreationDescription = binary.BigEndian.Uint16(b[0:2])
	offset := 2

	oct5 := b[0]
	if has1stBit(oct5) || has2ndBit(oct5) {
		if l < offset+4 {
			return io.ErrUnexpectedEOF
		}
		f.TEID = binary.BigEndian.Uint32(b[offset : offset+4])
		offset += 4
	}

	if has1stBit(oct5) || has3rdBit(oct5) || has5thBit(oct5) {
		if l < offset+4 {
			return io.ErrUnexpectedEOF
		}
		f.IPv4Address = net.IP(b[offset : offset+4]).To4()
		offset += 4
	}

	if has2ndBit(oct5) || has4thBit(oct5) || has6thBit(oct5) {
		if l < offset+16 {
			return io.ErrUnexpectedEOF
		}
		f.IPv6Address = net.IP(b[offset : offset+16]).To16()
		offset += 16
	}

	if has3rdBit(oct5) || has4thBit(oct5) {
		if l < offset+2 {
			return io.ErrUnexpectedEOF
		}
		f.PortNumber = binary.BigEndian.Uint16(b[offset : offset+2])
		offset += 2
	}

	if has7thBit(uint8(f.OuterHeaderCreationDescription & 0xff)) {
		if l < offset+3 {
			return io.ErrUnexpectedEOF
		}
		f.CTag = binary.BigEndian.Uint32(b[offset : offset+3])
		offset += 3
	}

	if has8thBit(uint8(f.OuterHeaderCreationDescription & 0xff)) {
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

	binary.BigEndian.PutUint16(b[0:2], f.OuterHeaderCreationDescription)
	offset := 2

	oct5 := uint8((f.OuterHeaderCreationDescription & 0xff00) >> 8)

	if has1stBit(oct5) || has2ndBit(oct5) {
		binary.BigEndian.PutUint32(b[offset:offset+4], f.TEID)
		offset += 4
	}

	if has1stBit(oct5) || has3rdBit(oct5) || has5thBit(oct5) {
		copy(b[offset:offset+4], f.IPv4Address)
		offset += 4
	}

	if has2ndBit(oct5) || has4thBit(oct5) || has6thBit(oct5) {
		copy(b[offset:offset+16], f.IPv6Address)
		offset += 16
	}

	if has3rdBit(oct5) || has4thBit(oct5) {
		binary.BigEndian.PutUint16(b[offset:offset+2], f.PortNumber)
	}

	if has7thBit(oct5) {
		p := make([]byte, 4)
		binary.BigEndian.PutUint32(p, f.CTag)
		copy(b[offset:offset+3], p[1:4])
		offset += 3
	}

	if has8thBit(oct5) {
		p := make([]byte, 4)
		binary.BigEndian.PutUint32(p, f.STag)
		copy(b[offset:offset+3], p[1:4])
	}

	return nil
}

// MarshalLen returns field length in integer.
func (f *OuterHeaderCreationFields) MarshalLen() int {
	l := 2

	if f.HasTEID() {
		l += 4
	}
	if f.HasIPv4() {
		l += 4
	}
	if f.HasIPv6() {
		l += 16
	}
	if f.HasPortNumber() {
		l += 2
	}
	if f.HasCTag() {
		l += 3
	}
	if f.HasSTag() {
		l += 3
	}

	return l
}

// HasTEID reports wether TEID field is set.
func (f *OuterHeaderCreationFields) HasTEID() bool {
	// The TEID field shall be present
	// if the Outer Header Creation Description requests
	// the creation of aGTP-U header. Otherwise it shall not be present.
	desc := uint8((f.OuterHeaderCreationDescription & 0xff00) >> 8)
	return has1stBit(desc) || has2ndBit(desc)
}

// HasIPv4 reports wether IPv4 Address field is set.
func (f *OuterHeaderCreationFields) HasIPv4() bool {
	// The IPv4 Address field shall be present
	// if the Outer Header Creation Description requests
	// the creation of an IPv4 header. Otherwise it shall not be present.
	desc := uint8((f.OuterHeaderCreationDescription & 0xff00) >> 8)
	return has1stBit(desc) || has3rdBit(desc) || has5thBit(desc)
}

// HasIPv6 reports wether IPv6 Address field is set.
func (f *OuterHeaderCreationFields) HasIPv6() bool {
	// The IPv6 Address field shall be present
	// if the Outer Header Creation Description requests
	// the creation of an IPv6 header. Otherwise it shall not be present.
	desc := uint8((f.OuterHeaderCreationDescription & 0xff00) >> 8)
	return has2ndBit(desc) || has4thBit(desc) || has6thBit(desc)
}

// HasPortNumber reports wether Port Number field is set.
func (f *OuterHeaderCreationFields) HasPortNumber() bool {
	// The Port Number field shall be present
	// if the Outer Header Creation Description requests
	// the creation of a UDP/IP header. Otherwise it shall not be present.
	desc := uint8((f.OuterHeaderCreationDescription & 0xff00) >> 8)
	return has3rdBit(desc) || has4thBit(desc)
}

// HasCTag reports wether C-TAG field is set.
func (f *OuterHeaderCreationFields) HasCTag() bool {
	// The C-TAG field shall be present
	// if the Outer Header Creation Description requests
	// the setting of the C-Tag in Ethernet packet. Otherwise it shall not be present.
	desc := uint8((f.OuterHeaderCreationDescription & 0xff00) >> 8)
	return has7thBit(desc)
}

// HasSTag reports wether S-TAG field is set.
func (f *OuterHeaderCreationFields) HasSTag() bool {
	// The S-TAG field shall be present
	// if the Outer Header Creation Description requests
	// the setting of the S-Tag in Ethernet packet. Otherwise it shall not be present.
	desc := uint8((f.OuterHeaderCreationDescription & 0xff00) >> 8)
	return has8thBit(desc)
}

// IsN19 reports wether Outer Header Creation Description has N19 Indication.
func (f *OuterHeaderCreationFields) IsN19() bool {
	desc := uint8(f.OuterHeaderCreationDescription & 0x00FF)
	return has1stBit(desc)
}

// IsN6 reports wether Outer Header Creation Description has N9 Indication
func (f *OuterHeaderCreationFields) IsN6() bool {
	desc := uint8(f.OuterHeaderCreationDescription & 0x00FF)
	return has2ndBit(desc)
}

// IsLLSSMCTEID reports wether Outer Header Creation Description has Low Layer SSM and C-TEID
// This bit has been introduced in release 17.2
func (f *OuterHeaderCreationFields) IsLLSSMCTEID() bool {
	desc := uint8(f.OuterHeaderCreationDescription & 0x00FF)
	return has3rdBit(desc)
}
