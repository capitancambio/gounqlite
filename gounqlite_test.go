package gounqlite

import (
	"bytes"
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
		t.Fatalf("Open(\":mem:\") error: %s", err)
	}
	err = c.Close()
	if err != nil {
		t.Fatalf("c.Close() error: %s", err)
	}
}

type keyValuePair struct {
	key   []byte
	value []byte
}

var keyValuePairs = []keyValuePair{
	{[]byte("key"), []byte("value")},
	{[]byte("hello"), []byte("世界")},
	{[]byte("hello"), []byte("world")},
	{[]byte("gordon"), []byte("gopher")},
	{[]byte{'f', 'o', 'o'}, []byte{'b', 'a', 'r'}},
	{[]byte{'\000'}, []byte{42}},
}

func TestStore(t *testing.T) {
	c, err := Open(":mem:")
	if err != nil {
		t.Fatalf("Open(\":mem:\") error: %s", err)
	}
	defer c.Close()

	for _, p := range keyValuePairs {
		if err := c.Store(p.key, p.value); err != nil {
			t.Fatalf("c.Store(%v, %v) error: %v", p.key, p.value, err)
		}
		b, err := c.Fetch(p.key)
		if err != nil {
			t.Fatalf("c.Fetch(%v) error: %v", p.key, err)
		}
		if !bytes.Equal(b, p.value) {
			t.Errorf("c.Fetch(%v) = %v, expected %v", p.key, b, p.value)
		}
	}
}

func TestDelete(t *testing.T) {
	c, err := Open(":mem:")
	if err != nil {
		t.Fatalf("Open(\":mem:\") error: %s", err)
	}
	defer c.Close()

	for _, p := range keyValuePairs {
		// Test delete non-existing records.
		if err := c.Delete(p.key); err != ErrNotFound {
			t.Fatalf("c.Delete(%v) = %v, expected %v", p.key, err, ErrNotFound)
		}
		// Store records.
		if err := c.Store(p.key, p.value); err != nil {
			t.Fatalf("c.Store(%v, %v) error: %v", p.key, p.value, err)
		}
		// Fetch and compare records to make sure that they were stored.
		b, err := c.Fetch(p.key)
		if err != nil {
			t.Fatalf("c.Fetch(%v) error: %v", p.key, err)
		}
		if !bytes.Equal(b, p.value) {
			t.Errorf("c.Fetch(%v) = %v, expected %v", p.key, b, p.value)
		}
		// Test delete records.
		if err := c.Delete(p.key); err != nil {
			t.Fatalf("c.Delete(%v) = %v, expected %v", p.key, err, nil)
		}
	}
}
