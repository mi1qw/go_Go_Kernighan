/*
Упражнение 5.2. Напишите функцию для заполнения отображения, ключами ко
торого являются имена элементов (р, div , span и т.д.), а значениями — количество
элементов с таким именем в дереве HTML-документа.


*/

// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 122.
//!+main

// Findlinks1 prints the links in an HTML document read from standard input.
package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

func main() {
	var elements = map[string]int{
		"div":    0,
		"p":      0,
		"span":   0,
		"a":      0,
		"i":      0,
		"img":    0,
		"li":     0,
		"ul":     0,
		"script": 0,
		"style":  0,
	}
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, elements, doc) {
		fmt.Printf("%#v\n", link)
	}
	fmt.Printf("\n%-5s %s\n", "name", "col")
	for k, v := range elements {
		fmt.Printf("%-5s %d\n", k, v)
	}
}

//!-main

// !+visit
// visit appends to links each link found in n and returns the result.
func visit(links []string, elements map[string]int, n *html.Node) []string {
	checkElement(elements, n)
	fmt.Printf("----%#v\n", *n)
	//println(&links, len(links), n.Parent, &n)
	//	println(&links, len(links), n.Parent, &n, n.Data)
	//fmt.Printf("dfata-- %#v\n", n.Data)
	if n.Type == html.ElementNode && n.Data == "a" {
		//fmt.Printf("%#v", *n)
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	if n.NextSibling != nil {
		var _ = "s"
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, elements, c)
	}
	return links
}

func checkElement(elements map[string]int, n *html.Node) {
	if _, ok := elements[n.Data]; ok {
		elements[n.Data]++
	}
}

//!-visit

/*
//!+html
package html

type Node struct {
	Type                    NodeType
	Data                    string
	Attr                    []Attribute
	FirstChild, NextSibling *Node
}

type NodeType int32

const (
	ErrorNode NodeType = iota
	TextNode
	DocumentNode
	ElementNode
	CommentNode
	DoctypeNode
)

type Attribute struct {
	Key, Val string
}

func Parse(r io.Reader) (*Node, error)
//!-html
*/
