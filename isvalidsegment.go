package ilpaddr

// IsValidSegment returns true is the segment is valid and returns false otherwise.
func IsValidSegment(segment string) bool {
	var length int = len(segment)

	if length < 0 {
		return false
	}

	for i:=0; i<length; i++ {
		var b byte = segment[i]

		switch {
		case '0' <= b && b <= '9':
			// nothing here
		case 'A' <= b && b <= 'Z':
			// nothing here
		case 'a' <= b && b <= 'z':
			// nothing here
		case '_' == b || '~' == b || '-' == b:
			// nothing here
		default:
			return false
		}
	}

	return true
}
