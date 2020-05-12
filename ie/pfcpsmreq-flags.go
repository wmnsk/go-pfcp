// Copyright 2019-2020 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewPFCPSMReqFlags creates a new PFCPSMReqFlags IE.
func NewPFCPSMReqFlags(flag uint8) *IE {
	return newUint8ValIE(PFCPSMReqFlags, flag)
}

// PFCPSMReqFlags returns PFCPSMReqFlags in []byte if the type of IE matches.
func (i *IE) PFCPSMReqFlags() ([]byte, error) {
	if i.Type != PFCPSMReqFlags {
		return nil, &InvalidTypeError{Type: i.Type}
	}

	return i.Payload, nil
}

// HasDROBU reports whether an IE has DROBU bit.
func (i *IE) HasDROBU() bool {
	if i.Type != PFCPSMReqFlags && i.Type != PFCPSRRspFlags {
		return false
	}
	if len(i.Payload) < 1 {
		return false
	}

	return has1stBit(i.Payload[0])
}

// HasSNDEM reports whether an IE has SNDEM bit.
func (i *IE) HasSNDEM() bool {
	if i.Type != PFCPSMReqFlags {
		return false
	}
	if len(i.Payload) < 1 {
		return false
	}

	return has2ndBit(i.Payload[0])
}

// HasQAURR reports whether an IE has QAURR bit.
func (i *IE) HasQAURR() bool {
	if i.Type != PFCPSMReqFlags {
		return false
	}
	if len(i.Payload) < 1 {
		return false
	}

	return has3rdBit(i.Payload[0])
}
