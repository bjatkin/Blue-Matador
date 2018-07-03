package main

import (
	"strings"
)

//THOUGHS - if there are 2 bottlenecks or more we can abandon the process and return an empty array
//a bottleneck is a single path that never reconects with a path behind it

func findPath(graph string) [][]string {
	g := newGraph()
	for _, node := range strings.Split(graph, "\n") {
		g.build(node)
	}

	paths := g.fullSearch()
	if len(paths) == 0 {
		return nil
	}
	return paths
}

func search(g nodeGraph, n *node, path map[*edge]int, test []string, all [][]string, depth, maxDepth int) [][]string {
	var end *node
	for _, e := range n.edges {
		if path[e] > 0 {
			continue
		}
		end = e.a
		if end == n {
			end = e.b
		}
		p := make(map[*edge]int)
		for k, v := range path {
			p[k] = v
		}
		p[e] = depth
		t := append(test, end.name)
		if depth == maxDepth {
			all = append(all, t)
			continue
		}
		res := search(g, end, p, t, all, depth+1, maxDepth)
		if len(res) > len(all) {
			all = res
		}
	}
	return all
}

type node struct {
	name  string
	edges []*edge
}

type edge struct {
	a *node
	b *node
}

type nodeGraph struct {
	nodes map[string]*node
	edges map[string]*edge
}

func (graph nodeGraph) fullSearch() [][]string {
	var allPaths [][]string
	for _, node := range graph.nodes {
		paths := graph.search(node, make(map[*edge]bool), []string{node.name}, [][]string{})
		for _, path := range paths {
			allPaths = append(allPaths, path)
		}
	}

	return allPaths
}

func (graph nodeGraph) search(start *node, visited map[*edge]bool, path []string, ret [][]string) [][]string {
	for _, edg := range start.edges {
		if visited[edg] {
			continue
		}

		end := edg.a
		if end == start {
			end = edg.b
		}
		newVisited := copyVisited(visited)
		newVisited[edg] = true

		newPath := append(path, end.name)

		if len(path) == len(graph.edges) {
			ret = append(ret, newPath)
			continue
		}

		ret = graph.search(end, newVisited, newPath, ret)
	}
	return ret
}

func copyVisited(old map[*edge]bool) map[*edge]bool {
	newMap := make(map[*edge]bool)
	for k, v := range old {
		newMap[k] = v
	}
	return newMap
}

func newGraph() nodeGraph {
	return nodeGraph{
		nodes: make(map[string]*node),
		edges: make(map[string]*edge),
	}
}

func (graph nodeGraph) findEdge(a, b *node) *edge {
	if abEdge, ok := graph.edges[a.name+b.name]; ok {
		return abEdge
	}
	if baEdge, ok := graph.edges[b.name+a.name]; ok {
		return baEdge
	}
	return nil
}

func (graph *nodeGraph) addEdge(a, b *node) *edge {
	add := graph.findEdge(a, b)
	if add == nil {
		add = &edge{
			a: a,
			b: b,
		}
		graph.edges[a.name+b.name] = add
	}

	return add
}

func (graph nodeGraph) findNode(name string) *node {
	if n, ok := graph.nodes[name]; ok {
		return n
	}
	return nil
}

func (graph *nodeGraph) addNode(name string) *node {
	if graph.findNode(name) == nil {
		graph.nodes[name] = &node{
			name: name,
		}
	}

	return graph.nodes[name]
}

func (graph *nodeGraph) build(outline string) {
	name := strings.Split(outline, " ")[0]
	a := graph.addNode(name)

	edges := strings.Split(outline, " ")[1:]
	for _, n := range edges {
		b := graph.addNode(n)
		e := graph.addEdge(a, b)

		a.edges = append(a.edges, e)
	}
}
