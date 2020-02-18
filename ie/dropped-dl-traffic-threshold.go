// Copyright 2019-2020 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import (
	"encoding/binary"
)

// NewDroppedDLTrafficThreshold creates a new DroppedDLTrafficThreshold IE.
func NewDroppedDLTrafficThreshold(dlpa, dlby bool, packets, bytes uint64) *IE {
	if dlpa {
		if dlby { // has both
			i := New(DroppedDLTrafficThreshold, make([]byte, 17))
			i.Payload[0] = 0x03
			binary.BigEndian.PutUint64(i.Payload[1:9], packets)
			binary.BigEndian.PutUint64(i.Payload[9:17], bytes)
			return i
		}

		// has DLPA only
		i := New(DroppedDLTrafficThreshold, make([]byte, 9))
		i.Payload[0] = 0x01
		binary.BigEndian.PutUint64(i.Payload[1:9], packets)
		return i
	}

	if dlby { // has DLBY only
		i := New(DroppedDLTrafficThreshold, make([]byte, 9))
		i.Payload[0] = 0x02
		binary.BigEndian.PutUint64(i.Payload[1:9], bytes)
		return i
	}

	return New(DroppedDLTrafficThreshold, []byte{0x00})
}

// DroppedDLTrafficThreshold returns DroppedDLTrafficThreshold in uint8 if the type of IE matches.
func (i *IE) DroppedDLTrafficThreshold() (uint8, error) {
	if i.Type != DroppedDLTrafficThreshold {
		return 0, &InvalidTypeError{Type: i.Type}
	}

	return i.Payload[0], nil
}

// HasDLBY reports whether DroppedDLTrafficThreshold IE has DLBY bit.
func (i *IE) HasDLBY() bool {
	v, err := i.DroppedDLTrafficThreshold()
	if err != nil {
		return false
	}

	return has2ndBit(v)
}

// HasDLPA reports whether DroppedDLTrafficThreshold IE has DLPA bit.
func (i *IE) HasDLPA() bool {
	v, err := i.DroppedDLTrafficThreshold()
	if err != nil {
		return false
	}

	return has1stBit(v)
}
