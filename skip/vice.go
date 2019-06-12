package skip

// Vice is the type used to enumerate all common behaviours for errors.
// It is mirrored with the main vice package.
type Vice uint64

// These values are mirrored with the main vice package.
const (
	NoVice  Vice = 0
	Timeout Vice = 1 << iota
	Temporary
	Closed
	AuthRequired
	AuthFailed
	Permission
	Conflict
	InvalidArgument
	NotFound
	Internal
)

// implementations of all the error types
type timeoutSeal struct{ sealError }

func (timeoutSeal) Timeout() bool { return true }

type temporarySeal struct{ sealError }

func (temporarySeal) Temporary() bool { return true }

type closedSeal struct{ sealError }

func (closedSeal) Closed() bool { return true }

type authRequiredSeal struct{ sealError }

func (authRequiredSeal) AuthRequired() bool { return true }

type authFailedSeal struct{ sealError }

func (authFailedSeal) AuthFailed() bool { return true }

type permissionSeal struct{ sealError }

func (permissionSeal) Permission() bool { return true }

type conflictSeal struct{ sealError }

func (conflictSeal) Conflict() bool { return true }

type invalidArgumentSeal struct{ sealError }

func (invalidArgumentSeal) InvalidArgument() bool { return true }

type notFoundErrorSeal struct{ sealError }

func (notFoundErrorSeal) NotFound() bool { return true }

type internalSeal struct{ sealError }

func (internalSeal) Internal() bool { return true }

// wrapped implmentations of all the error types

type timeoutWrap struct{ wrapError }

func (timeoutWrap) Timeout() bool { return true }

type temporaryWrap struct{ wrapError }

func (temporaryWrap) Temporary() bool { return true }

type closedWrap struct{ wrapError }

func (closedWrap) Closed() bool { return true }

type authRequiredWrap struct{ wrapError }

func (authRequiredWrap) AuthRequired() bool { return true }

type authFailedWrap struct{ wrapError }

func (authFailedWrap) AuthFailed() bool { return true }

type permissionWrap struct{ wrapError }

func (permissionWrap) Permission() bool { return true }

type conflictWrap struct{ wrapError }

func (conflictWrap) Conflict() bool { return true }

type invalidArgumentWrap struct{ wrapError }

func (invalidArgumentWrap) InvalidArgument() bool { return true }

type notFoundErrorWrap struct{ wrapError }

func (notFoundErrorWrap) NotFound() bool { return true }

type internalWrap struct{ wrapError }

func (internalWrap) Internal() bool { return true }
