// Copyright 2019-2020 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewPacketReplicationAndDetectionCarryOnInformation creates a new PacketReplicationAndDetectionCarryOnInformation IE.
func NewPacketReplicationAndDetectionCarryOnInformation(flag uint8) *IE {
	return newUint8ValIE(PacketReplicationAndDetectionCarryOnInformation, flag)
}

// PacketReplicationAndDetectionCarryOnInformation returns PacketReplicationAndDetectionCarryOnInformation in []byte if the type of IE matches.
func (i *IE) PacketReplicationAndDetectionCarryOnInformation() ([]byte, error) {
	if i.Type != PacketReplicationAndDetectionCarryOnInformation {
		return nil, &InvalidTypeError{Type: i.Type}
	}

	return i.Payload, nil
}

// HasPRIUEAI reports whether an IE has PRIUEAI bit.
func (i *IE) HasPRIUEAI() bool {
	if i.Type != PacketReplicationAndDetectionCarryOnInformation && i.Type != SxSRRspFlags {
		return false
	}
	if len(i.Payload) < 1 {
		return false
	}

	return has1stBit(i.Payload[0])
}

// HasPRINT19I reports whether an IE has PRINT19I bit.
func (i *IE) HasPRINT19I() bool {
	if i.Type != PacketReplicationAndDetectionCarryOnInformation {
		return false
	}
	if len(i.Payload) < 1 {
		return false
	}

	return has2ndBit(i.Payload[0])
}

// HasPRIN6I reports whether an IE has PRIN6I bit.
func (i *IE) HasPRIN6I() bool {
	if i.Type != PacketReplicationAndDetectionCarryOnInformation {
		return false
	}
	if len(i.Payload) < 1 {
		return false
	}

	return has3rdBit(i.Payload[0])
}

// HasDCARONI reports whether an IE has DCARONI bit.
func (i *IE) HasDCARONI() bool {
	if i.Type != PacketReplicationAndDetectionCarryOnInformation {
		return false
	}
	if len(i.Payload) < 1 {
		return false
	}

	return has4thBit(i.Payload[0])
}
