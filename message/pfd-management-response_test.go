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

func TestPFDManagementResponse(t *testing.T) {
	cases := []testutil.TestCase{
		{
			Description: "Normal",
			Structured: message.NewPFDManagementResponse(
				seq,
				ie.NewCause(ie.CauseRequestAccepted),
				ie.NewOffendingIE(ie.Cause),
			),
			Serialized: []byte{
				0x20, 0x04, 0x00, 0x0f, 0x11, 0x22, 0x33, 0x00,
				0x00, 0x13, 0x00, 0x01, 0x01,
				0x00, 0x28, 0x00, 0x02, 0x00, 0x13,
			},
		},
	}

	testutil.Run(t, cases, func(b []byte) (testutil.Serializable, error) {
		v, err := message.ParsePFDManagementResponse(b)
		if err != nil {
			return nil, err
		}
		v.Payload = nil
		return v, nil
	})
}
