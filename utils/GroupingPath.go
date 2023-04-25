package lemin

var DisjointPath = [][][]string{}

func GroupingDisjointPath() {
	for index, path := range Allpaths {
		Allpaths[index] = path[1:]
	}
	for index, path := range Allpaths {
		var group = [][]string{}
		group = append(group, path)
		DisjointPath = append(DisjointPath, group)
		for index2 := index + 1; index2 < len(Allpaths); index2++ {
			path2 := Allpaths[index2]
			if !DisjointCheck(group, path2) {
				group = append(group, path2)
			}
		}
		DisjointPath = append(DisjointPath, group)
	}
	Remove()
}

func DisjointCheck(group [][]string, p []string) bool {
	for i, r := range p { // r = rooms
		for _, path := range group {
			for _, r2 := range path {
				if r == r2 && i != len(p)-1 {
					return true
				}
			}
		}
	}
	return false
}
