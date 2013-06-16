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

func TestOpenCloseInMemoryDatabase(t *testing.T) {
	c, err := Open(":mem:")
	if err != nil {
		t.Errorf("got unexpected error for Open(\":mem:\"): %s", err)
	}
	err = c.Close()
	if err != nil {
		t.Errorf("got unexpected error for c.Close(): %s", err)
	}
}
