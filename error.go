package vice

import (
	"github.com/jbowes/vice/skip"
)

// New returns an error that formats as the given text, and implements the
// behaviour described by the given Vice. The argument skip is the number of
// frames to skip over. Caller(0) returns the frame for the caller of New.
//
// The returned error contains a Frame set to the caller's location and
// implements Formatter to show this information when printed with details.
//
// This func is intended to be used for implementing APIs on top of vice.
func New(v Vice, text string) error { return skip.New(skip.Vice(v), 1, text) }

// Errorf returns an error that formats as the format specifier text, and
// implements the behaviour described by the given Vice. The argument skip is
// the number of frames to skip over. Caller(0) returns the frame for the
// caller of Errorf.
//
// The returned error contains a Frame set to the caller's location and
// implements Formatter to show this information when printed with details.
//
// This func is intended to be used for implementing APIs on top of vice.
func Errorf(v Vice, format string, a ...interface{}) error {
	return skip.Errorf(skip.Vice(v), 1, format, a...)
}

// Wrap returns an error wrapping err with the supplied message, and a frame
// from the caller's stack. The returned error implements the behaviour
// described by the given Vice. The argument skip is the number of frames to
// skip over. Caller(0) returns the frame for the caller of Wrap. If err is
// nil, Wrap returns nil.
//
// The error returned implments the Unwrap method, for programatically
// extracting the error chain.
//
// This func is intended to be used for implementing APIs on top of vice.
func Wrap(err error, v Vice, text string) error {
	return skip.Wrap(err, skip.Vice(v), 1, text)
}

// Wrapf returns an error wrapping err with the supplied format specifier, and
// a frame from the caller's stack. The returned error implements the behaviour
// described by the given Vice. The argument skip is the number of frames to
// skip over. Caller(0) returns the frame for the caller of Wrapf. If err is
// nil, Wrapf returns nil.
//
// The error returned implments the Unwrap method, for programatically
// extracting the error chain.
//
// This func is intended to be used for implementing APIs on top of vice.
func Wrapf(err error, v Vice, format string, a ...interface{}) error {
	return skip.Wrapf(err, skip.Vice(v), 1, format, a...)
}

// Seal returns an error wrapping err with the supplied message, and a frame
// from the caller's stack. The returned error implements the behaviour
// described by the given Vice. The argument skip is the number of frames to
// skip over. Caller(0) returns the frame for the caller of Seal. If err is
// nil, Seal returns nil.
//
// The error returned does not implment the Unwrap method.
//
// This func is intended to be used for implementing APIs on top of vice.
func Seal(err error, v Vice, text string) error {
	return skip.Seal(err, skip.Vice(v), 1, text)
}

// Sealf returns an error wrapping err with the supplied format specifier, and
// a frame from the caller's stack. The returned error implements the behaviour
// described by the given Vice. The argument skip is the number of frames to
// skip over. Caller(0) returns the frame for the caller of Sealf. If err is
// nil, Sealfreturns nil.
//
// The error returned does not implment the Unwrap method.
//
// This func is intended to be used for implementing APIs on top of vice.
func Sealf(err error, v Vice, format string, a ...interface{}) error {
	return skip.Sealf(err, skip.Vice(v), 1, format, a...)
}
