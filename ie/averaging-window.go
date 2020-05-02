// Copyright 2019-2020 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import (
	"encoding/binary"
	"io"
)

// NewAveragingWindow creates a new AveragingWindow IE.
func NewAveragingWindow(window uint32) *IE {
	return newUint32ValIE(AveragingWindow, window)
}

// AveragingWindow returns AveragingWindow in uint32 if the type of IE matches.
func (i *IE) AveragingWindow() (uint32, error) {
	if i.Type != AveragingWindow {
		return 0, &InvalidTypeError{Type: i.Type}
	}
	if len(i.Payload) < 4 {
		return 0, io.ErrUnexpectedEOF
	}

	return binary.BigEndian.Uint32(i.Payload[0:4]), nil
}
