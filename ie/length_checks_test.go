// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie_test

import (
	"errors"
	"io"
	"testing"

	"github.com/wmnsk/go-pfcp/ie"
)

func TestSubsequentTimeQuotaShortPayload(t *testing.T) {
	if _, err := ie.New(ie.SubsequentTimeQuota, []byte{0x00, 0x01}).SubsequentTimeQuota(); !errors.Is(err, io.ErrUnexpectedEOF) {
		t.Errorf("got error %v, want %v", err, io.ErrUnexpectedEOF)
	}
}

func TestHasMACEmptyPayload(t *testing.T) {
	if ie.New(ie.TSNBridgeID, nil).HasMAC() {
		t.Error("got true, want false")
	}
}
