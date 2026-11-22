// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie_test

import (
	"errors"
	"io"
	"net"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/wmnsk/go-pfcp/ie"
)

func TestParseSourceIPAddressFields(t *testing.T) {
	cases := []struct {
		description string
		serialized  []byte
		structured  *ie.SourceIPAddressFields
		err         error
	}{
		{
			description: "IPv4AndMPL",
			serialized:  []byte{0x06, 127, 0, 0, 1, 24},
			structured: &ie.SourceIPAddressFields{
				Flags: 0x06, IPv4Address: net.IP{127, 0, 0, 1}, MaskPrefixLength: 24,
			},
		},
		{
			description: "MPLMissing",
			serialized:  []byte{0x06, 127, 0, 0, 1},
			err:         io.ErrUnexpectedEOF,
		},
	}

	for _, c := range cases {
		t.Run(c.description, func(t *testing.T) {
			got, err := ie.ParseSourceIPAddressFields(c.serialized)
			if !errors.Is(err, c.err) {
				t.Fatalf("got error %v, want %v", err, c.err)
			}
			if diff := cmp.Diff(c.structured, got); diff != "" {
				t.Error(diff)
			}
		})
	}
}

func TestSourceIPAddressMarshalToShortBuffer(t *testing.T) {
	f := ie.NewSourceIPAddressFields(net.ParseIP("127.0.0.1"), nil, 24)
	if err := f.MarshalTo(make([]byte, 5)); !errors.Is(err, io.ErrUnexpectedEOF) {
		t.Errorf("got error %v, want %v", err, io.ErrUnexpectedEOF)
	}
}
