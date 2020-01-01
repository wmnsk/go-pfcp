// Copyright 2019-2020 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// Interface definitions.
const (
	DstInterfaceAccess       uint8 = 0
	DstInterfaceCore         uint8 = 1
	DstInterfaceSGiLANN6LAN  uint8 = 2
	DstInterfaceCPFunction   uint8 = 3
	DstInterfaceLIFunction   uint8 = 4
	DstInterface5GVNInternal uint8 = 5
)

// NewDestinationInterface creates a new DestinationInterface IE.
func NewDestinationInterface(intf uint8) *IE {
	return newUint8ValIE(DestinationInterface, intf)
}

// DestinationInterface returns DestinationInterface in uint8 if the type of IE matches.
func (i *IE) DestinationInterface() (uint8, error) {
	if i.Type != DestinationInterface {
		return 0, &InvalidTypeError{Type: i.Type}
	}

	return i.Payload[0], nil
}
