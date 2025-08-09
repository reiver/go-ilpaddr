package ilpaddr

// Scheme returns the scheme of an ILP Address.
//
// If the ILP Address does not have a scheme,
// or there if some other problem, it returns an empty string.
func Scheme(ilpAddr string) string {
	if "" == ilpAddr {
		return ""
	}

	scheme, remainingRoute := FirstRest(ilpAddr)
	if "" == remainingRoute {
		return ""
	}

	// Return "" for inputs such as "g.", "local.", etc.
	if len(scheme) == len(ilpAddr) - 1 {
		return ""
	}

	return scheme
}
