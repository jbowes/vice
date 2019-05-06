package vice_test

import (
	"fmt"

	"github.com/jbowes/vice"
)

func ExampleIs() {
	err := vice.New(vice.Timeout, "request timed out")

	fmt.Println(vice.Is(err, vice.Timeout))
	// Output: true
}

func ExampleIs_error_chain() {
	err := vice.New(vice.Timeout, "request timed out")
	err2 := vice.Wrap(err, vice.Closed, "connection closed")

	fmt.Println(vice.Is(err2, vice.Timeout))
	// Output: true
}

func ExampleSeal() {
	err := vice.New(vice.Timeout, "request timed out")

	// Seal prevents programmatic inspection of any errors lower in the chain
	err2 := vice.Seal(err, vice.Closed, "connection closed")

	fmt.Println(vice.Is(err2, vice.Timeout))
	// Output: false
}
