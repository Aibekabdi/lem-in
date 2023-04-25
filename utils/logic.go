package lemin

import (
	"fmt"
	"os"
)

func Logic(c Colony, text []string) {
	ants := c.Ants
	if len(Allpaths[0]) == 2 {
		fmt.Println(Allpaths[0])
		for _, v := range text {
			fmt.Println(v)
		}
		fmt.Println()
		for i := 1; i <= ants; i++ {
			fmt.Print("l", i, "-", Allpaths[0][1], " ")
		}
		fmt.Println()
		os.Exit(0)
	}
	GroupingDisjointPath()
	g := OptimalDisjointPath(ants)
	Send(g, c.Ants, c, text)

}

func OptimalDisjointPath(ants int) [][]string {
	var min int
	var vertexes = make(map[int][][]string)
	for _, p := range DisjointPath {

		paths := len(p)
		max := len(p[len(p)-1])
		var cur = 0
		for _, path := range p {
			cur += max - len(path)
		}
		left := ants - cur
		if left == 0 {
			continue
		}
		tops := max + left/paths
		if left%paths != 0 {
			tops++
		}

		vertexes[tops] = p
		min = tops
	}

	for h := range vertexes {
		if h < min {
			min = h
		}
	}
	return vertexes[min]
}
