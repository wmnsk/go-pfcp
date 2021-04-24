// Copyright 2019-2021 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import (
	"encoding/binary"
	"io"
)

const (
	GBR_PAYLOAD_SIZE   int = 10
	GBR_UL_START_INDEX int = 1
	GBR_UL_END_INDEX   int = 5
	GBR_DL_START_INDEX int = 6
	GBR_DL_END_INDEX   int = 10
)

// NewGBR creates a new GBR IE.
func NewGBR(ul, dl uint32) *IE {
	i := New(GBR, make([]byte, GBR_PAYLOAD_SIZE))
	binary.BigEndian.PutUint32(i.Payload[GBR_UL_START_INDEX:GBR_UL_END_INDEX], ul)
	binary.BigEndian.PutUint32(i.Payload[GBR_DL_START_INDEX:GBR_DL_END_INDEX], dl)
	return i
}

// GBR returns GBR in []byte if the type of IE matches.
func (i *IE) GBR() ([]byte, error) {
	if len(i.Payload) < GBR_PAYLOAD_SIZE {
		return nil, io.ErrUnexpectedEOF
	}

	switch i.Type {
	case GBR:
		return i.Payload, nil
	case CreateQER:
		ies, err := i.CreateQER()
		if err != nil {
			return nil, err
		}
		for _, x := range ies {
			if x.Type == GBR {
				return x.GBR()
			}
		}
		return nil, ErrIENotFound
	case UpdateQER:
		ies, err := i.UpdateQER()
		if err != nil {
			return nil, err
		}
		for _, x := range ies {
			if x.Type == GBR {
				return x.GBR()
			}
		}
		return nil, ErrIENotFound
	default:
		return nil, &InvalidTypeError{Type: i.Type}
	}
}

// GBRUL returns GBRUL in uint32 if the type of IE matches.
func (i *IE) GBRUL() (uint32, error) {
	v, err := i.GBR()
	if err != nil {
		return 0, err
	}
	return binary.BigEndian.Uint32(v[GBR_UL_START_INDEX:GBR_UL_END_INDEX]), nil
}

// GBRDL returns GBRDL in uint32 if the type of IE matches.
func (i *IE) GBRDL() (uint32, error) {
	v, err := i.GBR()
	if err != nil {
		return 0, err
	}
	return binary.BigEndian.Uint32(v[GBR_DL_START_INDEX:GBR_DL_END_INDEX]), nil
}
