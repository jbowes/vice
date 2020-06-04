package vice

import (
	"errors"
	"testing"
)

func TestForError(t *testing.T) {
	for _, v := range vices {
		err := New(v, "test error")
		if v != ForError(err) {
			t.Fatalf("expected vice from error to equal error's specified vice")
		}
	}

	if ForError(errors.New("non vice error")) != NoVice {
		t.Fatalf("expected non vice error to return NoVice type")
	}
}
