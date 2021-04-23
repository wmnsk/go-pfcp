// Copyright 2019-2021 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package message_test

import (
	"testing"

	"github.com/wmnsk/go-pfcp/message"

	"github.com/wmnsk/go-pfcp/internal/testutil"
)

func TestHeader(t *testing.T) {
	cases := []testutil.TestCase{
		{
			Description: "Node",
			Structured: message.NewHeader(
				1, 0, 0, 0,
				0, // Message type
				0, // SEID
				0, // Sequence Number
				0,
				[]byte{ // Payload: ApplicationID IE
					0x00, 0x18, 0x00, 0x07, 0x67, 0x6f, 0x2d, 0x70, 0x66, 0x63, 0x70,
				},
			),
			Serialized: []byte{
				0x20, 0x00, 0x00, 0x0f, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x18, 0x00, 0x07, 0x67, 0x6f, 0x2d, 0x70, 0x66, 0x63, 0x70,
			},
		}, {
			Description: "Session",
			Structured: message.NewHeader(
				1, 0, 0, 1,
				50,                 // Message type
				0xffffffffffffffff, // SEID
				0xdadada,           // Sequence Number
				0,
				[]byte{ // Payload: ApplicationID IE
					0x00, 0x18, 0x00, 0x07, 0x67, 0x6f, 0x2d, 0x70, 0x66, 0x63, 0x70,
				},
			),
			Serialized: []byte{
				0x21, 0x32, 0x00, 0x17, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
				0xda, 0xda, 0xda, 0x00,
				0x00, 0x18, 0x00, 0x07, 0x67, 0x6f, 0x2d, 0x70, 0x66, 0x63, 0x70,
			},
		},
	}

	testutil.Run(t, cases, func(b []byte) (testutil.Serializable, error) {
		v, err := message.ParseHeader(b)
		if err != nil {
			return nil, err
		}
		return v, nil
	})
}
