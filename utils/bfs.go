package lemin

var Allpaths = [][]string{}

var visited = make(map[string]bool)
var queue = []string{}

func (graph *Graph) BFS(c Colony, room string) {
	if visited[room] {
		return
	}
	visited[room] = true
	queue = append(queue, room)
	if room != c.End {
		for _, v := range graph.Vertices {
			if v.Key == room {
				for _, adj := range v.Adjacent {
					graph.BFS(c, adj.Key)
				}
			}
		}
	} else {
		var currentPath = make([]string, len(queue))
		for index, pathVertice := range queue {
			currentPath[index] = pathVertice
		}
		Allpaths = append(Allpaths, currentPath)
	}
	queue = queue[:len(queue)-1]
	visited[room] = false
}
