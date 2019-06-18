package vice

import (
	"errors"
	"testing"
)

func TestVice(t *testing.T) {
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

func TestNoVice(t *testing.T) {
	t.Run("with Vice error", func(t *testing.T) {
		for i, v := range vices {
			err := New(v, "test")
			if Is(err, NoVice) {
				t.Errorf("%d element should be concrete Vice type", i)
			}
		}
	})

	t.Run("with NoVice error", func(t *testing.T) {
		if !Is(errors.New("test"), NoVice) {
			t.Errorf("Expected custom error to be NoVice")
		}

		if !Is(New(NoVice, "test"), NoVice) {
			t.Errorf("Expected NoVice error to be NoVice")
		}
	})
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
