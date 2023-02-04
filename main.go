package main

import (
	"coloring_map/src"
	"fmt"
	"time"
)

func testMapColoring(graph *src.Graph) {
	someError := false
	var fails uint
	for _, node := range graph.Nodes {
		for _, neighbor := range node.Neighbors {
			indexInGraph, exists := src.SearchInGraph(neighbor.Label, graph)
			if !exists {
				fmt.Println("ERROR IN TESTING: NON-EXIST NODE")
				someError = true
				fails++
				continue
			}

			if node.Color == 0 {
				fmt.Printf("NOT COLORED NODE: %s\n", node.Label)
				someError = true
				fails++
				continue
			}

			if graph.Nodes[indexInGraph].Color == node.Color {
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

func printTimeResults(parseTime, insertTime, coloringTime time.Duration) {
	fmt.Println("TIME RESULTS:")
	fmt.Println("\tParse time:", parseTime)
	fmt.Println("\tInsert nodes time:", insertTime)
	fmt.Println("\tColoring nodes time:", coloringTime)
}

func main() {
	maps := map[string]string{
		"Europe":        "sp pt fr an\nfr it be lu ch de an\nbe lu nl\nde be nl lu dk pl cz at ch li\nit ch si at\nat ch li si hu sk\nhr hu si ba rs me\nmk al bg el rs\nbg el ro tr\nrs me al ro hu ba\nba me\nsi hu\nsk hu cz ua pl\npl ru lt by ua cz\nua ro by ru md\nro mo hu\nfi se no ru\nno se ru\nlv ru lt ee bg\nru ee lt bg ge az\ntr am ge\nge az",
		"South America": "ar cl bo py br uy\nbo cl py pe br\npy br\nuy br\npe br ec co\nbr gy sr ve co fr\nve co\nec co",
	}

	var start time.Time
	var parseTime, insertTime, coloringTime time.Duration
	for mapName, countries := range maps {
		fmt.Printf("\n------%s------\n", mapName)
		start = time.Now()
		parsedEntry, order := src.ParseEntry(countries)
		parseTime = time.Since(start)

		var graph src.Graph
		start = time.Now()
		graph.InsertNodesInOrder(parsedEntry, order)
		insertTime = time.Since(start)

		start = time.Now()
		graph.Color()
		coloringTime = time.Since(start)

		testMapColoring(&graph)

		printTimeResults(parseTime, insertTime, coloringTime)

		fmt.Printf("\nCOLORING RESULT:\n")
		graph.PrintTuples()
	}
}
