package lemin

func Remove() {
	var curGroup [][][]string
	for i, group := range DisjointPath {
		if len(group) == 1 && i != 0 {
			continue
		}
		if !isSimilar(curGroup, group) {
			curGroup = append(curGroup, group)
		}
	}
	DisjointPath = curGroup
}

func isSimilar(newgroup [][][]string, group2 [][]string) bool {
	for _, group1 := range newgroup {
		if len(group1) == len(group2) {
			var isOk = make([]bool, len(group2))
			for i, path := range group2 {
				groupOfPath := group1[i]
				if len(path) >= len(groupOfPath) {
					isOk = append(isOk, true)
				} else {
					isOk = append(isOk, false)
				}
			}
			counter := 0
			for _, value := range isOk {
				if value {
					counter++
				}
			}
			if counter == len(group2) {
				return true
			}
		}
	}
	return false
}
