// Copyright 2019-2024 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewPrecedence creates a new Precedence IE.
func NewPrecedence(id uint32) *IE {
	return newUint32ValIE(Precedence, id)
}

// Precedence returns Precedence in uint32 if the type of IE matches.
func (i *IE) Precedence() (uint32, error) {
	switch i.Type {
	case Precedence:
		return i.ValueAsUint32()
	case CreatePDR:
		ies, err := i.CreatePDR()
		if err != nil {
			return 0, err
		}
		for _, x := range ies {
			if x.Type == Precedence {
				return x.Precedence()
			}
		}
		return 0, ErrIENotFound
	case UpdatePDR:
		ies, err := i.UpdatePDR()
		if err != nil {
			return 0, err
		}
		for _, x := range ies {
			if x.Type == Precedence {
				return x.Precedence()
			}
		}
		return 0, ErrIENotFound
	default:
		return 0, &InvalidTypeError{Type: i.Type}
	}
}
