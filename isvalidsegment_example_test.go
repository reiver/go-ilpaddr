package ilpaddr_test

import (
	"fmt"

	"github.com/reiver/go-ilpaddr"
)

func ExampleIsValidSegment_valid() {

	var segment string = "~ipr"

	result := ilpaddr.IsValidSegment(segment)

	fmt.Printf("SEGMENT: %q\n", segment)
	fmt.Printf("IS-VALID-SEGMENT: %t\n", result)

	// Output:
	// SEGMENT: "~ipr"
	// IS-VALID-SEGMENT: true
}

func ExampleIsValidSegment_invalid() {

	var segment string = "wow!"

	result := ilpaddr.IsValidSegment(segment)

	fmt.Printf("SEGMENT: %q\n", segment)
	fmt.Printf("IS-VALID-SEGMENT: %t\n", result)

	// Output:
	// SEGMENT: "wow!"
	// IS-VALID-SEGMENT: false
}
