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

func TestSDFFilter(t *testing.T) {
	cases := []struct {
		description string
		payload     []byte
		want        *ie.SDFFilterFields
		err         error
	}{
		{
			description: "FlowDescription",
			payload:     []byte{0x01, 0x00, 0x00, 0x03, 'f', 'o', 'o'},
			want: &ie.SDFFilterFields{
				Flags:           0x01,
				FDLength:        3,
				FlowDescription: "foo",
			},
		},
		{
			description: "FDLengthOverrunsPayload",
			payload:     []byte{0x01, 0x00, 0xff, 0xfe, 'a', 'b'},
			err:         io.ErrUnexpectedEOF,
		},
	}

	for _, c := range cases {
		t.Run(c.description, func(t *testing.T) {
			got, err := ie.New(ie.SDFFilter, c.payload).SDFFilter()
			if !errors.Is(err, c.err) {
				t.Fatalf("got error %v, want %v", err, c.err)
			}
			if diff := cmp.Diff(got, c.want); diff != "" {
				t.Error(diff)
			}
		})
	}
}
