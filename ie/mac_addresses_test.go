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

func TestParseMACAddressesDetectedAndRemovedFields(t *testing.T) {
	mac1, _ := net.ParseMAC("12:34:56:78:90:01")
	mac2, _ := net.ParseMAC("12:34:56:78:90:02")
	ctag := ie.NewCTAG(0x07, 1, 1, 100)
	stag := ie.NewSTAG(0x07, 1, 1, 200)
	noTag := ie.New(ie.CTAG, nil)

	cases := []struct {
		description string
		ctag, stag  *ie.IE
		macs        []net.HardwareAddr
	}{
		{"TwoMACsWithTags", ctag, stag, []net.HardwareAddr{mac1, mac2}},
		{"OneMACWithoutTags", noTag, noTag, []net.HardwareAddr{mac1}},
		{"NoMAC", ctag, noTag, nil},
	}

	for _, c := range cases {
		t.Run("Detected/"+c.description, func(t *testing.T) {
			want := ie.NewMACAddressesDetectedFields(c.ctag, c.stag, c.macs...)
			b, err := want.Marshal()
			if err != nil {
				t.Fatal(err)
			}
			got, err := ie.ParseMACAddressesDetectedFields(b)
			if err != nil {
				t.Fatal(err)
			}
			if diff := cmp.Diff(want, got); diff != "" {
				t.Error(diff)
			}
		})
		t.Run("Removed/"+c.description, func(t *testing.T) {
			want := ie.NewMACAddressesRemovedFields(c.ctag, c.stag, c.macs...)
			b, err := want.Marshal()
			if err != nil {
				t.Fatal(err)
			}
			got, err := ie.ParseMACAddressesRemovedFields(b)
			if err != nil {
				t.Fatal(err)
			}
			if diff := cmp.Diff(want, got); diff != "" {
				t.Error(diff)
			}
		})
	}
}

func TestParseMACAddressesDetectedFieldsMalformed(t *testing.T) {
	cases := []struct {
		description string
		serialized  []byte
		structured  *ie.MACAddressesDetectedFields
		err         error
	}{
		{
			description: "WithoutTagFields",
			serialized:  []byte{0x01, 0x12, 0x34, 0x56, 0x78, 0x90, 0x01},
			structured: &ie.MACAddressesDetectedFields{
				NumberOfMACAddresses: 1,
				MACAddresses:         []net.HardwareAddr{{0x12, 0x34, 0x56, 0x78, 0x90, 0x01}},
			},
		},
		{
			description: "MACMissing",
			serialized:  []byte{0x02, 0x12, 0x34, 0x56, 0x78, 0x90, 0x01},
			err:         io.ErrUnexpectedEOF,
		},
		{
			description: "CTAGOverrunsPayload",
			serialized:  []byte{0x00, 0x03, 0x01},
			err:         io.ErrUnexpectedEOF,
		},
	}

	for _, c := range cases {
		t.Run(c.description, func(t *testing.T) {
			got, err := ie.ParseMACAddressesDetectedFields(c.serialized)
			if !errors.Is(err, c.err) {
				t.Fatalf("got error %v, want %v", err, c.err)
			}
			if diff := cmp.Diff(c.structured, got); diff != "" {
				t.Error(diff)
			}
		})
	}
}
