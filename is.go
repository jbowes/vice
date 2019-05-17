package vice

import "golang.org/x/xerrors"

// Is reports whether any error in err's chain implements the behaviour
// described by the provided Vice, v.
func Is(err error, v Vice) bool {
	switch v {
	case Timeout:
		return isTimeout(err)
	case Temporary:
		return isTemporary(err)
	case Closed:
		return isClosed(err)
	case AuthRequired:
		return isAuthRequired(err)
	case AuthFailed:
		return isAuthFailed(err)
	case Permission:
		return isPermission(err)
	case Conflict:
		return isConflict(err)
	case InvalidArgument:
		return isInvalidArgument(err)
	case NotFound:
		return isNotFound(err)
	default:
		return false
	}
}

func isTimeout(err error) bool {
	var e interface{ Timeout() bool }
	if xerrors.As(err, &e) {
		return e.Timeout()
	}
	return false
}

func isTemporary(err error) bool {
	var e interface{ Temporary() bool }
	if xerrors.As(err, &e) {
		return e.Temporary()
	}
	return false
}

func isClosed(err error) bool {
	var e interface{ Closed() bool }
	if xerrors.As(err, &e) {
		return e.Closed()
	}
	return false
}

func isAuthRequired(err error) bool {
	var e interface{ AuthRequired() bool }
	if xerrors.As(err, &e) {
		return e.AuthRequired()
	}
	return false
}

func isAuthFailed(err error) bool {
	var e interface{ AuthFailed() bool }
	if xerrors.As(err, &e) {
		return e.AuthFailed()
	}
	return false
}

func isPermission(err error) bool {
	var e interface{ Permission() bool }
	if xerrors.As(err, &e) {
		return e.Permission()
	}
	return false
}

func isConflict(err error) bool {
	var e interface{ Conflict() bool }
	if xerrors.As(err, &e) {
		return e.Conflict()
	}
	return false
}

func isInvalidArgument(err error) bool {
	var e interface{ InvalidArgument() bool }
	if xerrors.As(err, &e) {
		return e.InvalidArgument()
	}
	return false
}

func isNotFound(err error) bool {
	var e interface{ NotFound() bool }
	if xerrors.As(err, &e) {
		return e.NotFound()
	}
	return false
}
