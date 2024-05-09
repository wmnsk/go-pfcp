// Copyright 2019-2024 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewDLDataPacketsSize creates a new DLDataPacketsSize IE.
func NewDLDataPacketsSize(size uint16) *IE {
	return newUint16ValIE(DLDataPacketsSize, size)
}

// DLDataPacketsSize returns DLDataPacketsSize in uint16 if the type of IE matches.
func (i *IE) DLDataPacketsSize() (uint16, error) {
	switch i.Type {
	case DLDataPacketsSize:
		return i.ValueAsUint16()
	case DownlinkDataReport:
		ies, err := i.DownlinkDataReport()
		if err != nil {
			return 0, err
		}
		for _, x := range ies {
			if x.Type == DLDataPacketsSize {
				return x.DLDataPacketsSize()
			}
		}
		return 0, ErrIENotFound
	default:
		return 0, &InvalidTypeError{Type: i.Type}
	}
}
