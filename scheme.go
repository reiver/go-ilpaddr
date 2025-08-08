package ilpaddr

import (
	"strings"
)

// Scheme returns the scheme of an ILP Address.
//
// If the ILP Address does not have a scheme,
// or there if some other problem, it returns an empty string.
func Scheme(ilpAddr string) string {
	if "" == ilpAddr {
		return ""
	}

	index := strings.Index(ilpAddr, separator)
	if index < 0 {
		return ""
	}

	// Return "" for inputs such as "g.", "local.", etc.
	if index == len(ilpAddr) - 1 {
		return ""
	}

	return ilpAddr[:index]
}
