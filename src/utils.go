package src

import "strings"

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

// ParseEntry Parses an entry like "sp fr\nfr it be", where sp is the node and the other
// two-chars code are its neighbors. Returns a slice that contains the original order of
// data.
func ParseEntry(entry string) (result map[string][]string, order []string) {
	result = make(map[string][]string)

	lines := strings.Split(entry, "\n")
	for _, line := range lines {
		data := strings.Split(line, " ")
		node := data[0]
		neighbors := data[1:]

		// Save correct position
		order = append(order, node)

		result[node] = neighbors
	}

	// Get missing nodes because of the way the entry is given
	justAdded := make([]string, 0)
	for node, neighbors := range result {
		for _, neighbor := range neighbors {
			// Map the neighbors that are not mapped
			if _, exists := result[neighbor]; !exists {
				justAdded = append(justAdded, neighbor)
				result[neighbor] = []string{node}

				order = append(order, neighbor)
			} else if _, recentlyAdded := SearchIn(neighbor, justAdded); exists && recentlyAdded {
				result[neighbor] = append(result[neighbor], node)
			}

			// Get missing links
			if _, exists := SearchIn(node, result[neighbor]); !exists {
				result[neighbor] = append(result[neighbor], node)
			}
		}
	}

	return
}
