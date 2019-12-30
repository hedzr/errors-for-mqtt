package errors

import (
	"regexp"
	"testing"
)

func TestCloseReasonConstants(t *testing.T) {
	for i := CloseNormal; i <= CloseUnsupportedWildcardSubscription; i++ {
		str := i.String()
		if matched, err := regexp.Match(`CloseReason\(\d+\)`, []byte(str)); err == nil && !matched {
			c := CloseNormal.Parse(str)
			t.Logf("  - %x (%d) => %q [%v]", int(c), int(c), str, c.IsValidReason())
		}
	}
}
