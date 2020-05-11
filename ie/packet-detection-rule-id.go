// Copyright 2019-2020 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import (
	"encoding/binary"
)

// NewPacketDetectionRuleID creates a new PacketDetectionRuleID IE.
func NewPacketDetectionRuleID(id uint16) *IE {
	return newUint16ValIE(PacketDetectionRuleID, id)
}

// PacketDetectionRuleID returns PacketDetectionRuleID in uint16 if the type of IE matches.
func (i *IE) PacketDetectionRuleID() (uint16, error) {
	switch i.Type {
	case PacketDetectionRuleID:
		return binary.BigEndian.Uint16(i.Payload[0:2]), nil
	case ApplicationDetectionInformation:
		ies, err := i.ApplicationDetectionInformation()
		if err != nil {
			return 0, err
		}
		for _, x := range ies {
			if x.Type == PacketDetectionRuleID {
				return x.PacketDetectionRuleID()
			}
		}
		return 0, ErrIENotFound
	default:
		return 0, ErrInvalidType
	}
}
