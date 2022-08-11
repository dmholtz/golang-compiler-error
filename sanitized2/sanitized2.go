package sanitized2

import "fmt"

type IEdge interface {
	To() int
}

type Edge struct {
	to int
}

func (e Edge) To() int {
	return e.to
}

type WeightedEdge struct {
	Edge
	weight int
}

func (e WeightedEdge) Weight() int {
	return e.weight
}

type Graph[N any, E IEdge] struct {
	nodes []N
	edges []E
}

func Accept[N any, E IEdge](gr Graph[N, E]) {
	for _, e := range gr.edges {
		to := e.To()
		fmt.Printf("To: %d\n", to)
	}
}

func main() {
	g := Graph[struct{}, WeightedEdge]{}
	Accept(g)
}
