// Copyright 2019-2021 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package message_test

import (
	"testing"

	"github.com/wmnsk/go-pfcp/message"

	"github.com/wmnsk/go-pfcp/internal/testutil"
)

func TestSessionDeletionRequest(t *testing.T) {
	cases := []testutil.TestCase{
		{
			Description: "Normal",
			Structured:  message.NewSessionDeletionRequest(mp, fo, seid, seq, pri),
			Serialized: []byte{
				0x21, 0x36, 0x00, 0x0c, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x11, 0x22, 0x33, 0x00,
			},
		},
	}

	testutil.Run(t, cases, func(b []byte) (testutil.Serializable, error) {
		v, err := message.ParseSessionDeletionRequest(b)
		if err != nil {
			return nil, err
		}
		v.Payload = nil
		return v, nil
	})
}
