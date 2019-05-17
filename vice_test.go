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
