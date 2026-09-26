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

func TestParseUserPlaneIPResourceInformationFields(t *testing.T) {
	marshal := func(f *ie.UserPlaneIPResourceInformationFields) []byte {
		b, err := f.Marshal()
		if err != nil {
			t.Fatal(err)
		}
		return b
	}
	ni := ie.NewUserPlaneIPResourceInformationFields(0x21, 0, "127.0.0.1", "", "some.apn", 0)
	niSI := ie.NewUserPlaneIPResourceInformationFields(0x61, 0, "127.0.0.1", "", "some.apn", 3)
	si := ie.NewUserPlaneIPResourceInformationFields(0x41, 0, "127.0.0.1", "", "", 3)

	cases := []struct {
		description string
		serialized  []byte
		structured  *ie.UserPlaneIPResourceInformationFields
		err         error
	}{
		{
			description: "NetworkInstance",
			serialized:  marshal(ni),
			structured:  ni,
		},
		{
			description: "NetworkInstanceAndSourceInterface",
			serialized:  marshal(niSI),
			structured:  niSI,
		},
		{
			description: "SourceInterface",
			serialized:  marshal(si),
			structured:  si,
		},
		{
			description: "SourceInterfaceMissing",
			serialized:  []byte{0x48, 0x06},
			err:         io.ErrUnexpectedEOF,
		},
	}

	for _, c := range cases {
		t.Run(c.description, func(t *testing.T) {
			got, err := ie.ParseUserPlaneIPResourceInformationFields(c.serialized)
			if !errors.Is(err, c.err) {
				t.Fatalf("got error %v, want %v", err, c.err)
			}
			if diff := cmp.Diff(c.structured, got); diff != "" {
				t.Error(diff)
			}
		})
	}
}

func TestTEIDRIOnEmptyPayload(t *testing.T) {
	if got := ie.New(ie.UserPlaneIPResourceInformation, nil).TEIDRI(); got != 0 {
		t.Errorf("got %d, want 0", got)
	}
}
