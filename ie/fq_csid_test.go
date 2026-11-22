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

func TestCSIDs(t *testing.T) {
	cases := []struct {
		description string
		ie          *ie.IE
		want        []uint16
		err         error
	}{
		{
			description: "IPv4",
			ie:          ie.NewFQCSID("127.0.0.1", 1, 2),
			want:        []uint16{1, 2},
		},
		{
			description: "IPv6",
			ie:          ie.NewFQCSID("2001::1", 1),
			want:        []uint16{1},
		},
		{
			description: "NoCSID",
			ie:          ie.NewFQCSID("127.0.0.1"),
		},
		{
			description: "ShortPayload",
			ie:          ie.New(ie.FQCSID, []byte{0x02}),
			err:         io.ErrUnexpectedEOF,
		},
	}

	for _, c := range cases {
		t.Run(c.description, func(t *testing.T) {
			got, err := c.ie.CSIDs()
			if !errors.Is(err, c.err) {
				t.Fatalf("got error %v, want %v", err, c.err)
			}
			if diff := cmp.Diff(c.want, got); diff != "" {
				t.Error(diff)
			}
		})
	}
}
