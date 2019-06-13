package vice

// Vice is the type used to enumerate all common behaviours for errors.
type Vice uint64

/*
The common error behaviours.

Excluding NoVice, each value corresponds to an interface with a single method
of the same name as the value that returns a bool.

For example, the Closed Vice is used to create and check for errors that
implement:

	interface {
		Closed() bool
	}
*/
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
