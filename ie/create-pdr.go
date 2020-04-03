// Copyright 2019-2020 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

func NewCreatePDR(teid uint16, pdrid *IE, precedence *IE, pdi *IE, outerHeaderRemoval *IE, farID *IE) *IE {
	return newGroupedIE(CreatePDR, teid, pdrid, precedence, pdi, outerHeaderRemoval, farID)
}
