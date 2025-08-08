package ilpaddr_test

import (
	"fmt"

	"github.com/reiver/go-ilpaddr"
)

func ExampleIsValidScheme_schemeUnknown() {
	var scheme string = "unknown"

	result := ilpaddr.IsValidScheme(scheme)

	fmt.Printf("IS-VALID-SCHEME: %t\n", result)

	// Output:
	// IS-VALID-SCHEME: false
}

func ExampleIsValidScheme_schemeGlobal() {
	var scheme string = "g"

	result := ilpaddr.IsValidScheme(scheme)

	fmt.Printf("IS-VALID-SCHEME: %t\n", result)

	// Output:
	// IS-VALID-SCHEME: true
}

func ExampleIsValidScheme_schemePrivate() {
	var scheme string = "private"

	result := ilpaddr.IsValidScheme(scheme)

	fmt.Printf("IS-VALID-SCHEME: %t\n", result)

	// Output:
	// IS-VALID-SCHEME: true
}

func ExampleIsValidScheme_schemeExample() {
	var scheme string = "example"

	result := ilpaddr.IsValidScheme(scheme)

	fmt.Printf("IS-VALID-SCHEME: %t\n", result)

	// Output:
	// IS-VALID-SCHEME: true
}

func ExampleIsValidScheme_schemeTest() {
	var scheme string = "test"

	result := ilpaddr.IsValidScheme(scheme)

	fmt.Printf("IS-VALID-SCHEME: %t\n", result)

	// Output:
	// IS-VALID-SCHEME: true
}

func ExampleIsValidScheme_schemeTest1() {
	var scheme string = "test1"

	result := ilpaddr.IsValidScheme(scheme)

	fmt.Printf("IS-VALID-SCHEME: %t\n", result)

	// Output:
	// IS-VALID-SCHEME: true
}

func ExampleIsValidScheme_schemeTest2() {
	var scheme string = "test2"

	result := ilpaddr.IsValidScheme(scheme)

	fmt.Printf("IS-VALID-SCHEME: %t\n", result)

	// Output:
	// IS-VALID-SCHEME: true
}

func ExampleIsValidScheme_schemeTest3() {
	var scheme string = "test3"

	result := ilpaddr.IsValidScheme(scheme)

	fmt.Printf("IS-VALID-SCHEME: %t\n", result)

	// Output:
	// IS-VALID-SCHEME: true
}

func ExampleIsValidScheme_schemeLocal() {
	var scheme string = "local"

	result := ilpaddr.IsValidScheme(scheme)

	fmt.Printf("IS-VALID-SCHEME: %t\n", result)

	// Output:
	// IS-VALID-SCHEME: true
}

func ExampleIsValidScheme_schemePeer() {
	var scheme string = "peer"

	result := ilpaddr.IsValidScheme(scheme)

	fmt.Printf("IS-VALID-SCHEME: %t\n", result)

	// Output:
	// IS-VALID-SCHEME: true
}

func ExampleIsValidScheme_schemeSelf() {
	var scheme string = "self"

	result := ilpaddr.IsValidScheme(scheme)

	fmt.Printf("IS-VALID-SCHEME: %t\n", result)

	// Output:
	// IS-VALID-SCHEME: true
}
