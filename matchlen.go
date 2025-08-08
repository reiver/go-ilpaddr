package ilpaddr

// MatchLen returns the "prefix match" length between two ILP Addresses.
func MatchLen(ilpAddr1 string, ilpAddr2 string) uint {
	if "" == ilpAddr1 {
		return 0
	}
	if "" == ilpAddr2 {
		return 0
	}

	var shortest, longest string = ilpAddr1, ilpAddr2
	if len(ilpAddr2) < len(ilpAddr1) {
		shortest, longest = ilpAddr2, ilpAddr1
	}

	var matchLen uint

	{
		var limit int = len(shortest)

		for i:=0; i<limit; i++ {
			var b1 byte = shortest[i]
			var b2 byte = longest[i]

			if b1 != b2 {
				return matchLen
			}

			if separatorByte == b1 {
				matchLen++
			}
		}
	}

	return matchLen+1
}
