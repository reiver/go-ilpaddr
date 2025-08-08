package ilpaddr_test

import (
	"fmt"

	"github.com/reiver/go-ilpaddr"
)

func ExampleHasScheme_schemeGlobal() {
	var ilpAddr string = "g.once.twice.thrice.fource"

	result := ilpaddr.HasScheme(ilpAddr, ilpaddr.SchemeGlobal)

	fmt.Printf("HAS-SCHEME: %t\n", result)

	// Output:
	// HAS-SCHEME: true
}

func ExampleHasScheme_schemePrivate() {
	var ilpAddr string = "private.once.twice.thrice.fource"

	result := ilpaddr.HasScheme(ilpAddr, ilpaddr.SchemePrivate)

	fmt.Printf("HAS-SCHEME: %t\n", result)

	// Output:
	// HAS-SCHEME: true
}

func ExampleHasScheme_schemeExample() {
	var ilpAddr string = "example.once.twice.thrice.fource"

	result := ilpaddr.HasScheme(ilpAddr, ilpaddr.SchemeExample)

	fmt.Printf("HAS-SCHEME: %t\n", result)

	// Output:
	// HAS-SCHEME: true
}

func ExampleHasScheme_schemeTest() {
	var ilpAddr string = "test.once.twice.thrice.fource"

	result := ilpaddr.HasScheme(ilpAddr, ilpaddr.SchemeTest)

	fmt.Printf("HAS-SCHEME: %t\n", result)

	// Output:
	// HAS-SCHEME: true
}

func ExampleHasScheme_schemeTest1() {
	var ilpAddr string = "test1.once.twice.thrice.fource"

	result := ilpaddr.HasScheme(ilpAddr, ilpaddr.SchemeTest1)

	fmt.Printf("HAS-SCHEME: %t\n", result)

	// Output:
	// HAS-SCHEME: true
}

func ExampleHasScheme_schemeTest2() {
	var ilpAddr string = "test2.once.twice.thrice.fource"

	result := ilpaddr.HasScheme(ilpAddr, ilpaddr.SchemeTest2)

	fmt.Printf("HAS-SCHEME: %t\n", result)

	// Output:
	// HAS-SCHEME: true
}

func ExampleHasScheme_schemeTest3() {
	var ilpAddr string = "test3.once.twice.thrice.fource"

	result := ilpaddr.HasScheme(ilpAddr, ilpaddr.SchemeTest3)

	fmt.Printf("HAS-SCHEME: %t\n", result)

	// Output:
	// HAS-SCHEME: true
}

func ExampleHasScheme_schemeLocal() {
	var ilpAddr string = "local.once.twice.thrice.fource"

	result := ilpaddr.HasScheme(ilpAddr, ilpaddr.SchemeLocal)

	fmt.Printf("HAS-SCHEME: %t\n", result)

	// Output:
	// HAS-SCHEME: true
}

func ExampleHasScheme_schemePeer() {
	var ilpAddr string = "peer.once.twice.thrice.fource"

	result := ilpaddr.HasScheme(ilpAddr, ilpaddr.SchemePeer)

	fmt.Printf("HAS-SCHEME: %t\n", result)

	// Output:
	// HAS-SCHEME: true
}

func ExampleHasScheme_schemeSelf() {
	var ilpAddr string = "self.once.twice.thrice.fource"

	result := ilpaddr.HasScheme(ilpAddr, ilpaddr.SchemeSelf)

	fmt.Printf("HAS-SCHEME: %t\n", result)

	// Output:
	// HAS-SCHEME: true
}
