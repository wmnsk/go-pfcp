// Copyright 2019-2020 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package message_test

import (
	"github.com/wmnsk/go-pfcp/ie"
	"github.com/wmnsk/go-pfcp/internal/testutil"
	"github.com/wmnsk/go-pfcp/message"
	"testing"
	"time"
)

func TestAssociationSetupResponse(t *testing.T) {

	nodeID := ie.NewNodeID("172.55.55.101", "", "")
	rets := ie.NewRecoveryTimeStamp(time.Date(2020, 3, 3, 12, 22, 22, 22, time.Local))
	cause := ie.NewCause(ie.CauseRequestAccepted)
	cp := ie.NewCPFunctionFeatures(0x00)

	cases := []testutil.TestCase{
		{
			Description: "Normal",
			Structured: message.NewAssociationSetupResponse(
				nodeID, cause, rets, cp,
			),
			Serialized: []byte{
				0x20, 0x06, 0x00, 0x1f, 0x00, 0x00, 0x00, 0x00, 0x00, 0x3c, 0x00, 0x05, 0x00, 0xac, 0x37, 0x37, 0x65, 0x00, 0x13, 0x00, 0x01, 0x01, 0x00, 0x60, 0x00, 0x04, 0xe2, 0x08, 0x59, 0xfe, 0x00, 0x59, 0x00, 0x01, 0x00,
			},
		},
	}

	testutil.Run(t, cases, func(b []byte) (testutil.Serializable, error) {
		v, err := message.ParseAssociationSetupResponse(b)
		if err != nil {
			return nil, err
		}
		v.Payload = nil
		return v, nil
	})
}
