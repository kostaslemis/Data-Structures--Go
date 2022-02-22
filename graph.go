package main

import "fmt"

// Graph represents an adjacency list graph
type Graph struct {
	vertices []*Vertex
}


// Vertex represents a graph vertex
type Vertex struct {
	key int
	adjacent []*Vertex
}

// AddVertex adds a Vertex to the Graph
func (graph *Graph) AddVertex(key int) {
	if contains(graph.vertices, key) {
		err := fmt.Errorf("vertex %v not added because it is an existing key", key)
		fmt.Println(err.Error())
	} else {
		graph.vertices = append(graph.vertices, &Vertex{key: key})
	}
}

// AddEdge adds an edge to the graph
func (graph *Graph) AddEdge(from,to int) {
	// get vertex
	fromVertex := graph.getVertex(from)
	toVertex := graph.getVertex(to)
	// check error
	if fromVertex == nil || toVertex == nil {
		err := fmt.Errorf("invalid edge (%v-->%v)", from, to)
		fmt.Println(err.Error())
	} else if contains(fromVertex.adjacent, to)	{
		err := fmt.Errorf("existing edge ((%v-->%v)", from, to)
		fmt.Println(err.Error())
	} else {
		// add edge
		fromVertex.adjacent = append(fromVertex.adjacent, toVertex)
	}
}


// getVertex returns a pointer to the Vertex with a key integer
func (graph *Graph) getVertex(key int) *Vertex {
	for i, vertex := range graph.vertices {
		if vertex.key == key {
			return graph.vertices[i]
		}
	}
	return nil
}

// contains
func contains(s []*Vertex, key int) bool {
	for _,vertex := range s {
		if key == vertex.key {
			return true
		}
	}
	return false
}


// Print will print the adjacent list for each vertex of the graph
func (graph *Graph) Print() {
	for _,vertex := range graph.vertices {
		fmt.Printf("\nVertex %v : ", vertex.key)
		for _,vertex := range vertex.adjacent {
			fmt.Printf(" %v", vertex.key)
		}
	}
}


func main() {
	graph := &Graph{}

	for i := 0; i < 5; i++ {
		graph.AddVertex(i)
	}

	graph.AddEdge(0,2)
	graph.AddEdge(6,1)
	graph.AddEdge(2,4)
	graph.AddEdge(2,4)

	graph.Print()
}
