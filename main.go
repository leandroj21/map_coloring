package main

import (
	"fmt"
	"strings"
)

const amountOfColors int = 4

type Edge struct {
	label string
}

type Node struct {
	// Node name
	label string
	// 0 for not visited, -1 for already looked nodes and 1 for visited
	visited           int
	previousNodeIndex int
	neighbors         []Edge
	color             int
}

type Graph struct {
	nodes []*Node
	// nil by default
	indexes map[string]int
	inited  bool
}

// Inserts a node in the graph
func (graph *Graph) insertNode(nodeLabel string, edgesLabels []string) {
	// Init indexes map
	if !graph.inited {
		graph.indexes = make(map[string]int)
		graph.inited = true
	}

	node := Node{label: nodeLabel, previousNodeIndex: -1}
	for _, edgeLabel := range edgesLabels {
		edge := Edge{edgeLabel}
		node.neighbors = append(node.neighbors, edge)

		if idx, exists := graph.indexes[edgeLabel]; exists {
			if _, isIn := searchInNeighbors(nodeLabel, graph.nodes[idx].neighbors); !isIn {
				graph.nodes[idx].neighbors = append(graph.nodes[idx].neighbors, Edge{label: nodeLabel})
			}
		}
	}

	graph.nodes = append(graph.nodes, &node)
	graph.indexes[nodeLabel] = len(graph.nodes) - 1
	return
}

// Prints the current state of the graph
func (graph *Graph) print() {
	for index, node := range graph.nodes {
		fmt.Printf("(%2d) %s --> %v\n", index, node.label, node.neighbors)
	}
}

// Searches a given string [str] in a slice of strings
// Returns a tuple of the index and the boolean value of the existence
func searchIn(str string, slice []string) (int, bool) {
	for idx, value := range slice {
		if value == str {
			return idx, true
		}
	}

	return -1, false
}

func searchInNeighbors(str string, slice []Edge) (int, bool) {
	for idx, value := range slice {
		if value.label == str {
			return idx, true
		}
	}

	return -1, false
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
			} else if _, recentlyAdded := searchIn(neighbor, justAdded); exists && recentlyAdded {
				result[neighbor] = append(result[neighbor], node)
			}

			// Get missing links
			if _, exists := searchIn(node, result[neighbor]); !exists {
				result[neighbor] = append(result[neighbor], node)
			}
		}
	}

	return
}

func (graph *Graph) isSafeToColor(color, index int) bool {
	node := graph.nodes[index]
	for _, neighbor := range node.neighbors {
		if color == graph.nodes[graph.indexes[neighbor.label]].color {
			return false
		}
	}

	return true
}

func getNextColor(color int) (next int) {
	next = color + 1
	if next > amountOfColors {
		next = 1
	}
	return
}

// Colors every node, but a node and any of its neighbors cannot have
// the same color
func (graph *Graph) color() {
	currentColor := 1
	for nodeIndex, node := range graph.nodes {
		if node.color != 0 {
			continue
		}

		for {
			if graph.isSafeToColor(currentColor, nodeIndex) {
				graph.nodes[nodeIndex].color = currentColor
				break
			}
			currentColor = getNextColor(currentColor)
		}
	}
}

func testColoringMap(graph *Graph) {
	someError := false
	var fails uint
	fmt.Println("\nTESTING MAP COLORING...")
	for _, node := range graph.nodes {
		for _, neighbor := range node.neighbors {
			if graph.nodes[graph.indexes[neighbor.label]].color == node.color {
				fmt.Printf("%s and %s FAILED\n", node.label, neighbor.label)
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
func testInsert(g *Graph) {
	for idxReal, value := range g.nodes {
		if g.indexes[value.label] != idxReal {
			fmt.Printf("INDEX ERROR: %s (idx in map: %d) should be in %d\n", value.label, g.indexes[value.label], idxReal)
		}
	}
}

func main() {
	europeMap := "sp pt fr an\nfr it be lu ch de an\nbe lu nl\nde be nl lu dk pl cz at ch li\nit ch si at\nat ch li si hu sk\nhr hu si ba rs me\nmk al bg el rs\nbg el ro tr\nrs me al ro hu ba\nba me\nsi hu\nsk hu cz ua pl\npl ru lt by ua cz\nua ro by ru md\nro mo hu\nfi se no ru\nno se ru\nlv ru lt ee bg\nru ee lt bg ge az\ntr am ge\nge az"

	parsedEntry := parseEntry(europeMap)
	var graph Graph
	for node, neighbors := range parsedEntry {
		graph.insertNode(node, neighbors)
	}
	testInsert(&graph)
	graph.color()
	testColoringMap(&graph)
}
