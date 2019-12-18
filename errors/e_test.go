package errors

import (
	"io"
	"testing"
)

func TestIs(t *testing.T) {
	e := ErrCodePacketIncomplete.New("asd")
	if !e.EqualRecursive((ErrCodePacketIncomplete)) {
		t.Fatal("EqualRecursive() failed")
	}

	ex := New("aa").Nest(e, io.EOF)
	if !ex.EqualRecursive((ErrCodePacketIncomplete)) {
		t.Fatal("nested EqualRecursive() failed")
	}
}
