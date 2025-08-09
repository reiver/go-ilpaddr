package ilpaddr

import (
	"strings"
)

// splitRoute splits a route into the first-sement, and the remaining route.
//
// For example:
//
//	var route string = "acmecorp.sales.199.~ipr.cdfa5e16-e759-4ba3-88f6-8b9dc83c1868.2"
//	
//	segment, rest := splitRoute(route)
//	// segment == "acmecorp"
//	// rest    == "sales.199.~ipr.cdfa5e16-e759-4ba3-88f6-8b9dc83c1868.2"
func splitRoute(route string) (firstSegment string, remainingRoute string) {
	if "" == route {
		return "", ""
	}

	index := strings.Index(route, separatorString)
	if index < 0 {
		return route, ""
	}

	return route[:index], route[index+len(separatorString):]
}
