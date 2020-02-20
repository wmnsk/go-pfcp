// Copyright 2019-2020 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import "io"

// NewOuterHeaderRemoval creates a new OuterHeaderRemoval IE.
func NewOuterHeaderRemoval(desc, ext uint8) *IE {
	return newUint16ValIE(OuterHeaderRemoval, uint16(desc)<<8|uint16(ext))
}

// OuterHeaderRemoval returns OuterHeaderRemoval in []byte if the type of IE matches.
func (i *IE) OuterHeaderRemoval() ([]byte, error) {
	if i.Type != OuterHeaderRemoval {
		return nil, &InvalidTypeError{Type: i.Type}
	}

	return i.Payload, nil
}

// OuterHeaderRemovalDescription returns OuterHeaderRemovalDescription in uint8 if the type of IE matches.
func (i *IE) OuterHeaderRemovalDescription() (uint8, error) {
	if i.Type != OuterHeaderRemoval {
		return 0, &InvalidTypeError{Type: i.Type}
	}
	if len(i.Payload) < 1 {
		return 0, io.ErrUnexpectedEOF
	}

	return i.Payload[0], nil
}

// GTPUExternsionHeaderDeletion returns GTPUExternsionHeaderDeletion in uint8 if the type of IE matches.
func (i *IE) GTPUExternsionHeaderDeletion() (uint8, error) {
	if i.Type != OuterHeaderRemoval {
		return 0, &InvalidTypeError{Type: i.Type}
	}
	if len(i.Payload) < 2 {
		return 0, io.ErrUnexpectedEOF
	}

	return i.Payload[1], nil
}
