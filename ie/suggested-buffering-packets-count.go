// Copyright 2019-2020 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import (
	"io"
)

// NewSuggestedBufferingPacketsCount creates a new SuggestedBufferingPacketsCount IE.
func NewSuggestedBufferingPacketsCount(count uint8) *IE {
	return newUint8ValIE(SuggestedBufferingPacketsCount, count)
}

// SuggestedBufferingPacketsCount returns SuggestedBufferingPacketsCount in uint8 if the type of IE matches.
func (i *IE) SuggestedBufferingPacketsCount() (uint8, error) {
	if i.Type != SuggestedBufferingPacketsCount {
		return 0, &InvalidTypeError{Type: i.Type}
	}
	if len(i.Payload) < 1 {
		return 0, io.ErrUnexpectedEOF
	}

	return i.Payload[0], nil
}
