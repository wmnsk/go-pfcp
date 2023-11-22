package ie_test

import (
	"io"
	"net"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/wmnsk/go-pfcp/ie"
)

func TestParseMACAddressFields(t *testing.T) {
	cases := []struct {
		description string
		serialized  []byte
		structured  *ie.MACAddressFields
		err         error
	}{
		{
			description: "EmptyPayload",
			serialized:  []byte{},
			structured:  nil,
			err:         io.ErrUnexpectedEOF,
		},
		{
			description: "TooSmallPayload",
			serialized:  []byte{1},
			structured:  nil,
			err:         io.ErrUnexpectedEOF,
		},
		{
			description: "SmallestValidPayload",
			serialized:  []byte{0, 0},
			structured:  &ie.MACAddressFields{},
		},
		{
			description: "SourceMACAddress",
			serialized:  []byte{1, 0, 0, 0, 0, 0, 1},
			structured: &ie.MACAddressFields{
				Flags:            0x1,
				SourceMACAddress: net.HardwareAddr{0, 0, 0, 0, 0, 1},
			},
		},
		{
			description: "SourceMACAddressTooShort",
			serialized:  []byte{1, 0},
			structured:  nil,
			err:         io.ErrUnexpectedEOF,
		},
		{
			description: "DestinationMACAddress",
			serialized:  []byte{2, 0, 0, 0, 0, 0, 1},
			structured: &ie.MACAddressFields{
				Flags:                 0x2,
				DestinationMACAddress: net.HardwareAddr{0, 0, 0, 0, 0, 1},
			},
		},
		{
			description: "DestinationMACAddressTooShort",
			serialized:  []byte{2, 0},
			structured:  nil,
			err:         io.ErrUnexpectedEOF,
		},
		{
			description: "UpperSourceMACAddress",
			serialized:  []byte{4, 0, 0, 0, 0, 0, 1},
			structured: &ie.MACAddressFields{
				Flags:                 0x4,
				UpperSourceMACAddress: net.HardwareAddr{0, 0, 0, 0, 0, 1},
			},
		},
		{
			description: "UpperSourceMACAddressTooShort",
			serialized:  []byte{4, 0},
			structured:  nil,
			err:         io.ErrUnexpectedEOF,
		},
		{
			description: "UpperDestinationMACAddress",
			serialized:  []byte{8, 0, 0, 0, 0, 0, 1},
			structured: &ie.MACAddressFields{
				Flags:                      0x8,
				UpperDestinationMACAddress: net.HardwareAddr{0, 0, 0, 0, 0, 1},
			},
		},
		{
			description: "UpperDestinationMACAddressTooShort",
			serialized:  []byte{4, 0},
			structured:  nil,
			err:         io.ErrUnexpectedEOF,
		},
		{
			description: "AllCombined",
			serialized:  []byte{15, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 3, 0, 0, 0, 0, 0, 4},
			structured: &ie.MACAddressFields{
				Flags:                      0x0f,
				SourceMACAddress:           net.HardwareAddr{0, 0, 0, 0, 0, 1},
				DestinationMACAddress:      net.HardwareAddr{0, 0, 0, 0, 0, 2},
				UpperSourceMACAddress:      net.HardwareAddr{0, 0, 0, 0, 0, 3},
				UpperDestinationMACAddress: net.HardwareAddr{0, 0, 0, 0, 0, 4},
			},
		},
	}
	for _, c := range cases {
		t.Run(c.description, func(t *testing.T) {
			got, err := ie.ParseMACAddressFields(c.serialized)
			if err != c.err {
				t.Errorf("expected error %v but got %v", c.err, err)
				t.Fatal(err)
			}

			if diff := cmp.Diff(got, c.structured); diff != "" {
				t.Error(diff)
			}
		})
	}
}
