// Copyright 2019-2024 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// SteeringFunctionality definitions.
const (
	SteeringFunctionalityATSSSLL uint8 = 0
	SteeringFunctionalityMPTCP   uint8 = 1
)

// NewSteeringFunctionality creates a new SteeringFunctionality IE.
func NewSteeringFunctionality(sfunc uint8) *IE {
	return newUint8ValIE(SteeringFunctionality, sfunc&0x0f)
}

// SteeringFunctionality returns SteeringFunctionality in uint8 if the type of IE matches.
func (i *IE) SteeringFunctionality() (uint8, error) {
	switch i.Type {
	case SteeringFunctionality:
		return i.ValueAsUint8()
	case CreateMAR:
		ies, err := i.CreateMAR()
		if err != nil {
			return 0, err
		}
		for _, x := range ies {
			if x.Type == SteeringFunctionality {
				return x.SteeringFunctionality()
			}
		}
		return 0, ErrIENotFound
	case UpdateMAR:
		ies, err := i.UpdateMAR()
		if err != nil {
			return 0, err
		}
		for _, x := range ies {
			if x.Type == SteeringFunctionality {
				return x.SteeringFunctionality()
			}
		}
		return 0, ErrIENotFound
	default:
		return 0, &InvalidTypeError{Type: i.Type}
	}
}
