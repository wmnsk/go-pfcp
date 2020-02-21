// Copyright 2019-2020 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import (
	"encoding/binary"
	"io"
)

// NewSubsequentVolumeQuota creates a new SubsequentVolumeQuota IE.
func NewSubsequentVolumeQuota(flags uint8, total, ul, dl uint64) *IE {
	i := New(SubsequentVolumeQuota, []byte{flags})

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

// SubsequentVolumeQuota returns SubsequentVolumeQuota in []byte if the type of IE matches.
func (i *IE) SubsequentVolumeQuota() ([]byte, error) {
	if i.Type != SubsequentVolumeQuota {
		return nil, &InvalidTypeError{Type: i.Type}
	}
	if len(i.Payload) < 1 {
		return nil, io.ErrUnexpectedEOF
	}

	return i.Payload, nil
}

// SubsequentVolumeQuotaTotal returns SubsequentVolumeQuotaTotal in uint64 if the type of IE matches.
func (i *IE) SubsequentVolumeQuotaTotal() (uint64, error) {
	paylod, err := i.SubsequentVolumeQuota()
	if err != nil {
		return 0, err
	}

	if has1stBit(paylod[0]) && len(paylod) >= 8 {
		return binary.BigEndian.Uint64(paylod[1:9]), nil
	}
	return 0, nil
}

// SubsequentVolumeQuotaUplink returns SubsequentVolumeQuotaUplink in uint64 if the type of IE matches.
func (i *IE) SubsequentVolumeQuotaUplink() (uint64, error) {
	paylod, err := i.SubsequentVolumeQuota()
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

// SubsequentVolumeQuotaDownlink returns SubsequentVolumeQuotaDownlink in uint64 if the type of IE matches.
func (i *IE) SubsequentVolumeQuotaDownlink() (uint64, error) {
	paylod, err := i.SubsequentVolumeQuota()
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
