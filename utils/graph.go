package lemin

//add vertex
func (g *Graph) AddVeertex(k string) bool {
	if contains(g.Vertices, k) {
		return false
	} else {
		g.Vertices = append(g.Vertices, &Vertex{Key: k})
		return true
	}
}

//add edge
func (g *Graph) AddEdge(from, to string) bool {
	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)
	// check error
	if fromVertex == nil || toVertex == nil {
		return false
	} else if contains(fromVertex.Adjacent, to) {
		return true
	} else {
		fromVertex.Adjacent = append(fromVertex.Adjacent, toVertex)
		toVertex.Adjacent = append(toVertex.Adjacent, fromVertex)
		return true
	}
}

func (g *Graph) getVertex(k string) *Vertex {
	for i, v := range g.Vertices {
		if v.Key == k {
			return g.Vertices[i]
		}
	}
	return nil
}

//contains
func contains(s []*Vertex, k string) bool {
	for _, v := range s {
		if k == v.Key {
			return true
		}
	}
	return false
}

//////////////////////////////////////////////////////////////////
func (g *Graph) FillGraph(room []string, edges [][]string) bool {
	for _, l := range room {
		err := g.AddVeertex(l)
		if !err {
			return err
		}
	}

	for _, l := range edges {
		err := g.AddEdge(l[0], l[1])
		if !err {
			return err
		}
	}
	return true
}
