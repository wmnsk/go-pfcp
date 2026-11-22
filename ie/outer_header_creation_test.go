// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie_test

import (
	"errors"
	"io"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/wmnsk/go-pfcp/ie"
)

func TestParseOuterHeaderCreationFields(t *testing.T) {
	marshal := func(f *ie.OuterHeaderCreationFields) []byte {
		b, err := f.Marshal()
		if err != nil {
			t.Fatal(err)
		}
		return b
	}
	ctag := ie.NewOuterHeaderCreationFields(0x4000, 0, "", "", 0, 0x123456, 0)
	both := ie.NewOuterHeaderCreationFields(0xc000, 0, "", "", 0, 0x123456, 0x654321)

	cases := []struct {
		description string
		serialized  []byte
		structured  *ie.OuterHeaderCreationFields
		err         error
	}{
		{
			description: "CTag",
			serialized:  marshal(ctag),
			structured:  ctag,
		},
		{
			description: "CTagAndSTag",
			serialized:  marshal(both),
			structured:  both,
		},
		{
			description: "CTagTooShort",
			serialized:  []byte{0x40, 0x00, 0x12, 0x34},
			err:         io.ErrUnexpectedEOF,
		},
	}

	for _, c := range cases {
		t.Run(c.description, func(t *testing.T) {
			got, err := ie.ParseOuterHeaderCreationFields(c.serialized)
			if !errors.Is(err, c.err) {
				t.Fatalf("got error %v, want %v", err, c.err)
			}
			if diff := cmp.Diff(c.structured, got); diff != "" {
				t.Error(diff)
			}
		})
	}
}
