package main

import (
	"fmt"
	"strings"
	"time"
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
	indexes      map[string]int
	inited       bool
	nodesColored int
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
func (graph *Graph) printState() {
	for index, node := range graph.nodes {
		fmt.Printf("(%2d) %s --> %v\n", index, node.label, node.neighbors)
	}
}

func (graph *Graph) printTuples() {
	for _, node := range graph.nodes {
		fmt.Printf("(%s, %d)%s", node.label, node.color, ", ")
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

func (graph *Graph) getIndexUncoloredNeighbor(index int) int {
	for _, neighbor := range graph.nodes[index].neighbors {
		neighborIndex := graph.indexes[neighbor.label]
		if graph.nodes[neighborIndex].color == 0 {
			return neighborIndex
		}
	}
	return -1
}

// Colors every node, but a node and any of its neighbors cannot have
// the same color
func (graph *Graph) color(index int) bool {
	if graph.nodesColored >= len(graph.nodes) {
		return true
	}

	for c := 1; c <= amountOfColors; c++ {
		if !graph.isSafeToColor(c, index) {
			continue
		}

		graph.nodes[index].color = c
		graph.nodesColored++

		nextIdx := index + 1
		if graph.color(nextIdx) {
			return true
		}

		graph.nodes[index].color = 0
		graph.nodesColored--
	}

	return false
}

func testColoringMap(graph *Graph) {
	someError := false
	var fails uint
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

	var graph Graph
	start = time.Now()
	for node, neighbors := range parsedEntry {
		graph.insertNode(node, neighbors)
	}
	insertTime = time.Since(start)

	testInsert(&graph)

	start = time.Now()
	graph.color(0)
	coloringTime = time.Since(start)

	testColoringMap(&graph)

	printTimeResults(parseTime, insertTime, coloringTime)
}
