// Copyright 2019-2020 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import (
	"encoding/binary"
)

// NewVolumeQuota creates a new VolumeQuota IE.
func NewVolumeQuota(tovol, ulvol, dlvol bool, total, ul, dl uint64) *IE {
	b := []byte{0x00}
	offset := 1
	if tovol {
		b[0] |= 0x01
		b = append(b, make([]byte, 8)...)
		binary.BigEndian.PutUint64(b[offset:offset+8], total)
		offset += 8
	}
	if ulvol {
		b[0] |= 0x02
		b = append(b, make([]byte, 8)...)
		binary.BigEndian.PutUint64(b[offset:offset+8], ul)
		offset += 8
	}
	if dlvol {
		b[0] |= 0x04
		b = append(b, make([]byte, 8)...)
		binary.BigEndian.PutUint64(b[offset:offset+8], dl)
	}

	return New(VolumeQuota, b)
}

// VolumeQuota returns VolumeQuota in uint8 if the type of IE matches.
func (i *IE) VolumeQuota() (uint8, error) {
	if i.Type != VolumeQuota {
		return 0, &InvalidTypeError{Type: i.Type}
	}

	return i.Payload[0], nil
}

// HasDLVOL reports whether VolumeQuota IE has DLVOL bit.
func (i *IE) HasDLVOL() bool {
	v, err := i.VolumeQuota()
	if err != nil {
		return false
	}

	return has3rdBit(v)
}

// HasULVOL reports whether VolumeQuota IE has ULVOL bit.
func (i *IE) HasULVOL() bool {
	v, err := i.VolumeQuota()
	if err != nil {
		return false
	}

	return has2ndBit(v)
}

// HasTOVOL reports whether VolumeQuota IE has TOVOL bit.
func (i *IE) HasTOVOL() bool {
	v, err := i.VolumeQuota()
	if err != nil {
		return false
	}

	return has1stBit(v)
}
