package ilpaddr_test

import (
	"fmt"

	"github.com/reiver/go-ilpaddr"
)

func ExampleMatchLen_once() {
	var ilpAddr1 string = "g.apple.banana.cherry"
	var ilpAddr2 string = "g.us-fed.ach.0.acmebank.swx0a0.acmecorp.sales.199.~ipr.cdfa5e16-e759-4ba3-88f6-8b9dc83c1868.2"

	matchLen := ilpaddr.MatchLen(ilpAddr1, ilpAddr2)

	fmt.Printf("MATCH-LENGTH: %d\n", matchLen)

	// Output:
	// MATCH-LENGTH: 1
}

func ExampleMatchLen_twice() {
	var ilpAddr1 string = "g.apple.banana.cherry"
	var ilpAddr2 string = "g.apple.ach.0.acmebank.swx0a0.acmecorp.sales.199.~ipr.cdfa5e16-e759-4ba3-88f6-8b9dc83c1868.2"

	matchLen := ilpaddr.MatchLen(ilpAddr1, ilpAddr2)

	fmt.Printf("MATCH-LENGTH: %d\n", matchLen)

	// Output:
	// MATCH-LENGTH: 2
}

func ExampleMatchLen_thrice() {
	var ilpAddr1 string = "g.apple.banana.cherry"
	var ilpAddr2 string = "g.apple.banana.0.acmebank.swx0a0.acmecorp.sales.199.~ipr.cdfa5e16-e759-4ba3-88f6-8b9dc83c1868.2"

	matchLen := ilpaddr.MatchLen(ilpAddr1, ilpAddr2)

	fmt.Printf("MATCH-LENGTH: %d\n", matchLen)

	// Output:
	// MATCH-LENGTH: 3
}

func ExampleMatchLen_fource() {
	var ilpAddr1 string = "g.apple.banana.cherry"
	var ilpAddr2 string = "g.apple.banana.cherry.acmebank.swx0a0.acmecorp.sales.199.~ipr.cdfa5e16-e759-4ba3-88f6-8b9dc83c1868.2"

	matchLen := ilpaddr.MatchLen(ilpAddr1, ilpAddr2)

	fmt.Printf("MATCH-LENGTH: %d\n", matchLen)

	// Output:
	// MATCH-LENGTH: 4
}
