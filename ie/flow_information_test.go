// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie_test

import (
	"errors"
	"io"
	"testing"

	"github.com/wmnsk/go-pfcp/ie"
)

func TestFlowDescription(t *testing.T) {
	cases := []struct {
		description string
		payload     []byte
		want        string
		err         error
	}{
		{
			description: "Valid",
			payload:     []byte{0x01, 0x00, 0x03, 'f', 'o', 'o'},
			want:        "foo",
		},
		{
			description: "EmptyPayload",
			payload:     []byte{},
			err:         io.ErrUnexpectedEOF,
		},
		{
			description: "NoLength",
			payload:     []byte{0x01},
			err:         io.ErrUnexpectedEOF,
		},
		{
			description: "LengthCountsHeader",
			payload:     []byte{0x01, 0x00, 0x05, 'f', 'o'},
			err:         io.ErrUnexpectedEOF,
		},
	}

	for _, c := range cases {
		t.Run(c.description, func(t *testing.T) {
			got, err := ie.New(ie.FlowInformation, c.payload).FlowDescription()
			if !errors.Is(err, c.err) {
				t.Fatalf("got error %v, want %v", err, c.err)
			}
			if got != c.want {
				t.Errorf("got %q, want %q", got, c.want)
			}
		})
	}
}
