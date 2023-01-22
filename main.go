package main

import (
	"coloring_map/src"
	"fmt"
	"strings"
	"time"
)

func testColoringMap(graph *src.Graph) {
	someError := false
	var fails uint
	for _, node := range graph.Nodes {
		for _, neighbor := range node.Neighbors {
			if graph.Nodes[graph.Indexes[neighbor.Label]].Color == node.Color {
				fmt.Printf("%s and %s FAILED\n", node.Label, neighbor.Label)
				someError = true
				fails++
			}
		}
	}

	if someError {
		fmt.Println("Map coloring does not work. Fails:", fails)
	} else {
		fmt.Println("Map coloring works.")
	}
}

// Test if all the indexes saved in graph.indexes are correct
func testInsert(g *src.Graph) {
	for idxReal, value := range g.Nodes {
		if g.Indexes[value.Label] != idxReal {
			fmt.Printf("INDEX ERROR: %s (idx in map: %d) should be in %d\n",
				value.Label,
				g.Indexes[value.Label],
				idxReal,
			)
		}
	}
}

// Parses an entry like "sp fr\nfr it be", where sp is the node and the other
// two-chars code are its neighbors
func parseEntry(entry string) (result map[string][]string) {
	result = make(map[string][]string)

	lines := strings.Split(entry, "\n")
	for _, line := range lines {
		data := strings.Split(line, " ")
		node := data[0]
		neighbors := data[1:]
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
			} else if _, recentlyAdded := src.SearchIn(neighbor, justAdded); exists && recentlyAdded {
				result[neighbor] = append(result[neighbor], node)
			}

			// Get missing links
			if _, exists := src.SearchIn(node, result[neighbor]); !exists {
				result[neighbor] = append(result[neighbor], node)
			}
		}
	}

	return
}

func printTimeResults(parseTime, insertTime, coloringTime time.Duration) {
	fmt.Println("TIME RESULTS:")
	fmt.Println("\tParse time:", parseTime)
	fmt.Println("\tInsert nodes time:", insertTime)
	fmt.Println("\tColoring nodes time:", coloringTime)
}

func main() {
	//europeMap := "sp pt fr an\nfr it be lu ch de an\nbe lu nl\nde be nl lu dk pl cz at ch li\nit ch si at\nat ch li si hu sk\nhr hu si ba rs me\nmk al bg el rs\nbg el ro tr\nrs me al ro hu ba\nba me\nsi hu\nsk hu cz ua pl\npl ru lt by ua cz\nua ro by ru md\nro mo hu\nfi se no ru\nno se ru\nlv ru lt ee bg\nru ee lt bg ge az\ntr am ge\nge az"
	southAmericaMap := "ar cl bo py br uy\nbo cl py pe br\npy br\nuy br\npe br ec co\nbr gy sr ve co fr\nve co\nec co"

	var start time.Time
	var parseTime, insertTime, coloringTime time.Duration

	start = time.Now()
	parsedEntry := parseEntry(southAmericaMap)
	parseTime = time.Since(start)

	var graph src.Graph
	start = time.Now()
	for node, neighbors := range parsedEntry {
		graph.InsertNode(node, neighbors)
	}
	insertTime = time.Since(start)

	testInsert(&graph)

	start = time.Now()
	graph.Color()
	coloringTime = time.Since(start)

	testColoringMap(&graph)

	printTimeResults(parseTime, insertTime, coloringTime)
}
