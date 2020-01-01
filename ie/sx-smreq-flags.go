// Copyright 2019-2020 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewSxSMReqFlags creates a new SxSMReqFlags IE.
func NewSxSMReqFlags(flag uint8) *IE {
	return newUint8ValIE(SxSMReqFlags, flag)
}

// SxSMReqFlags returns SxSMReqFlags in []byte if the type of IE matches.
func (i *IE) SxSMReqFlags() ([]byte, error) {
	if i.Type != SxSMReqFlags {
		return nil, &InvalidTypeError{Type: i.Type}
	}

	return i.Payload, nil
}

// HasDROBU reports whether up function features has DROBU bit.
func (i *IE) HasDROBU() bool {
	if i.Type != SxSMReqFlags && i.Type != SxSRRspFlags {
		return false
	}
	if len(i.Payload) < 1 {
		return false
	}

	return has1stBit(i.Payload[0])
}

// HasSNDEM reports whether up function features has SNDEM bit.
func (i *IE) HasSNDEM() bool {
	if i.Type != SxSMReqFlags {
		return false
	}
	if len(i.Payload) < 1 {
		return false
	}

	return has2ndBit(i.Payload[0])
}

// HasQAURR reports whether up function features has QAURR bit.
func (i *IE) HasQAURR() bool {
	if i.Type != SxSMReqFlags {
		return false
	}
	if len(i.Payload) < 1 {
		return false
	}

	return has3rdBit(i.Payload[0])
}
