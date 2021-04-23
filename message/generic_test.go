// Copyright 2019-2021 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package message_test

import (
	"testing"

	"github.com/wmnsk/go-pfcp/ie"
	"github.com/wmnsk/go-pfcp/message"

	"github.com/wmnsk/go-pfcp/internal/testutil"
)

func TestGeneric(t *testing.T) {
	cases := []testutil.TestCase{
		{
			Description: "Normal",
			Structured: message.NewGeneric(
				message.MsgTypeSessionEstablishmentRequest,
				testutil.TestBearerInfo.SEID, testutil.TestBearerInfo.Seq,
				ie.NewApplicationID("go-pfcp"),
			),
			Serialized: []byte{
				0x21, 0x32, 0x00, 0x17, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88,
				0x00, 0x00, 0x01, 0x00,
				0x00, 0x18, 0x00, 0x07, 0x67, 0x6f, 0x2d, 0x70, 0x66, 0x63, 0x70,
			},
		}, {
			Description: "No-SEID",
			Structured: message.NewGenericWithoutSEID(
				message.MsgTypeHeartbeatRequest,
				testutil.TestBearerInfo.Seq,
				ie.NewApplicationID("go-pfcp"),
			),
			Serialized: []byte{
				0x20, 0x01, 0x00, 0x0f, 0x00, 0x00, 0x01, 0x00,
				0x00, 0x18, 0x00, 0x07, 0x67, 0x6f, 0x2d, 0x70, 0x66, 0x63, 0x70,
			},
		},
	}

	testutil.Run(t, cases, func(b []byte) (testutil.Serializable, error) {
		v, err := message.ParseGeneric(b)
		if err != nil {
			return nil, err
		}
		v.Payload = nil
		return v, nil
	})
}
