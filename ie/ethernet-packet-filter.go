// Copyright 2019-2020 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewEthernetPacketFilter creates a new EthernetPacketFilter IE.
func NewEthernetPacketFilter(fid, fprop, mac, etype, ctag, stag, sdfFilter *IE) *IE {
	return newGroupedIE(EthernetPacketFilter, 0, fid, fprop, mac, etype, ctag, stag, sdfFilter)
}

// EthernetPacketFilter returns the IEs above EthernetPacketFilter if the type of IE matches.
func (i *IE) EthernetPacketFilter() ([]*IE, error) {
	if i.Type != EthernetPacketFilter {
		return nil, &InvalidTypeError{Type: i.Type}
	}

	return ParseMultiIEs(i.Payload)
}
