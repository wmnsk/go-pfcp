// Copyright 2019-2021 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import "io"

// NewDataStatus creates a new DataStatus IE.
func NewDataStatus(flag uint8) *IE {
	return newUint8ValIE(DataStatus, flag)
}

// DataStatus returns DataStatus in uint8 if the type of IE matches.
func (i *IE) DataStatus() (uint8, error) {
	if i.Type != DataStatus {
		return 0, &InvalidTypeError{Type: i.Type}
	}
	if len(i.Payload) < 1 {
		return 0, io.ErrUnexpectedEOF
	}

	return i.Payload[0], nil
}
