package main

import "fmt"

func main() {
	// graph with adjacency list
	graph := map[string][]string{
		"a": {"b", "c"},
		"b": {"d"},
		"c": {"e"},
		"d": {"f"},
		"e": {"a"}, //cycle in graph
		"f": {},
	}
	dfs(graph, "a")
	bfs(graph, "a")
}

// dfs traverse the graph in inorder traversal, where first the children nodes of root are visited before the siblings
func dfs(graph map[string][]string, startingNode string) {
	fmt.Println("dfs traversal")
	// dfs can be traversed using a simple stack mechanism
	// where the visiting nodes can be pushed on stack, and once visited pop it from the stack
	// and push its children on the stack
	// in next iteration, pop the stack (LIFO), this will ensure that first the childrens of recent pushed nodes are visited
	// before their siblings
	// Initialize the stack with starting node as its first element
	stack := []string{startingNode}
	// Also in order to overcome cycles in graph, we have to keep track of nodes which are already visited
	visited := map[string]bool{startingNode: true}
	// traverse until there are elements in stack
	for len(stack) > 0 {
		// the last element on stack is to be poped, print it
		n := len(stack) - 1
		popedNode := stack[n]
		fmt.Println(popedNode)
		// reset the stack by removing poped element
		stack = stack[:n]
		// traverse the childrens of poped node, and push them on stack
		for _, node := range graph[popedNode] {
			// only add the node on stack if is not visited before
			if !visited[node] {
				stack = append(stack, node)
				// after adding node on stack, mark it as visited, so we dont fall in cycle
				visited[node] = true
			}
		}
	}
	// once the stack is empty, the dfs traversal is complete
}

func bfs(graph map[string][]string, startingNode string) {
	fmt.Println("bfs traversal")
	// bfs can be traversed using a simple queue mechanism
	// where the visiting nodes are pushed on queue, and once visited, it is dequed
	// and its childrens are enqueued
	// in the next iteration, the dequed element will be [FIFO] the sibling of visited node
	// and the its childrens will be visited next in order
	// initialized the queue with starting node as first element of queue
	queue := []string{startingNode}
	// Also in order to overcome cycles in graph, we have to keep track of nodes which are already visited
	visited := map[string]bool{startingNode: true}
	// traverse untile queue is empty
	for len(queue) > 0 {
		// first element of queue has to be dequed
		dequedElement := queue[0]
		fmt.Println(dequedElement)
		// reset the queue by removing first element
		queue = queue[1:]
		// traverse the adjaceny list of dequed element and add the childrens on queue
		for _, node := range graph[dequedElement] {
			// only add the node on queue if is not visited before
			if !visited[node] {
				queue = append(queue, node)
				// after adding node on queue, mark it as visited, so we dont fall in cycle
				visited[node] = true
			}
		}
	}
	// once the queue is empty, the dfs traversal is complete
}
