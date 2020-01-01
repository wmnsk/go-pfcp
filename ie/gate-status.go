// Copyright 2019-2020 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// GateStatus definitions.
const (
	GateStatusOpen   uint8 = 0
	GateStatusClosed uint8 = 1
)

// NewGateStatus creates a new GateStatus IE.
func NewGateStatus(ul, dl uint8) *IE {
	return newUint8ValIE(GateStatus, (ul<<2)|dl)
}

// GateStatus returns GateStatus in uint8 if the type of IE matches.
func (i *IE) GateStatus() (uint8, error) {
	if i.Type != GateStatus {
		return 0, &InvalidTypeError{Type: i.Type}
	}

	return i.Payload[0], nil
}

// GateStatusUL returns GateStatusUL in uint8 if the type of IE matches.
func (i *IE) GateStatusUL() (uint8, error) {
	if i.Type != GateStatus {
		return 0, &InvalidTypeError{Type: i.Type}
	}

	return (i.Payload[0] >> 2) & 0xff, nil
}

// GateStatusDL returns GateStatusDL in uint8 if the type of IE matches.
func (i *IE) GateStatusDL() (uint8, error) {
	if i.Type != GateStatus {
		return 0, &InvalidTypeError{Type: i.Type}
	}

	return i.Payload[0] & 0xff, nil
}
