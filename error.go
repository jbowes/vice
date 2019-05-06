package vice

import (
	"github.com/jbowes/vice/skip"
)

func New(v Vice, text string) error { return skip.New(skip.Vice(v), 1, text) }

func Errorf(v Vice, format string, a ...interface{}) error {
	return skip.Errorf(skip.Vice(v), 1, format, a...)
}

func Wrap(err error, v Vice, text string) error { return skip.Wrap(err, skip.Vice(v), 1, text) }

func Wrapf(err error, v Vice, format string, a ...interface{}) error {
	return skip.Wrapf(err, skip.Vice(v), 1, format, a...)
}

func Seal(err error, v Vice, text string) error { return skip.Seal(err, skip.Vice(v), 1, text) }

func Sealf(err error, v Vice, format string, a ...interface{}) error {
	return skip.Sealf(err, skip.Vice(v), 1, format, a...)
}
