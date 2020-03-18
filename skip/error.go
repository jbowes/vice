package skip

import (
	"fmt"

	"golang.org/x/xerrors"
)

// New returns an error that formats as the given text, and implements the
// behaviour described by the given Vice. The argument skip is the number of
// frames to skip over. Caller(0) returns the frame for the caller of New.
//
// The returned error contains a Frame set to the caller's location and
// implements Formatter to show this information when printed with details.
//
// This func is intended to be used for implementing APIs on top of vice.
func New(v Vice, skip uint, text string) error {
	return sealed(&sealError{msg: text}, v, skip)
}

// Errorf returns an error that formats as the format specifier, and implements
// the behaviour described by the given Vice. The argument skip is  the number
// of frames to skip over. Caller(0) returns the frame for the caller of Errorf.
//
// The returned error contains a Frame set to the caller's location and
// implements Formatter to show this information when printed with details.
//
// This func is intended to be used for implementing APIs on top of vice.
func Errorf(v Vice, skip uint, format string, a ...interface{}) error {
	return sealed(&sealError{msg: fmt.Sprintf(format, a...)}, v, skip)
}

// Wrap returns an error wrapping err with the supplied text, and a frame
// from the caller's stack. The returned error implements the behaviour
// described by the given Vice. The argument skip is the number of frames to
// skip over. Caller(0) returns the frame for the caller of Wrap. If err is
// nil, Wrap returns nil.
//
// The error returned implments the Unwrap method, for programatically
// extracting the error chain.
//
// This func is intended to be used for implementing APIs on top of vice.
func Wrap(err error, v Vice, skip uint, text string) error {
	if err == nil {
		return nil
	}
	return wrapped(&wrapError{sealError{err: err, msg: text}}, v, skip)
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
func Wrapf(err error, v Vice, skip uint, format string, a ...interface{}) error {
	if err == nil {
		return nil
	}
	return wrapped(&wrapError{sealError{err: err, msg: fmt.Sprintf(format, a...)}}, v, skip)
}

// Seal returns an error wrapping err with the supplied text, and a frame
// from the caller's stack. The returned error implements the behaviour
// described by the given Vice. The argument skip is the number of frames to
// skip over. Caller(0) returns the frame for the caller of Seal. If err is
// nil, Seal returns nil.
//
// The error returned does not implment the Unwrap method.
//
// This func is intended to be used for implementing APIs on top of vice.
func Seal(err error, v Vice, skip uint, text string) error {
	if err == nil {
		return nil
	}
	return sealed(&sealError{err: err, msg: text}, v, skip)
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
func Sealf(err error, v Vice, skip uint, format string, a ...interface{}) error {
	if err == nil {
		return nil
	}
	return sealed(&sealError{err: err, msg: fmt.Sprintf(format, a...)}, v, skip)
}

func sealed(serr *sealError, v Vice, skip uint) error {
	serr.frame = xerrors.Caller(int(2 + skip))
	switch v {
	case Timeout:
		return &timeoutSeal{*serr}
	case Temporary:
		return &temporarySeal{*serr}
	case Closed:
		return &closedSeal{*serr}
	case AuthRequired:
		return &authRequiredSeal{*serr}
	case AuthFailed:
		return &authFailedSeal{*serr}
	case Permission:
		return &permissionSeal{*serr}
	case Conflict:
		return &conflictSeal{*serr}
	case InvalidArgument:
		return &invalidArgumentSeal{*serr}
	case NotFound:
		return &notFoundErrorSeal{*serr}
	case Internal:
		return &internalSeal{*serr}
	case Canceled:
		return &canceledSeal{*serr}
	case PreconditionFailed:
		return &preconditionFailedSeal{*serr}
	default:
		return serr
	}
}

func wrapped(werr *wrapError, v Vice, skip uint) error {
	werr.frame = xerrors.Caller(int(2 + skip))
	switch v {
	case Timeout:
		return &timeoutWrap{*werr}
	case Temporary:
		return &temporaryWrap{*werr}
	case Closed:
		return &closedWrap{*werr}
	case AuthRequired:
		return &authRequiredWrap{*werr}
	case AuthFailed:
		return &authFailedWrap{*werr}
	case Permission:
		return &permissionWrap{*werr}
	case Conflict:
		return &conflictWrap{*werr}
	case InvalidArgument:
		return &invalidArgumentWrap{*werr}
	case NotFound:
		return &notFoundErrorWrap{*werr}
	case Internal:
		return &internalWrap{*werr}
	case Canceled:
		return &canceledWrap{*werr}
	case PreconditionFailed:
		return &preconditionFailedWrap{*werr}
	default:
		return werr
	}
}

type sealError struct {
	msg   string
	err   error
	frame xerrors.Frame
}

func (e *sealError) Error() string              { return fmt.Sprint(e) }
func (e *sealError) Format(s fmt.State, v rune) { xerrors.FormatError(e, s, v) }

func (e *sealError) FormatError(p xerrors.Printer) (next error) {
	p.Print(e.msg)
	e.frame.Format(p)
	return e.err
}

type wrapError struct {
	sealError
}

func (e *wrapError) Unwrap() error { return e.err }
