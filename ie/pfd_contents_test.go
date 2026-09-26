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

func TestParsePFDContentsFields(t *testing.T) {
	all := ie.NewPFDContentsFields(
		"permit out ip from any to any", "http://example.com", "example.com", "custom", "dnp",
		[]string{"afd1", "afd2"}, []string{"aurl1"}, []string{"adnp1"},
	)
	allBytes, err := all.Marshal()
	if err != nil {
		t.Fatal(err)
	}

	cases := []struct {
		description string
		serialized  []byte
		structured  *ie.PFDContentsFields
		err         error
	}{
		{
			description: "AllFields",
			serialized:  allBytes,
			structured:  all,
		},
		{
			description: "FDAndURL",
			serialized:  []byte{0x03, 0x00, 0x00, 0x01, 'a', 0x00, 0x02, 'b', 'c'},
			structured: &ie.PFDContentsFields{
				Flags: 0x03, FDLength: 1, FlowDescription: "a", URLLength: 2, URL: "bc",
			},
		},
		{
			description: "FDLengthOverrunsPayload",
			serialized:  []byte{0x01, 0x00, 0xff, 0xff, 0x00, 0x00},
			err:         io.ErrUnexpectedEOF,
		},
		{
			description: "FDLengthMissing",
			serialized:  []byte{0x01, 0x00, 0x00},
			err:         io.ErrUnexpectedEOF,
		},
		{
			description: "AFDLengthOverrunsPayload",
			serialized:  []byte{0x10, 0x00, 0x1a, 0x46},
			err:         io.ErrUnexpectedEOF,
		},
	}

	for _, c := range cases {
		t.Run(c.description, func(t *testing.T) {
			got, err := ie.ParsePFDContentsFields(c.serialized)
			if !errors.Is(err, c.err) {
				t.Fatalf("got error %v, want %v", err, c.err)
			}
			if diff := cmp.Diff(c.structured, got); diff != "" {
				t.Error(diff)
			}
		})
	}
}
