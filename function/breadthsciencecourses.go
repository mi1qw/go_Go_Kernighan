// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 139.

// Findlinks3 crawls the web, starting with the URLs on the command line.
package main

import "fmt"

// prereqs maps computer science courses to their prerequisites.
var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

type Node struct {
	prev  []*Node
	next  []*Node
	value string
}

// !+breadthFirst
// breadthFirst calls f for each item in the worklist.
// Any items returned by f are added to the worklist.
// f is called at most once for each item.
func breadthFirst(f func(item string, m map[string]*Node,
	prereqs map[string][]string) []string,
	prereqs map[string][]string) map[string]*Node {

	worklist := make([]string, 0, len(prereqs))
	for k := range prereqs {
		worklist = append(worklist, k)
	}

	seen := make(map[string]bool)
	nodes := make(map[string]*Node)

	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true

				worklist = append(worklist, f(item, nodes, prereqs)...)
			}
		}
	}
	return nodes
}

//!-breadthFirst

// !+crawl
func crawl(s string, m map[string]*Node, prereqs map[string][]string) []string {
	list := prereqs[s]
	prev := getNode(s, m)
	for _, n := range list {
		nxt := getNode(n, m)
		prev.next = append(prev.next, nxt)
		nxt.prev = append(nxt.prev, prev)
	}
	return list
}

//!-crawl

func getNode(s string, m map[string]*Node) *Node {
	nodPtr, ok := m[s]
	if !ok {
		nodPtr = &Node{
			value: s,
			prev:  []*Node{},
			next:  []*Node{},
		}
		m[s] = nodPtr
	}
	return nodPtr
}
func screenNodes(m map[string]*Node) {
	depth := 0
	var f func(nod *Node)
	f = func(nod *Node) {
		fmt.Printf("%*s%s\n", depth*2, "", nod.value)
		depth++
		for _, node := range nod.next {
			f(node)
		}
		depth--
	}
	for _, nod := range m {
		if len(nod.prev) == 0 {
			f(nod)
			fmt.Printf("-------------------\n")
		}
	}
}

// !+main
func main() {
	nodes := breadthFirst(crawl, prereqs)
	screenNodes(nodes)
}

//!-main
