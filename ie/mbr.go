// Copyright 2019-2021 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import (
	"encoding/binary"
	"io"
)

const (
	MBR_PAYLOAD_SIZE   int = 10
	MBR_UL_START_INDEX int = 1
	MBR_UL_END_INDEX   int = 5
	MBR_DL_START_INDEX int = 6
	MBR_DL_END_INDEX   int = 10
)

// NewMBR creates a new MBR IE.
func NewMBR(ul, dl uint32) *IE {
	i := New(MBR, make([]byte, MBR_PAYLOAD_SIZE))
	binary.BigEndian.PutUint32(i.Payload[MBR_UL_START_INDEX:MBR_UL_END_INDEX], ul)
	binary.BigEndian.PutUint32(i.Payload[MBR_DL_START_INDEX:MBR_DL_END_INDEX], dl)
	return i
}

// MBR returns MBR in []byte if the type of IE matches.
func (i *IE) MBR() ([]byte, error) {
	if len(i.Payload) < MBR_PAYLOAD_SIZE {
		return nil, io.ErrUnexpectedEOF
	}

	switch i.Type {
	case MBR:
		return i.Payload, nil
	case CreateQER:
		ies, err := i.CreateQER()
		if err != nil {
			return nil, err
		}
		for _, x := range ies {
			if x.Type == MBR {
				return x.MBR()
			}
		}
		return nil, ErrIENotFound
	case UpdateQER:
		ies, err := i.UpdateQER()
		if err != nil {
			return nil, err
		}
		for _, x := range ies {
			if x.Type == MBR {
				return x.MBR()
			}
		}
		return nil, ErrIENotFound
	default:
		return nil, &InvalidTypeError{Type: i.Type}
	}
}

// MBRUL returns MBRUL in uint32 if the type of IE matches.
func (i *IE) MBRUL() (uint32, error) {
	v, err := i.MBR()
	if err != nil {
		return 0, err
	}
	return binary.BigEndian.Uint32(v[MBR_UL_START_INDEX:MBR_UL_END_INDEX]), nil
}

// MBRDL returns MBRDL in uint32 if the type of IE matches.
func (i *IE) MBRDL() (uint32, error) {
	v, err := i.MBR()
	if err != nil {
		return 0, err
	}
	return binary.BigEndian.Uint32(v[MBR_DL_START_INDEX:MBR_DL_END_INDEX]), nil
}
