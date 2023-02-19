/*
Упражнение 5.2. Напишите функцию для заполнения отображения, ключами ко
торого являются имена элементов (р, div , span и т.д.), а значениями — количество
элементов с таким именем в дереве HTML-документа.
Упражнение 5.3. Напишите функцию для вывода содержимого всех текстовых
узлов в дереве документа HTML. Не входите в элементы <script> и <style>,
скольку их содержимое в веб-браузере не является видимым.
Упражнение 5.4. Расширьте функцию visit так, чтобы она извлекала другие раз
новидности ссылок из документа, такие как изображения, сценарии и листы стилей.
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
	"regexp"
)

var textNodes []NodeVal
var imgLinkHrefs []NodeVal
var nameElemHrefMap = map[string]struct{}{
	"img":    {},
	"link":   {},
	"script": {},
	"style":  {},
}

type NodeVal struct {
	node *html.Node
	val  string
}

func main() {
	// todo добавить конец строки в регулярку
	//reg, err := regexp.Compile("^\n+\\s*\n*\\z")
	reg, err := regexp.Compile("^[\n\\s]*\\z")
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

	for _, node := range textNodes {
		str := node.val
		if !reg.MatchString(str) {
			fmt.Printf("%s", str)
		}
	}

	fmt.Printf("\n\n")
	for _, val := range imgLinkHrefs {
		fmt.Printf("%s\n", val.val)
	}
}

// !+visit
// visit appends to links each link found in n and returns the result.
func visit(links []string, elements map[string]int, n *html.Node) []string {
	checkElement(elements, n)
	saveElement(html.TextNode,
		func(*html.Node) bool {
			return n.Parent.Data != "script" &&
				n.Parent.Data != "noscript" &&
				n.Parent.Data != "style"
		},
		func(*html.Node) string {
			return n.Data
		},
		n, &textNodes)

	saveElement(html.ElementNode,
		func(*html.Node) bool {
			if _, ok := nameElemHrefMap[n.Data]; ok {
				return true
			}
			return false
		},
		func(*html.Node) string {
			for _, a := range n.Attr {
				if a.Key == "href" || a.Key == "src" {
					return a.Val
				}
			}
			return "nil"
		},
		n, &imgLinkHrefs)

	//fmt.Printf("----%#v\n", *n)
	//println(&links, len(links), n.Parent, &n)
	//	println(&links, len(links), n.Parent, &n, n.Data)
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
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

// textNodes *[]NodeVal
func saveElement(typeNode html.NodeType,
	predicate func(*html.Node) bool,
	get func(*html.Node) string,
	n *html.Node, list *[]NodeVal) {
	if n.Type == typeNode && predicate(n) {
		*list = append(*list, NodeVal{n, get(n)})
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
