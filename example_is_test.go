package vice_test

import (
	"fmt"

	"github.com/jbowes/vice"
)

// DatabaseError implements the error interface, and the Closed error value
// interface, but is not an error generated or defined in vice.
type DatabaseError struct {
	closed bool
}

func (d *DatabaseError) Error() string { return "database error" }

func (d *DatabaseError) Closed() bool { return d.closed }

func someDatabaseOperation() error {
	return &DatabaseError{closed: true}
}

func ExampleIs_checkingOtherImplementors() {
	if err := someDatabaseOperation(); vice.Is(err, vice.Closed) {
		fmt.Println("database is already closed")
	}
	// Output: database is already closed
}
