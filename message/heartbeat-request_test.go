// Copyright 2019-2021 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package message_test

import (
	"net"
	"testing"
	"time"

	"github.com/wmnsk/go-pfcp/ie"
	"github.com/wmnsk/go-pfcp/message"

	"github.com/wmnsk/go-pfcp/internal/testutil"
)

func TestHeartbeatRequest(t *testing.T) {
	cases := []testutil.TestCase{
		{
			Description: "Normal",
			Structured: message.NewHeartbeatRequest(seq,
				ie.NewRecoveryTimeStamp(time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)),
				ie.NewSourceIPAddress(net.ParseIP("127.0.0.1"), net.ParseIP("2001::1"), 0),
			),
			Serialized: []byte{
				0x20, 0x01, 0x00, 0x25, 0x11, 0x22, 0x33, 0x00,
				0x00, 0x60, 0x00, 0x04, 0xdf, 0xd5, 0x2c, 0x00,
				0x00, 0xc0, 0x00, 0x15,
				0x03,
				0x7f, 0x00, 0x00, 0x01,
				0x20, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01,
			},
		},
	}

	testutil.Run(t, cases, func(b []byte) (testutil.Serializable, error) {
		v, err := message.ParseHeartbeatRequest(b)
		if err != nil {
			return nil, err
		}
		v.Payload = nil
		return v, nil
	})
}
