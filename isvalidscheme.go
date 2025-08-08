package ilpaddr

// IsValidScheme returns true is the scheme is valid and returns false otherwise.
func IsValidScheme(scheme string) bool {
	switch scheme {
	case SchemeGlobal, SchemePrivate, SchemeExample, SchemeTest, SchemeTest1, SchemeTest2, SchemeTest3, SchemeLocal, SchemePeer, SchemeSelf:
		return true
	default:
		return false
	}
}
