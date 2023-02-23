// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 133.

// Упражнение 5.8. Измените функцию forEachNode так, чтобы функции рrе и
// post возвращали булево значение, указывающее, следует ли продолжать обход де
// рева. Воспользуйтесь ими для написания функции ElementBylD с приведенной
// ниже сигнатурой, которая находит первый HTML-элемент с указанным атрибутом
// id. Функция должна прекращать обход дерева, как только соответствующий элемент найден:
// func ElementByID(doc *html.Node, id string) *html.Node
// http://www.gopl.io/

package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		outline1(url)
	}
}

func outline1(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	node := ElementById(doc, "toc")
	if node != nil {
		fmt.Printf("%v", node.Attr)
	}
	return nil
}

func ElementById(doc *html.Node, id string) *html.Node {
	node := forEachNode1(doc, startElement_, endElement_, predicateID, id)
	return node
}
func predicateID(n *html.Node, id string) bool {
	if n.Attr != nil {
		for _, attribute := range n.Attr {
			if attribute.Key == "id" && attribute.Val == id {
				return true
			}
		}
	}
	return false
}

// !+forEachNode
// forEachNode calls the functions pre(x) and post(x) for each node
// x in the tree rooted at n. Both functions are optional.
// pre is called before the children are visited (preorder) and
// post is called after (postorder).
func forEachNode1(n *html.Node,
	pre, post func(*html.Node, func(*html.Node, string) bool, string) bool,
	predicate func(*html.Node, string) bool,
	str string) *html.Node {
	if pre != nil && pre(n, predicate, str) {
		return n
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		node := forEachNode1(c, pre, post, predicate, str)
		if node != nil {
			return node
		}
	}

	if post != nil {
		post(n, predicate, str)
	}
	return nil
}

//!-forEachNode

func startElement_(n *html.Node, predicate func(*html.Node, string) bool,
	str string) bool {
	if n.Type == html.ElementNode {
		return predicate(n, str)
	}
	return false
}

func endElement_(n *html.Node, predicate func(*html.Node, string) bool,
	str string) bool {
	if n.Type == html.ElementNode {
		//fmt.Printf("%s\n", n.Data)
	}
	return false
}
