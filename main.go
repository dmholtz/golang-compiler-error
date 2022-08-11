package main

import "fmt"

type WeightType interface {
	float64 | int
}

type IEdge interface {
	To() int
}

type Edge struct {
	to int
}

func (e Edge) To() int {
	return e.to
}

type WeightedEdge[W WeightType] struct {
	Edge
	weight W
}

func (e WeightedEdge[W]) Weight() W {
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
	g := Graph[struct{}, WeightedEdge[int]]{}
	Accept(g)
}
