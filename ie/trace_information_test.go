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

func TestParseTraceInformationFields(t *testing.T) {
	v4 := ie.NewTraceInformationFields("123", "45", "abc", []byte{1, 2}, 7, []byte{3, 4}, net.ParseIP("10.0.0.1"))
	v4Bytes, err := v4.Marshal()
	if err != nil {
		t.Fatal(err)
	}
	v6 := ie.NewTraceInformationFields("123", "45", "abc", []byte{1}, 7, []byte{3}, net.ParseIP("2001::1"))
	v6Bytes, err := v6.Marshal()
	if err != nil {
		t.Fatal(err)
	}
	empty := ie.NewTraceInformationFields("123", "45", "abc", nil, 7, nil, net.ParseIP("10.0.0.1"))
	emptyBytes, err := empty.Marshal()
	if err != nil {
		t.Fatal(err)
	}

	cases := []struct {
		description string
		serialized  []byte
		structured  *ie.TraceInformationFields
		err         error
	}{
		{
			description: "IPv4",
			serialized:  v4Bytes,
			structured:  v4,
		},
		{
			description: "IPv6",
			serialized:  v6Bytes,
			structured:  v6,
		},
		{
			description: "NoEventsOrInterfaces",
			serialized:  emptyBytes,
			structured:  empty,
		},
		{
			description: "TriggeringEventsLengthMissing",
			serialized:  v4Bytes[:6],
			err:         io.ErrUnexpectedEOF,
		},
		{
			description: "SessionTraceDepthMissing",
			serialized:  v4Bytes[:9],
			err:         io.ErrUnexpectedEOF,
		},
	}

	for _, c := range cases {
		t.Run(c.description, func(t *testing.T) {
			got, err := ie.ParseTraceInformationFields(c.serialized)
			if !errors.Is(err, c.err) {
				t.Fatalf("got error %v, want %v", err, c.err)
			}
			if diff := cmp.Diff(c.structured, got); diff != "" {
				t.Error(diff)
			}
		})
	}
}

func TestTraceInformationMarshalToShortBuffer(t *testing.T) {
	f := ie.NewTraceInformationFields("123", "45", "abc", nil, 7, nil, net.ParseIP("10.0.0.1"))
	if err := f.MarshalTo(make([]byte, 6)); !errors.Is(err, io.ErrUnexpectedEOF) {
		t.Errorf("got error %v, want %v", err, io.ErrUnexpectedEOF)
	}
}
