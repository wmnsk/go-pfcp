// Copyright 2019-2020 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import (
	"encoding/binary"
	"io"
)

// NewDLFlowLevelMarking creates a new DLFlowLevelMarking IE.
func NewDLFlowLevelMarking(flags uint8, ttc, sci uint16) *IE {
	fields := NewDLFlowLevelMarkingFields(flags, ttc, sci)
	b, err := fields.Marshal()
	if err != nil {
		return nil
	}

	return New(DLFlowLevelMarking, b)
}

// DLFlowLevelMarking returns DLFlowLevelMarking in *DLFlowLevelMarkingFields if the type of IE matches.
func (i *IE) DLFlowLevelMarking() (*DLFlowLevelMarkingFields, error) {
	if i.Type != DLFlowLevelMarking {
		return nil, &InvalidTypeError{Type: i.Type}
	}

	f, err := ParseDLFlowLevelMarkingFields(i.Payload)
	if err != nil {
		return nil, err
	}
	return f, nil
}

// HasTTC reports whether DLFlowLevelMarking IE has TTC bit.
func (i *IE) HasTTC() bool {
	if i.Type != DLFlowLevelMarking {
		return false
	}
	if len(i.Payload) < 1 {
		return false
	}

	return has2ndBit(i.Payload[0])
}

// HasSCI reports whether DLFlowLevelMarking IE has SCI bit.
func (i *IE) HasSCI() bool {
	if i.Type != DLFlowLevelMarking {
		return false
	}
	if len(i.Payload) < 1 {
		return false
	}

	return has1stBit(i.Payload[0])
}

// DLFlowLevelMarkingFields represents a fields contained in DLFlowLevelMarking IE.
type DLFlowLevelMarkingFields struct {
	Flags                  uint8
	ToSTrafficClass        uint16
	ServiceClassIdentifier uint16
}

func NewDLFlowLevelMarkingFields(flags uint8, ttc, sci uint16) *DLFlowLevelMarkingFields {
	f := &DLFlowLevelMarkingFields{Flags: flags}

	if has1stBit(flags) {
		f.ToSTrafficClass = ttc
	}

	if has2ndBit(flags) {
		f.ServiceClassIdentifier = sci
	}

	return f
}

// ParseDLFlowLevelMarkingFields parses b into DLFlowLevelMarkingFields.
func ParseDLFlowLevelMarkingFields(b []byte) (*DLFlowLevelMarkingFields, error) {
	f := &DLFlowLevelMarkingFields{}
	if err := f.UnmarshalBinary(b); err != nil {
		return nil, err
	}
	return f, nil
}

// UnmarshalBinary parses b into IE.
func (f *DLFlowLevelMarkingFields) UnmarshalBinary(b []byte) error {
	l := len(b)
	if l < 1 {
		return io.ErrUnexpectedEOF
	}

	f.Flags = b[0]
	offset := 1

	if has1stBit(f.Flags) {
		if l < offset+2 {
			return io.ErrUnexpectedEOF
		}
		f.ToSTrafficClass = binary.BigEndian.Uint16(b[offset : offset+2])
		offset += 2
	}

	if has2ndBit(f.Flags) {
		if l < offset+2 {
			return io.ErrUnexpectedEOF
		}
		f.ServiceClassIdentifier = binary.BigEndian.Uint16(b[offset : offset+2])
	}

	return nil
}

// Marshal returns the serialized bytes of DLFlowLevelMarkingFields.
func (f *DLFlowLevelMarkingFields) Marshal() ([]byte, error) {
	b := make([]byte, f.MarshalLen())
	if err := f.MarshalTo(b); err != nil {
		return nil, err
	}
	return b, nil
}

// MarshalTo puts the byte sequence in the byte array given as b.
func (f *DLFlowLevelMarkingFields) MarshalTo(b []byte) error {
	l := len(b)
	if l < 2 {
		return io.ErrUnexpectedEOF
	}

	b[0] = f.Flags
	offset := 1

	if has1stBit(f.Flags) {
		binary.BigEndian.PutUint16(b[offset:offset+2], f.ToSTrafficClass)
		offset += 2
	}

	if has2ndBit(f.Flags) {
		binary.BigEndian.PutUint16(b[offset:offset+2], f.ServiceClassIdentifier)
	}

	return nil
}

// MarshalLen returns field length in integer.
func (f *DLFlowLevelMarkingFields) MarshalLen() int {
	l := 1
	if has1stBit(f.Flags) {
		l += 2
	}
	if has2ndBit(f.Flags) {
		l += 2
	}

	return l
}
