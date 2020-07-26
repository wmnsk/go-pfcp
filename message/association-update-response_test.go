// Copyright 2019-2020 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package message_test

import (
	"testing"

	"github.com/wmnsk/go-pfcp/ie"
	"github.com/wmnsk/go-pfcp/message"

	"github.com/wmnsk/go-pfcp/internal/testutil"
)

func TestAssociationUpdateResponse(t *testing.T) {
	var seq uint32 = 3
	cases := []testutil.TestCase{
		{
			Description: "Normal",
			Structured: message.NewAssociationUpdateResponse(seq,
				ie.NewNodeID("", "", "go-pfcp.epc.3gppnetwork.org"),
				ie.NewCause(ie.CauseRequestAccepted),
				ie.NewUPFunctionFeatures(0x01, 0x02),
				ie.NewCPFunctionFeatures(0x3f),
			),
			Serialized: []byte{
				0x20, 0x08, 0x00, 0x35, 0x00, 0x00, 0x03, 0x00,
				0x00, 0x3c, 0x00, 0x1d, 0x02, 0x07, 0x67, 0x6f, 0x2d, 0x70, 0x66, 0x63, 0x70, 0x03, 0x65, 0x70, 0x63, 0x0b, 0x33, 0x67, 0x70, 0x70, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x03, 0x6f, 0x72, 0x67,
				0x00, 0x13, 0x00, 0x01, 0x01,
				0x00, 0x2b, 0x00, 0x02, 0x01, 0x02,
				0x00, 0x59, 0x00, 0x01, 0x3f,
			},
		},
	}

	testutil.Run(t, cases, func(b []byte) (testutil.Serializable, error) {
		v, err := message.ParseAssociationUpdateResponse(b)
		if err != nil {
			return nil, err
		}
		v.Payload = nil
		return v, nil
	})
}
