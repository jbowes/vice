/*
Package vice provides bad behaviours for Go 2/1.13+ error values.

vice defines common types of errors, as seen in existing packages, as expressed
through methods on the errors of the form:
	IsXXX() bool

Errors of these types can be created and checked by vice. When used with Go 2
error value wrapping, intermediate code layers don't have to know about the
behaviours or types of errors, and may add additional information, without
altering how the error is handled in the topmost layer:
	// Create an error when your database can't find a user record
	err := vice.New(vice.NotFound, "user not found")

	// many intermediate layers of code, passing the error on, and wrapping it
	// ...

	// In your central request handler middleware
	if vice.Is(err, vice.NotFound) {
		// respond appropriately to the client
	}
*/
package vice // import "github.com/jbowes/vice"
