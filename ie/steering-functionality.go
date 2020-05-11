// Copyright 2019-2020 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import "io"

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
	if i.Type != SteeringFunctionality {
		return 0, &InvalidTypeError{Type: i.Type}
	}
	if len(i.Payload) == 0 {
		return 0, io.ErrUnexpectedEOF
	}

	return i.Payload[0], nil
}
