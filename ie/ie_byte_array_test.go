// Copyright 2019-2023 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/wmnsk/go-pfcp/ie"
)

func TestByteArrayIEs(t *testing.T) {
	cases := []struct {
		description string
		structured  *ie.IE
		decoded     []byte
		decoderFunc func(*ie.IE) ([]byte, error)
	}{
		{
			description: "ApplyAction/pre-16.3.0",
			structured:  ie.NewApplyAction(0x04), // Flag BUFF is set
			decoded:     []byte{0x04},
			decoderFunc: func(i *ie.IE) ([]byte, error) { return i.ApplyAction() },
		}, {
			description: "ApplyAction/post-16.3.0-compat",
			structured:  ie.NewApplyAction(0x04, 0x00), // Flag BUFF is set
			decoded:     []byte{0x04, 0x00},
			decoderFunc: func(i *ie.IE) ([]byte, error) { return i.ApplyAction() },
		}, {
			description: "ApplyAction/post-16.3.0",
			structured:  ie.NewApplyAction(0x04, 0x02), //Flags BUFF and BDPN are set
			decoded:     []byte{0x04, 0x02},
			decoderFunc: func(i *ie.IE) ([]byte, error) { return i.ApplyAction() },
		}, {
			description: "ApplyAction/pre-16.3.0/CreateFAR",
			structured: ie.NewCreateFAR(
				ie.NewFARID(0xffffffff),
				ie.NewApplyAction(0x04), // Flag BUFF is set
			),
			decoded:     []byte{0x04},
			decoderFunc: func(i *ie.IE) ([]byte, error) { return i.ApplyAction() },
		}, {
			description: "ApplyAction/post-16.3.0-compat/CreateFAR",
			structured: ie.NewCreateFAR(
				ie.NewFARID(0xffffffff),
				ie.NewApplyAction(0x04, 0x00), // Flag BUFF is set
			),
			decoded:     []byte{0x04, 0x00},
			decoderFunc: func(i *ie.IE) ([]byte, error) { return i.ApplyAction() },
		}, {
			description: "ApplyAction/post-16.3.0/CreateFAR",
			structured: ie.NewCreateFAR(
				ie.NewFARID(0xffffffff),
				ie.NewApplyAction(0x04, 0x02), // Flags BUFF and BDPN are set
			),
			decoded:     []byte{0x04, 0x02},
			decoderFunc: func(i *ie.IE) ([]byte, error) { return i.ApplyAction() },
		}, {
			description: "ApplyAction/pre-16.3.0/UpdateFAR",
			structured: ie.NewUpdateFAR(
				ie.NewFARID(0xffffffff),
				ie.NewApplyAction(0x04), // Flag BUFF is set
			),
			decoded:     []byte{0x04},
			decoderFunc: func(i *ie.IE) ([]byte, error) { return i.ApplyAction() },
		}, {
			description: "ApplyAction/post-16.3.0-compat/UpdateFAR",
			structured: ie.NewUpdateFAR(
				ie.NewFARID(0xffffffff),
				ie.NewApplyAction(0x04, 0x00), // Flag BUFF is set
			),
			decoded:     []byte{0x04, 0x00},
			decoderFunc: func(i *ie.IE) ([]byte, error) { return i.ApplyAction() },
		}, {
			description: "ApplyAction/post-16.3.0/UpdateFAR",
			structured: ie.NewUpdateFAR(
				ie.NewFARID(0xffffffff),
				ie.NewApplyAction(0x04, 0x02), // Flags BUFF and BDPN are set
			),
			decoded:     []byte{0x04, 0x02},
			decoderFunc: func(i *ie.IE) ([]byte, error) { return i.ApplyAction() },
		}, {
			description: "CPFunctionFeatures",
			structured:  ie.NewCPFunctionFeatures(0x3f),
			decoded:     []byte{0x3f},
			decoderFunc: func(i *ie.IE) ([]byte, error) { return i.CPFunctionFeatures() },
		},
	}

	for _, c := range cases {
		t.Run(c.description, func(t *testing.T) {
			got, err := c.decoderFunc(c.structured)
			if err != nil {
				t.Fatal(err)
			}

			if diff := cmp.Diff(got, c.decoded); diff != "" {
				t.Error(diff)
			}
		})
	}
}
