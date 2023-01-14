package main

import (
	"fmt"
	"strings"
)

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
			if _, exists := result[neighbor]; !exists {
				justAdded = append(justAdded, neighbor)
				result[neighbor] = []string{node}
			} else if _, recentlyAdded := searchIn(neighbor, justAdded); exists && recentlyAdded {
				result[neighbor] = append(result[neighbor], node)
			}
		}
	}
	return
}

func main() {
	europeMap := "sp pt fr an\nfr it be lu ch de an\nbe lu nl\nde be nl lu dk pl cz at ch li\nit ch si at\nat ch li si hu sk\nhr hu si ba rs me\nmk al bg el rs\nbg el ro tr\nrs me al ro hu ba\nba me\nsi hu\nsk hu cz ua pl\npl ru lt by ua cz\nua ro by ru md\nro mo hu\nfi se no ru\nno se ru\nlv ru lt ee bg\nru ee lt bg ge az\ntr am ge\nge az"
	parsedEntry := parseEntry(europeMap)

	var graph Graph
	for node, neighbors := range parsedEntry {
		graph.insertNode(node, neighbors)
	}
	graph.print()
}
