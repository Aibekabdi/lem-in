package main

import (
	"fmt"
	lemin "lemin/utils"
	"os"
)

func main() {
	// -------------------- working with file--------------------------------//
	arg := os.Args[1:]
	if len(arg) != 1 {
		fmt.Println("non number of argumets")
		os.Exit(0)
	}
	text, errReader := lemin.Reader(arg[0])
	if errReader != nil {
		fmt.Println("ERROR: invalid data format")
		os.Exit(0)
	}
	rooms, edges, colony := lemin.TakeInfo(text)
	// ---------------------- end --------------------------------------------//

	// ---------------------- graph -----------------------------------------//
	graph := &lemin.Graph{}
	err := graph.FillGraph(rooms, edges)
	if !err {
		fmt.Println("ERROR: invalid data format")
		os.Exit(0)
	}
	// ---------------------- end --------------------------------------------//
	graph.BFS(colony, colony.Start)
	lemin.Checker(colony)
	lemin.QuickSort(lemin.Allpaths, 0, len(lemin.Allpaths)-1)
	lemin.Logic(colony, text)
}
