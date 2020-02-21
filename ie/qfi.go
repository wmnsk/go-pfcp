// Copyright 2019-2020 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewQFI creates a new QFI IE.
func NewQFI(qfi uint8) *IE {
	return newUint8ValIE(QFI, qfi)
}
