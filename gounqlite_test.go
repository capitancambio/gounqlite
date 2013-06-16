package gounqlite

import (
	"testing"
)

func TestErrnoError(t *testing.T) {
	for k, v := range errText {
		if s := k.Error(); s != v {
			t.Errorf("Errno(%d).Error() = %s, expected %s", k, s, v)
		}
	}
}
