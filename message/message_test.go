// Copyright 2019-2021 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package message_test

import "net"

var (
	mac1, _ = net.ParseMAC("12:34:56:78:90:01")
	mac2, _ = net.ParseMAC("12:34:56:78:90:02")
	mac3, _ = net.ParseMAC("12:34:56:78:90:03")
	mac4, _ = net.ParseMAC("12:34:56:78:90:04")

	mp   uint8  = 0                  // Flags
	fo   uint8  = 0                  // Flags
	seid uint64 = 0x1122334455667788 // SEID
	seq  uint32 = 0x112233           // Sequence Number
	pri  uint8  = 0                  // Message Priority
)
