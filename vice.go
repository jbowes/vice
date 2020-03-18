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
	// NoVice indicates there is an unkown error.
	NoVice Vice = 0

	// Timeout means the operation has timed out before the end of the
	// operation.
	Timeout Vice = 1 << iota

	// Temporary indicates a temporary underlying issue. The consumer should
	// retry.
	Temporary

	// Closed means the consumer closed the connection to the underlying
	// service.
	Closed

	// AuthRequired indicates the operation needs authentication, but none is
	// provided.
	AuthRequired

	// AuthFailed indicates the operation could not be authenticated using the
	// provided authentication schema.
	AuthFailed

	// Permission indicates that the consumer does not have permissions to
	// execute the specified operation.
	Permission

	// Conflict indicates that an attempt to create a resource failed because it
	// already exists.
	Conflict

	// InvalidArgument indicates the client provided bad input. This differs
	// from PreconditionFailed as this is based on the given arguments rather
	// than the state of the system. For example, an invalid HTTP scheme.
	InvalidArgument

	// NotFound indicates that a requested resource was not found.
	NotFound

	// Internal indicates that the service expected different behaviour than it
	// produced. If you see one of these errors, the service is very broken.
	Internal

	// Canceled means the operation has been canceled, typically by the
	// consumer.
	Canceled

	// PreconditionFailed indicates that the state of the system does not meet the
	// requirements to fulfull the request.
	// For example, a directory to be deleted may be non-empty, an rmdir
	// operation is applied to a non-directory, etc.
	PreconditionFailed
)

// vices is a fixed length slice of Vice verbs used to iterate over them
// accordingly. It only contains actual vices, NoVice is not included.
var vices = [...]Vice{
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
	Canceled,
	PreconditionFailed,
}
