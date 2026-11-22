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

func TestParseRemoteGTPUPeerFields(t *testing.T) {
	cases := []struct {
		description string
		serialized  []byte
		structured  *ie.RemoteGTPUPeerFields
		err         error
	}{
		{
			description: "DI",
			serialized:  []byte{0x04, 0x00, 0x01, 0x01},
			structured:  &ie.RemoteGTPUPeerFields{Flags: 0x04, DILength: 1, DestinationInterface: 1},
		},
		{
			description: "DILengthZero",
			serialized:  []byte{0x04, 0x00, 0x00},
			err:         io.ErrUnexpectedEOF,
		},
		{
			description: "DIMissing",
			serialized:  []byte{0x04, 0x00, 0x01},
			err:         io.ErrUnexpectedEOF,
		},
	}

	for _, c := range cases {
		t.Run(c.description, func(t *testing.T) {
			got, err := ie.ParseRemoteGTPUPeerFields(c.serialized)
			if !errors.Is(err, c.err) {
				t.Fatalf("got error %v, want %v", err, c.err)
			}
			if diff := cmp.Diff(c.structured, got); diff != "" {
				t.Error(diff)
			}
		})
	}
}
