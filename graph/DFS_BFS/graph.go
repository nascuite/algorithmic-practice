package main

import (
	"fmt"
	"strings"
)

func main() {
	g := NewGraph()

	g.AddVertex("ajinkya")
	g.AddVertex("francesc")
	g.AddVertex("manish")
	g.AddVertex("albert")

	g.AddEdge("albert", "ajinkya")
	g.AddEdge("ajinkya", "albert")
	g.AddEdge("francesc", "ajinkya")
	g.AddEdge("francesc", "manish")
	g.AddEdge("manish", "francesc")
	g.AddEdge("manish", "albert")

	g.DFS("francesc")
	g.CreatePath("francesc", "albert")
}

func NewGraph() Graph {
	return Graph{
		adjacency: make(map[string][]string),
	}
}

type Graph struct {
	adjacency map[string][]string
}

func (g *Graph) AddVertex(vertex string) bool {
	if _, ok := g.adjacency[vertex]; ok {
		fmt.Println("vertex %v already exists", vertex)
		return false
	}
	g.adjacency[vertex] = []string{}
	return true
}

func (g *Graph) AddEdge(vertex, neighbor string) bool {
	if _, ok := g.adjacency[vertex]; !ok {
		fmt.Println("vertex $v does not exists", vertex)
		return false
	}
	if ok := contains(g.adjacency[vertex], neighbor); ok {
		fmt.Println("neighbor %v already exists for vertex %v", neighbor, vertex)
		return false
	}

	g.adjacency[vertex] = append(g.adjacency[vertex], neighbor)
	return true
}

func contains(slice []string, item string) bool {
	for _, substring := range slice {
		if substring == item {
			return true
		}
	}
	return false
}

func (g Graph) DFS(startVertex string) {
	visited := g.createVisited()
	g.dfsRecursive(startVertex, visited)

}

func (g Graph) createVisited() map[string]bool {
	visited := make(map[string]bool, len(g.adjacency))
	for key := range g.adjacency {
		visited[key] = false
	}
	return visited
}

func (g Graph) dfsRecursive(startVertex string, visited map[string]bool) {
	visited[startVertex] = true
	fmt.Println("DFS", startVertex)
	for _, neighbor := range g.adjacency[startVertex] {
		if !visited[neighbor] {
			g.dfsRecursive(neighbor, visited)
		}
	}
}

func (g Graph) CreatePath(firstVertex, secondVertex string) bool {
	visited := g.createVisited()
	var path, queue []string

	queue = append(queue, firstVertex)
	visited[firstVertex] = true

	for len(queue) > 0 {
		currentVertex := queue[0]
		queue = queue[1:]

		path = append(path, currentVertex)
		edges := g.adjacency[currentVertex]
		if contains(edges, secondVertex) {
			path = append(path, secondVertex)
			fmt.Println(strings.Join(path, "->"))
			return true
		}

		for _, neighbor := range g.adjacency[currentVertex] {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
	}
	fmt.Println("no path found")
	return false
}
