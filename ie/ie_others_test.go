// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/wmnsk/go-pfcp/ie"
)

func TestOffendingIE(t *testing.T) {
	structured := ie.NewOffendingIE(ie.Cause)
	decoded := ie.Cause

	got, err := structured.OffendingIE()
	if err != nil {
		t.Fatal(err)
	}

	if diff := cmp.Diff(got, decoded); diff != "" {
		t.Error(diff)
	}
}
