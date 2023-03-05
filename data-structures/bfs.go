package main

import "fmt"

func (g *Graph2) BFS(node int) {
	visited := make(map[int]bool, 0)
	
	for _, v := range g.vertices {
		visited[v.value] = false
	}
	
	visited[node] = true
	
	queue := &Queue{}
	queue.Enqueue(node)
	fmt.Println("initial queue is", (*queue))
	
	for !queue.IsEmpty() {
		front := (*queue)[0]
		frontAdj := g.GetVertex(front)
		for _, v := range frontAdj.adjacent {
			if !visited[front] {
				visited[front] = true
			}
			if !queue.Find(v.value) && !visited[v.value] {
				queue.Enqueue(v.value)
			}
		}
		fmt.Println("queue is", (*queue))
		fmt.Println("dequeue:", (*queue)[0])
		queue.Dequeue()
	}
}