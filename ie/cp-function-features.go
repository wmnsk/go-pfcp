// Copyright 2019-2024 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import "io"

// NewCPFunctionFeatures creates a new CPFunctionFeatures IE.
func NewCPFunctionFeatures(features ...uint8) *IE {
	return New(CPFunctionFeatures, features)
}

// CPFunctionFeatures returns CPFunctionFeatures in []byte if the type of IE matches.
func (i *IE) CPFunctionFeatures() ([]byte, error) {
	if i.Type != CPFunctionFeatures {
		return nil, &InvalidTypeError{Type: i.Type}
	}
	if len(i.Payload) < 1 {
		return nil, io.ErrUnexpectedEOF
	}
	return i.Payload, nil
}

// HasLOAD reports whether an IE has LOAD bit.
func (i *IE) HasLOAD() bool {
	if i.Type != CPFunctionFeatures {
		return false
	}
	if len(i.Payload) < 1 {
		return false
	}

	return has1stBit(i.Payload[0])
}

// HasOVRL reports whether an IE has OVRL bit.
func (i *IE) HasOVRL() bool {
	if i.Type != CPFunctionFeatures {
		return false
	}
	if len(i.Payload) < 1 {
		return false
	}

	return has2ndBit(i.Payload[0])
}

// HasARDR reports whether an IE has ARDR bit.
func (i *IE) HasARDR() bool {
	if i.Type != CPFunctionFeatures {
		return false
	}
	if len(i.Payload) < 1 {
		return false
	}

	return has7thBit(i.Payload[0])
}

// HasUIAUR reports whether an IE has UIAUR bit.
func (i *IE) HasUIAUR() bool {
	if i.Type != CPFunctionFeatures {
		return false
	}
	if len(i.Payload) < 1 {
		return false
	}

	return has8thBit(i.Payload[0])
}

// HasPSUCC reports whether an IE has PSUCC bit.
func (i *IE) HasPSUCC() bool {
	if i.Type != CPFunctionFeatures {
		return false
	}
	if len(i.Payload) < 2 {
		return false
	}

	return has1stBit(i.Payload[1])
}

// HasRPGUR reports whether an IE has RPGUR bit.
func (i *IE) HasRPGUR() bool {
	if i.Type != CPFunctionFeatures {
		return false
	}
	if len(i.Payload) < 2 {
		return false
	}

	return has2ndBit(i.Payload[1])
}
