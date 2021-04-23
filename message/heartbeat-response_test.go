// Copyright 2019-2021 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package message_test

import (
	"testing"
	"time"

	"github.com/wmnsk/go-pfcp/ie"
	"github.com/wmnsk/go-pfcp/message"

	"github.com/wmnsk/go-pfcp/internal/testutil"
)

func TestHeartbeatResponse(t *testing.T) {
	cases := []testutil.TestCase{
		{
			Description: "Normal",
			Structured: message.NewHeartbeatResponse(seq,
				ie.NewRecoveryTimeStamp(time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)),
			),
			Serialized: []byte{
				0x20, 0x02, 0x00, 0x0c, 0x11, 0x22, 0x33, 0x00,
				0x00, 0x60, 0x00, 0x04, 0xdf, 0xd5, 0x2c, 0x00,
			},
		},
	}

	testutil.Run(t, cases, func(b []byte) (testutil.Serializable, error) {
		v, err := message.ParseHeartbeatResponse(b)
		if err != nil {
			return nil, err
		}
		v.Payload = nil
		return v, nil
	})
}
