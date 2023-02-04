package src

import "fmt"

const amountOfColors int = 4
const tuplesInOneLine int = 5

type Edge struct {
	Label string
}

type Node struct {
	// Node name
	Label     string
	Neighbors []Edge
	Color     int
}

type Graph struct {
	Nodes        []*Node
	nodesColored int
}

// InsertNode Inserts a node in the graph
func (graph *Graph) InsertNode(nodeLabel string, edgesLabels []string) {
	node := Node{Label: nodeLabel}
	for _, edgeLabel := range edgesLabels {
		edge := Edge{edgeLabel}
		node.Neighbors = append(node.Neighbors, edge)
	}

	graph.Nodes = append(graph.Nodes, &node)
	return
}

// InsertNodesInOrder inserts nodes in a specific order into the node's array
func (graph *Graph) InsertNodesInOrder(data map[string][]string, order []string) {
	for _, nodeLabel := range order {
		graph.InsertNode(nodeLabel, data[nodeLabel])
	}
}

// PrintState Prints the current state of the graph
func (graph *Graph) PrintState() {
	for index, node := range graph.Nodes {
		fmt.Printf("(%2d) %s --> %v\n", index, node.Label, node.Neighbors)
	}
}

// PrintTuples prints the graph in tuples, such as ([node], [node color])
func (graph *Graph) PrintTuples() {
	lineCount := 0
	for _, node := range graph.Nodes {
		fmt.Printf("(%s, %d)%s", node.Label, node.Color, ", ")
		lineCount++

		if lineCount >= tuplesInOneLine {
			fmt.Println("")
			lineCount = 0
		}
	}
}

func (graph *Graph) isSafeToColor(Color, index int) bool {
	node := graph.Nodes[index]
	for _, neighbor := range node.Neighbors {
		indexInGraph, exists := SearchInGraph(neighbor.Label, graph)
		if !exists {
			continue
		}

		if Color == graph.Nodes[indexInGraph].Color {
			return false
		}
	}

	return true
}

// Colors every node, but a node and any of its Neighbors cannot have
// the same color
func (graph *Graph) coloringGraph(index int) bool {
	if graph.nodesColored >= len(graph.Nodes) {
		return true
	}

	for c := 1; c <= amountOfColors; c++ {
		if !graph.isSafeToColor(c, index) {
			continue
		}

		graph.Nodes[index].Color = c
		graph.nodesColored++

		nextIdx := index + 1
		if graph.coloringGraph(nextIdx) {
			return true
		}

		graph.Nodes[index].Color = 0
		graph.nodesColored--
	}

	return false
}

func (graph *Graph) Color() {
	graph.coloringGraph(0)
}
