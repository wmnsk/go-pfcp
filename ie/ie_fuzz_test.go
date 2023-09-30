package ie_test

import (
	"testing"

	"github.com/wmnsk/go-pfcp/ie"
)

func FuzzParse(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte) {
		if _, err := ie.Parse(b); err != nil {
			t.Skip()
		}
	})
}

func FuzzParseMultiIEs(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte) {
		ie.SetIsGroupedFun(func(t uint16) bool { return true })
		if _, err := ie.ParseMultiIEs(b); err != nil {
			t.Skip()
		}
	})
}

func FuzzValueAs(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte) {
		for typ := uint16(0); typ <= 65535; typ++ {
			i := ie.New(typ, b)
			if _, err := i.ValueAsUint8(); err != nil {
				t.Skip()
			}
			if _, err := i.ValueAsUint16(); err != nil {
				t.Skip()
			}
			if _, err := i.ValueAsUint32(); err != nil {
				t.Skip()
			}
			if _, err := i.ValueAsUint64(); err != nil {
				t.Skip()
			}
			if _, err := i.ValueAsString(); err != nil {
				t.Skip()
			}
			if _, err := i.ValueAsFQDN(); err != nil {
				t.Skip()
			}
			ie.SetIsGroupedFun(func(t uint16) bool { return true })
			if _, err := i.ValueAsGrouped(); err != nil {
				t.Skip()
			}
		}
	})
}
