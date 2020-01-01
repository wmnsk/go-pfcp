// Copyright 2019-2020 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewSxSRRspFlags creates a new SxSRRspFlags IE.
func NewSxSRRspFlags(flag uint8) *IE {
	return newUint8ValIE(SxSRRspFlags, flag)
}

// SxSRRspFlags returns SxSRRspFlags in []byte if the type of IE matches.
func (i *IE) SxSRRspFlags() ([]byte, error) {
	if i.Type != SxSRRspFlags {
		return nil, &InvalidTypeError{Type: i.Type}
	}

	return i.Payload, nil
}
