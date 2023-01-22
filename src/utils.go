package src

// SearchIn Searches a given string [str] in a slice of strings
// Returns a tuple of the index and the boolean value of the existence
func SearchIn(str string, slice []string) (int, bool) {
	for idx, value := range slice {
		if value == str {
			return idx, true
		}
	}

	return -1, false
}

func SearchInNeighbors(str string, slice []Edge) (int, bool) {
	for idx, value := range slice {
		if value.Label == str {
			return idx, true
		}
	}

	return -1, false
}
