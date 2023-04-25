package lemin

type Graph struct {
	Vertices []*Vertex
}
type Vertex struct {
	Key      string
	Adjacent []*Vertex
}

type Colony struct {
	Start string
	End   string
	Ants  int
}
