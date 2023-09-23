// Copyright 2019-2022 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import (
	"io"
)

// RAT Type definitions.
const (
	RATTypeUTRAN         uint8 = 1
	RATTypeGERAN         uint8 = 2
	RATTypeWLAN          uint8 = 3
	RATTypeGAN           uint8 = 4
	RATTypeHSPAEvolution uint8 = 5
	RATTypeWBEUTRAN      uint8 = 6
	RATTypeVirtual       uint8 = 7
	RATTypeEUTRANNBIoT   uint8 = 8
	RATTypeLTEM          uint8 = 9
	RATTypeNR            uint8 = 10
)

// NewRATType creates a new RATType IE.
func NewRATType(typ uint8) *IE {
	return newUint8ValIE(RATType, typ)
}

// RATType returns RATType in uint8 if the type of IE matches.
func (i *IE) RATType() (uint8, error) {
	if i.Type != RATType {
		return 0, &InvalidTypeError{Type: i.Type}
	}
	if len(i.Payload) < 1 {
		return 0, io.ErrUnexpectedEOF
	}
	return i.Payload[0], nil
}
