package ilpaddr

// MatchLen returns the "prefix match" length between two ILP Addresses.
func MatchLen(ilpAddr1 string, ilpAddr2 string) uint {
	if "" == ilpAddr1 {
		return 0
	}
	if "" == ilpAddr2 {
		return 0
	}
	if separatorByte == ilpAddr1[len(ilpAddr1)-1] {
		return 0
	}
	if separatorByte == ilpAddr2[len(ilpAddr2)-1] {
		return 0
	}

	var matchLen uint

	var route1 string
	var route2 string
	{
		var scheme1 string
		var scheme2 string

		scheme1, route1 = FirstRest(ilpAddr1)
		scheme2, route2 = FirstRest(ilpAddr2)

		if scheme1 != scheme2 {
			return 0
		}
		if !IsValidScheme(scheme1) {
			return 0
		}
		if "" == route1 {
			return 0
		}
		if "" == route2 {
			return 0
		}
	}

	matchLen++

	for {
		var segment1 string
		var segment2 string

		segment1, route1 = FirstRest(route1)
		segment2, route2 = FirstRest(route2)

		if segment1 != segment2 {
			break
		}
		if !IsValidSegment(segment1) {
			return 0
		}

		matchLen++

		if "" == route1 {
			break
		}
		if "" == route2 {
			break
		}

	}

	return matchLen
}
