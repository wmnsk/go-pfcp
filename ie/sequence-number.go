// Copyright 2019-2024 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewSequenceNumber creates a new SequenceNumber IE.
func NewSequenceNumber(seq uint32) *IE {
	return newUint32ValIE(SequenceNumber, seq)
}

// SequenceNumber returns SequenceNumber in uint32 if the type of IE matches.
func (i *IE) SequenceNumber() (uint32, error) {
	switch i.Type {
	case SequenceNumber:
		return i.ValueAsUint32()
	case LoadControlInformation:
		ies, err := i.LoadControlInformation()
		if err != nil {
			return 0, err
		}
		for _, x := range ies {
			if x.Type == SequenceNumber {
				return x.SequenceNumber()
			}
		}
		return 0, ErrIENotFound
	case OverloadControlInformation:
		ies, err := i.OverloadControlInformation()
		if err != nil {
			return 0, err
		}
		for _, x := range ies {
			if x.Type == SequenceNumber {
				return x.SequenceNumber()
			}
		}
		return 0, ErrIENotFound
	default:
		return 0, &InvalidTypeError{Type: i.Type}
	}

}
