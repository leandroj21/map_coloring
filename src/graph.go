package src

import "fmt"

const amountOfColors int = 4

type Edge struct {
	Label string
}

type Node struct {
	// Node name
	Label string
	// 0 for not visited, -1 for already looked Nodes and 1 for visited
	Visited   int
	Neighbors []Edge
	Color     int
}

type Graph struct {
	Nodes []*Node
	// nil by default
	Indexes      map[string]int
	inited       bool
	nodesColored int
}

// InsertNode Inserts a node in the graph
func (graph *Graph) InsertNode(nodeLabel string, edgesLabels []string) {
	// Init Indexes map
	if !graph.inited {
		graph.Indexes = make(map[string]int)
		graph.inited = true
	}

	node := Node{Label: nodeLabel}
	for _, edgeLabel := range edgesLabels {
		edge := Edge{edgeLabel}
		node.Neighbors = append(node.Neighbors, edge)

		if idx, exists := graph.Indexes[edgeLabel]; exists {
			if _, isIn := SearchInNeighbors(nodeLabel, graph.Nodes[idx].Neighbors); !isIn {
				graph.Nodes[idx].Neighbors = append(graph.Nodes[idx].Neighbors, Edge{Label: nodeLabel})
			}
		}
	}

	graph.Nodes = append(graph.Nodes, &node)
	graph.Indexes[nodeLabel] = len(graph.Nodes) - 1
	return
}

// Prints the current state of the graph
func (graph *Graph) printState() {
	for index, node := range graph.Nodes {
		fmt.Printf("(%2d) %s --> %v\n", index, node.Label, node.Neighbors)
	}
}

func (graph *Graph) printTuples() {
	for _, node := range graph.Nodes {
		fmt.Printf("(%s, %d)%s", node.Label, node.Color, ", ")
	}
}

func (graph *Graph) isSafeToColor(Color, index int) bool {
	node := graph.Nodes[index]
	for _, neighbor := range node.Neighbors {
		if Color == graph.Nodes[graph.Indexes[neighbor.Label]].Color {
			return false
		}
	}

	return true
}

func (graph *Graph) getIndexUncoloredNeighbor(index int) int {
	for _, neighbor := range graph.Nodes[index].Neighbors {
		neighborIndex := graph.Indexes[neighbor.Label]
		if graph.Nodes[neighborIndex].Color == 0 {
			return neighborIndex
		}
	}
	return -1
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
