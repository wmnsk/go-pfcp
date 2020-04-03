// Copyright 2019-2020 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

func NewPDI(teid uint16, sourceInterface *IE, fteid *IE, ueIpAddress *IE) *IE {
	return newGroupedIE(PDI, teid, sourceInterface, fteid, ueIpAddress)
}
