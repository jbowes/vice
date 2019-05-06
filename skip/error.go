package skip

import (
	"fmt"

	"golang.org/x/xerrors"
)

func New(v Vice, skip uint, text string) error {
	return sealed(&sealError{msg: text}, v, skip)
}

func Errorf(v Vice, skip uint, format string, a ...interface{}) error {
	return sealed(&sealError{msg: fmt.Sprintf(format, a...)}, v, skip)
}

func Wrap(err error, v Vice, skip uint, text string) error {
	return wrapped(&wrapError{sealError{err: err, msg: text}}, v, skip)
}

func Wrapf(err error, v Vice, skip uint, format string, a ...interface{}) error {
	return wrapped(&wrapError{sealError{err: err, msg: fmt.Sprintf(format, a...)}}, v, skip)
}

func Seal(err error, v Vice, skip uint, text string) error {
	return sealed(&sealError{err: err, msg: text}, v, skip)
}

func Sealf(err error, v Vice, skip uint, format string, a ...interface{}) error {
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
