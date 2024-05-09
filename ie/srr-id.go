// Copyright 2019-2024 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewSRRID creates a new SRRID IE.
func NewSRRID(id uint8) *IE {
	return newUint8ValIE(SRRID, id)
}

// SRRID returns SRRID in uint8 if the type of IE matches.
func (i *IE) SRRID() (uint8, error) {
	switch i.Type {
	case SRRID:
		return i.ValueAsUint8()
	case RemoveSRR:
		ies, err := i.RemoveSRR()
		if err != nil {
			return 0, err
		}

		for _, x := range ies {
			if x.Type == SRRID {
				return x.SRRID()
			}
		}
		return 0, ErrIENotFound
	case CreateSRR:
		ies, err := i.CreateSRR()
		if err != nil {
			return 0, err
		}

		for _, x := range ies {
			if x.Type == SRRID {
				return x.SRRID()
			}
		}
		return 0, ErrIENotFound
	case UpdateSRR:
		ies, err := i.UpdateSRR()
		if err != nil {
			return 0, err
		}

		for _, x := range ies {
			if x.Type == SRRID {
				return x.SRRID()
			}
		}
		return 0, ErrIENotFound
	case SessionReport:
		ies, err := i.SessionReport()
		if err != nil {
			return 0, err
		}

		for _, x := range ies {
			if x.Type == SRRID {
				return x.SRRID()
			}
		}
		return 0, ErrIENotFound
	default:
		return 0, &InvalidTypeError{Type: i.Type}
	}
}
