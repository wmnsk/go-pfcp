// Copyright 2019-2020 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import (
	"encoding/binary"
	"io"
)

// NewSubsequentVolumeThreshold creates a new SubsequentVolumeThreshold IE.
func NewSubsequentVolumeThreshold(flags uint8, total, ul, dl uint64) *IE {
	i := New(SubsequentVolumeThreshold, []byte{flags})

	offset := 1
	if has1stBit(flags) {
		i.Payload = append(i.Payload, make([]byte, 8)...)
		binary.BigEndian.PutUint64(i.Payload[offset:offset+8], total)
		offset += 8
	}
	if has2ndBit(flags) {
		i.Payload = append(i.Payload, make([]byte, 8)...)
		binary.BigEndian.PutUint64(i.Payload[offset:offset+8], ul)
		offset += 8
	}
	if has3rdBit(flags) {
		i.Payload = append(i.Payload, make([]byte, 8)...)
		binary.BigEndian.PutUint64(i.Payload[offset:offset+8], dl)
	}

	i.SetLength()
	return i
}

// SubsequentVolumeThreshold returns SubsequentVolumeThreshold in []byte if the type of IE matches.
func (i *IE) SubsequentVolumeThreshold() ([]byte, error) {
	if i.Type != SubsequentVolumeThreshold {
		return nil, &InvalidTypeError{Type: i.Type}
	}
	if len(i.Payload) < 1 {
		return nil, io.ErrUnexpectedEOF
	}

	return i.Payload, nil
}

// SubsequentVolumeThresholdTotal returns SubsequentVolumeThresholdTotal in uint64 if the type of IE matches.
func (i *IE) SubsequentVolumeThresholdTotal() (uint64, error) {
	paylod, err := i.SubsequentVolumeThreshold()
	if err != nil {
		return 0, err
	}

	if has1stBit(paylod[0]) && len(paylod) >= 8 {
		return binary.BigEndian.Uint64(paylod[1:9]), nil
	}
	return 0, nil
}

// SubsequentVolumeThresholdUplink returns SubsequentVolumeThresholdUplink in uint64 if the type of IE matches.
func (i *IE) SubsequentVolumeThresholdUplink() (uint64, error) {
	paylod, err := i.SubsequentVolumeThreshold()
	if err != nil {
		return 0, err
	}

	if !has2ndBit(paylod[0]) {
		return 0, nil
	}

	offset := 1
	if has1stBit(i.Payload[0]) {
		offset += 8
	}

	if len(paylod) < offset+8 {
		return 0, nil
	}
	return binary.BigEndian.Uint64(paylod[offset : offset+8]), nil
}

// SubsequentVolumeThresholdDownlink returns SubsequentVolumeThresholdDownlink in uint64 if the type of IE matches.
func (i *IE) SubsequentVolumeThresholdDownlink() (uint64, error) {
	paylod, err := i.SubsequentVolumeThreshold()
	if err != nil {
		return 0, err
	}

	if !has2ndBit(paylod[0]) {
		return 0, nil
	}

	offset := 1
	if has1stBit(i.Payload[0]) {
		offset += 8
	}
	if has2ndBit(i.Payload[0]) {
		offset += 8
	}

	if len(paylod) < offset+8 {
		return 0, nil
	}
	return binary.BigEndian.Uint64(paylod[offset : offset+8]), nil
}
