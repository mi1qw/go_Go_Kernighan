// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 133.

// У пражнение 5.7. Разработайте startElement и endElement для обобщенного
// вывода HTML. Выводите узлы комментариев, текстовые узлы и атрибуты каждого
// элемента (<а href=*...’ >). Используйте сокращенный вывод наподобие <img/>
// вместо <img></img>, когда элемент не имеет дочерних узлов.
// http://www.gopl.io/

// Outline prints the outline of an HTML document tree.
package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
	"regexp"
)

var reg *regexp.Regexp

func main() {
	reg, _ = regexp.Compile("^[^\\w\\$\\*\\+\\?\\{\\}\\[\\]\\\\\\|\\(\\)]*$")
	for _, url := range os.Args[1:] {
		outline(url)
	}
}

func outline(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	//!+call
	forEachNode(doc, startElement, endElement)
	//!-call

	return nil
}

// !+forEachNode
// forEachNode calls the functions pre(x) and post(x) for each node
// x in the tree rooted at n. Both functions are optional.
// pre is called before the children are visited (preorder) and
// post is called after (postorder).
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

//!-forEachNode

// !+startend
func checkSpace(s string) bool {
	if reg.MatchString(s) {
		return true
	}
	return false
}

var depth int

func prntAttr(n *html.Node, depth int) {
	if n.Attr != nil {
		if len(n.Attr) > 3 {
			for _, attribute := range n.Attr {
				fmt.Printf("\n%*s%s=\"%s\"",
					depth*2, "", attribute.Key, attribute.Val)
			}
		} else {
			for _, attribute := range n.Attr {
				fmt.Printf(" %s=\"%s\"", attribute.Key, attribute.Val)
			}
		}
	}
}

func getPrevType(n *html.Node) html.NodeType {
	for n != nil && checkSpace(n.Data) {
		n = n.PrevSibling
	}
	if n == nil {
		return html.ElementNode
	}
	return n.Type
}

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		if getPrevType(n.PrevSibling) == html.TextNode {
			fmt.Printf("<%s", n.Data)
		} else {
			fmt.Printf("\n%*s<%s", depth*2, "", n.Data)
		}
		depth++
		prntAttr(n, depth)
		if n.LastChild == nil {
			fmt.Printf("/>")
		} else {
			fmt.Printf(">")
		}
	} else {
		if n.Type == html.CommentNode {
			fmt.Printf("\n<!-- %s -->", n.Data)
		} else {
			if !checkSpace(n.Data) {
				if n.Parent != nil && n.Parent.Type == html.ElementNode {
					fmt.Printf("%s", n.Data)
				} else {
					fmt.Printf("%*s%s", depth*2, "", n.Data)
				}
			}
		}
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		typeN := getPrevType(n.LastChild)
		if n.LastChild != nil {
			if typeN == html.TextNode {
				fmt.Printf("</%s>", n.Data)
			} else {
				fmt.Printf("\n%*s</%s>", depth*2, "", n.Data)
			}
		}
	}
}

//!-startend
