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

func TestQFIInDownlinkDataServiceInformation(t *testing.T) {
	cases := []struct {
		description string
		ie          *ie.IE
		want        uint8
		err         error
	}{
		{
			description: "QFIOnly",
			ie:          ie.NewDownlinkDataServiceInformation(false, true, 0, 9),
			want:        9,
		},
		{
			description: "PPIAndQFI",
			ie:          ie.NewDownlinkDataServiceInformation(true, true, 3, 9),
			want:        9,
		},
		{
			description: "EmptyPayload",
			ie:          ie.New(ie.DownlinkDataServiceInformation, []byte{}),
			err:         io.ErrUnexpectedEOF,
		},
		{
			description: "QFIMissing",
			ie:          ie.New(ie.DownlinkDataServiceInformation, []byte{0x03, 0x03}),
			err:         io.ErrUnexpectedEOF,
		},
	}

	for _, c := range cases {
		t.Run(c.description, func(t *testing.T) {
			got, err := c.ie.QFI()
			if !errors.Is(err, c.err) {
				t.Fatalf("got error %v, want %v", err, c.err)
			}
			if got != c.want {
				t.Errorf("got %d, want %d", got, c.want)
			}
		})
	}
}
