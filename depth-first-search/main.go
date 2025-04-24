package main

import "fmt"

func DFS(node int, graph map[int][]int, visited map[int]bool) {
	if visited[node] {
		return
	}

	visited[node] = true
	fmt.Println("Visited:", node)

	for _, neighbor := range graph[node] {
		DFS(neighbor, graph, visited)
	}
}

func main() {
	graph := map[int][]int{
		1: {2, 3},
		2: {1, 4},
		3: {1, 5},
		4: {2, 5},
		5: {3, 4},
	}

	visited := make(map[int]bool)
	DFS(1, graph, visited)
}
