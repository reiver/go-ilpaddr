package ilpaddr

import (
	"strings"
)

// HasScheme returns true of the ILP Address (ilpAddr) has the specified scheme.
func HasScheme(ilpAddr string, scheme string) bool {
	if "" == scheme {
		return false
	}
	if "" == ilpAddr {
		return false
	}

	var prefix string = scheme + separator
	return strings.HasPrefix(ilpAddr, prefix)
}
