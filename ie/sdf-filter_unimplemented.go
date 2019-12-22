// Copyright 2019 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewSDFFilter creates a new SDFFilter IE.
func NewSDFFilter() *IE {
	return New(SDFFilter, []byte{})
}

// SDFFilter returns SDFFilter in structured format if the type of IE matches.
func (i *IE) SDFFilter() (*SDFFilterFields, error) {
	if i.Type != SDFFilter {
		return nil, &InvalidTypeError{Type: i.Type}
	}

	s, err := ParseSDFFilterFields(i.Payload)
	if err != nil {
		return nil, err
	}
	return s, nil
}

// SDFFilterFields represents a fields contained in SDFFilter IE.
//
// DO NOT USE THIS: This IE is not fully implemented yet.
type SDFFilterFields struct {
	Flags                  uint8
	FDLength               uint16
	FlowDescription        string
	ToSTrafficClass        uint16
	SecurityParameterIndex uint32
	FlowLabel              uint32 // 3 octets
	SDFFilterID            uint32
}

// NewSDFFilterFields creates a new NewSDFFilterFields.
func NewSDFFilterFields() *SDFFilterFields {
	return &SDFFilterFields{}
}

// HasBID reports whether CHID flag is set.
func (f *SDFFilterFields) HasBID() bool {
	return has5thBit(f.Flags)
}

// SetBIDFlag sets CHID flag in SDFFilter.
func (f *SDFFilterFields) SetBIDFlag() {
	f.Flags |= 0x10
}

// HasFL reports whether CHID flag is set.
func (f *SDFFilterFields) HasFL() bool {
	return has4thBit(f.Flags)
}

// SetFLFlag sets CHID flag in SDFFilter.
func (f *SDFFilterFields) SetFLFlag() {
	f.Flags |= 0x08
}

// HasSPI reports whether CH flag is set.
func (f *SDFFilterFields) HasSPI() bool {
	return has3rdBit(f.Flags)
}

// SetSPIFlag sets CH flag in SDFFilter.
func (f *SDFFilterFields) SetSPIFlag() {
	f.Flags |= 0x04
}

// HasTTC reports whether TTC flag is set.
func (f *SDFFilterFields) HasTTC() bool {
	return has2ndBit(f.Flags)
}

// SetTTCFlag sets CHID flag in SDFFilter.
func (f *SDFFilterFields) SetTTCFlag() {
	f.Flags |= 0x02
}

// HasFD reports whether FD flag is set.
func (f *SDFFilterFields) HasFD() bool {
	return has1stBit(f.Flags)
}

// SetFDFlag sets CHID flag in SDFFilter.
func (f *SDFFilterFields) SetFDFlag() {
	f.Flags |= 0x01
}

// ParseSDFFilterFields parses b into SDFFilterFields.
func ParseSDFFilterFields(b []byte) (*SDFFilterFields, error) {
	f := &SDFFilterFields{}
	if err := f.UnmarshalBinary(b); err != nil {
		return nil, err
	}
	return f, nil
}

// UnmarshalBinary parses b into IE.
func (f *SDFFilterFields) UnmarshalBinary(b []byte) error {
	l := len(b)
	if l < 2 {
		return ErrTooShortToParse
	}

	f.Flags = b[0]
	//offset := 1

	return nil
}

// Marshal returns the serialized bytes of SDFFilterFields.
func (f *SDFFilterFields) Marshal() ([]byte, error) {
	b := make([]byte, f.MarshalLen())
	if err := f.MarshalTo(b); err != nil {
		return nil, err
	}
	return b, nil
}

// MarshalTo puts the byte sequence in the byte array given as b.
func (f *SDFFilterFields) MarshalTo(b []byte) error {
	l := len(b)
	if l < 1 {
		return ErrTooShortToParse
	}

	b[0] = f.Flags
	//offset := 1

	return nil
}

// MarshalLen returns field length in integer.
func (f *SDFFilterFields) MarshalLen() int {
	return 0
}
