// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 136.

/*
У пражнение 5.11. Преподаватель линейной алгебры (linear algebra) считает, что
до его курса следует прослушать курс матанализа (calculus). Перепишите функцию
topoSort так, чтобы она сообщала о наличии циклов.
*/

// The toposort program prints the nodes of a DAG in topological order.
package main

import (
	"fmt"
)

// !+table
// prereqs maps computer science courses to their prerequisites.
var prereqs = map[string]map[string]bool{
	"algorithms":     {"data structures": false},
	"calculus":       {"linear algebra": false},
	"linear algebra": {"calculus": false},

	"compilers": {
		"data structures":       false,
		"formal languages":      false,
		"computer organization": false,
	},

	"data structures":       {"discrete math": false, "compilers": false},
	"databases":             {"data structures": false},
	"discrete math":         {"intro to programming": false},
	"formal languages":      {"discrete math": false},
	"networks":              {"operating systems": false},
	"operating systems":     {"data structures": false, "computer organization": false},
	"programming languages": {"data structures": false, "computer organization": false},
}

//!-table

// !+main
func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func isCyclic(parent, child string) bool {
	if parent == "" || child == "" {
		return false
	}
	_, ok := prereqs[child][parent]
	return ok
}

//var parent, child string

func topoSort(m map[string]map[string]bool) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items map[string]bool, parent string)

	visitAll = func(items map[string]bool, parent string) {
		for item := range items {
			if !seen[item] {
				seen[item] = true
				if isCyclic(parent, item) {
					fmt.Printf("%s -> %s, has cyclic dependency\n", parent, item)
				}

				visitAll(m[item], item)
				order = append(order, item)
			}
		}
	}

	var keys = map[string]bool{}
	for k := range m {
		keys[k] = false
	}

	visitAll(keys, "")
	return order
}

//!-main
