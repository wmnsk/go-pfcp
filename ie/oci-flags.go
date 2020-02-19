// Copyright 2019-2020 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import (
	"io"
)

// NewOCIFlags creates a new OCIFlags IE.
func NewOCIFlags(flags uint8) *IE {
	return newUint8ValIE(OCIFlags, flags)
}

// OCIFlags returns OCIFlags in uint8 if the type of IE matches.
func (i *IE) OCIFlags() (uint8, error) {
	if i.Type != OCIFlags {
		return 0, &InvalidTypeError{Type: i.Type}
	}
	if len(i.Payload) < 1 {
		return 0, io.ErrUnexpectedEOF
	}

	return i.Payload[0], nil
}
