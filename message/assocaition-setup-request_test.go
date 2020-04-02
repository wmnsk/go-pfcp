// Copyright 2019-2020 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.
package message_test

import (
	"testing"
	"time"

	"github.com/wmnsk/go-pfcp/ie"
	"github.com/wmnsk/go-pfcp/internal/testutil"
	"github.com/wmnsk/go-pfcp/message"
)

func TestAssociationSetupRequest(t *testing.T) {
	cases := []testutil.TestCase{
		{
			Description: "Normal",
			Structured:  message.NewAssociationSetupRequest(ie.NewNodeID("172.55.55.102", "", ""), ie.NewRecoveryTimeStamp(time.Now()), ie.NewUPFunctionFeatures(0x10, 0x00)),

			Serialized: []byte{
				0x20, 0x05, 0x00, 0x1b, 0x00, 0x00, 0x00, 0x00, 0x00, 0x3c, 0x00, 0x05, 0x00, 0xac, 0x37, 0x37,
				0x66, 0x00, 0x60, 0x00, 0x04, 0xe2, 0x30, 0x2c, 0xad, 0x00, 0x2b, 0x00, 0x02, 0x10, 0x00,
			},
		},
	}

	testutil.Run(t, cases, func(b []byte) (testutil.Serializable, error) {
		v, err := message.ParseAssociationSetupRequest(b)
		if err != nil {
			return nil, err
		}
		v.Payload = nil
		return v, nil
	})
}
