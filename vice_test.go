package vice

import (
	"errors"
	"testing"
)

func TestVice(t *testing.T) {
	vices := []Vice{
		Timeout,
		Temporary,
		Closed,
		AuthRequired,
		AuthFailed,
		Permission,
		Conflict,
		InvalidArgument,
		NotFound,
		Internal,
	}

	for _, v := range vices {
		err := New(v, "test")
		if !Is(err, v) {
			t.Error("Did not create correct error form")
		}

		inner := errors.New("inside")

		err = Wrap(inner, v, "wrapping")
		if !Is(err, v) {
			t.Error("Did not create correct error form")
		}
	}
}

func TestWrapNil(t *testing.T) {
	inner := errors.New("no auth")
	wrapped := Wrap(inner, AuthRequired, "wrapping")

	if wrapped == nil {
		t.Error("Wrapped error should not be nil")
	}

	maybeError := func() error {
		return nil
	}()
	wrapped = Wrap(maybeError, Conflict, "try wrapping")
	if wrapped != nil {
		t.Error("Wrapping a nil error should return nil")
	}
}
