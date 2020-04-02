// Copyright 2019-2020 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

func NewCreateFAR(teid uint16, farID *IE, applyAction *IE, forwardingParameter *IE) *IE {
	return newGroupedIE(CreateFAR, teid, farID, applyAction, forwardingParameter)
}
