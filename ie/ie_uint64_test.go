// Copyright 2019-2022 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/wmnsk/go-pfcp/ie"
)

func TestUint64IEs(t *testing.T) {
	cases := []struct {
		description string
		structured  *ie.IE
		decoded     uint64
		decoderFunc func(*ie.IE) (uint64, error)
	}{
		{
			description: "Multiplier/ValueDigits",
			structured:  ie.NewMultiplier(0xffffffffffffffff, 0x11223344),
			decoded:     0xffffffffffffffff,
			decoderFunc: func(i *ie.IE) (uint64, error) { return i.ValueDigits() },
		},
	}

	for _, c := range cases {
		t.Run(c.description, func(t *testing.T) {
			got, err := c.decoderFunc(c.structured)
			if err != nil {
				t.Fatal(err)
			}

			if diff := cmp.Diff(got, c.decoded); diff != "" {
				t.Error(diff)
			}
		})
	}
}
