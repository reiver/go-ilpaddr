package ilpaddr_test

import (
	"fmt"

	"github.com/reiver/go-ilpaddr"
)

func ExampleScheme_schemeGlobal() {
	var ilpAddr string = "g.once.twice.thrice.fource"

	scheme := ilpaddr.Scheme(ilpAddr)

	fmt.Printf("SCHEME: %q\n", scheme)

	// Output:
	// SCHEME: "g"
}

func ExampleScheme_schemePrivate() {
	var ilpAddr string = "private.once.twice.thrice.fource"

	scheme := ilpaddr.Scheme(ilpAddr)

	fmt.Printf("SCHEME: %q\n", scheme)

	// Output:
	// SCHEME: "private"
}

func ExampleScheme_schemeExample() {
	var ilpAddr string = "example.once.twice.thrice.fource"

	 scheme := ilpaddr.Scheme(ilpAddr)

	fmt.Printf("SCHEME: %q\n", scheme)

	// Output:
	// SCHEME: "example"
}

func ExampleScheme_schemeTest() {
	var ilpAddr string = "test.once.twice.thrice.fource"

	 scheme := ilpaddr.Scheme(ilpAddr)

	fmt.Printf("SCHEME: %q\n", scheme)

	// Output:
	// SCHEME: "test"
}

func ExampleScheme_schemeTest1() {
	var ilpAddr string = "test1.once.twice.thrice.fource"

	 scheme := ilpaddr.Scheme(ilpAddr)

	fmt.Printf("SCHEME: %q\n", scheme)

	// Output:
	// SCHEME: "test1"
}

func ExampleScheme_schemeTest2() {
	var ilpAddr string = "test2.once.twice.thrice.fource"

	 scheme := ilpaddr.Scheme(ilpAddr)

	fmt.Printf("SCHEME: %q\n", scheme)

	// Output:
	// SCHEME: "test2"
}

func ExampleScheme_schemeTest3() {
	var ilpAddr string = "test3.once.twice.thrice.fource"

	 scheme := ilpaddr.Scheme(ilpAddr)

	fmt.Printf("SCHEME: %q\n", scheme)

	// Output:
	// SCHEME: "test3"
}

func ExampleScheme_schemeLocal() {
	var ilpAddr string = "local.once.twice.thrice.fource"

	 scheme := ilpaddr.Scheme(ilpAddr)

	fmt.Printf("SCHEME: %q\n", scheme)

	// Output:
	// SCHEME: "local"
}

func ExampleScheme_schemePeer() {
	var ilpAddr string = "peer.once.twice.thrice.fource"

	 scheme := ilpaddr.Scheme(ilpAddr)

	fmt.Printf("SCHEME: %q\n", scheme)

	// Output:
	// SCHEME: "peer"
}

func ExampleScheme_schemeSelf() {
	var ilpAddr string = "self.once.twice.thrice.fource"

	 scheme := ilpaddr.Scheme(ilpAddr)

	fmt.Printf("SCHEME: %q\n", scheme)

	// Output:
	// SCHEME: "self"
}
