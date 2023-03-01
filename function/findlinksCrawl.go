// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 139.

/*
У пражнение 5.13. Модифицируйте функцию crawl так, чтобы она делала локальные
копии найденных ею страниц, при необходимости создавая каталоги. Не делайте
копии страниц, полученных из других доменов. Например, если исходная страница
поступает с адреса golang.org, сохраняйте все страницы оттуда, но не сохраняйте
страницы, например, сvimeo.com.

http://www.gopl.io/
https://www.java.com/
https://golang-blog.blogspot.com/
*/

// Findlinks3 crawls the web, starting with the URLs on the command line.
package main

import (
	"bufio"
	"fmt"
	"gopl.io/ch5/links"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/user"
	"path"
	"strings"
)

type SaveCnfg struct {
	userPath   string
	folderName string
	url        string
	host       string
}

// !+breadthFirst
// breadthFirst calls f for each item in the worklist.
// Any items returned by f are added to the worklist.
// f is called at most once for each item.
func breadthFirst(f func(item string, cfg SaveCnfg) []string, cfg SaveCnfg) {
	seen := make(map[string]bool)
	var worklist = []string{cfg.url}
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item, cfg)...)
			}
		}
	}
}

//!-breadthFirst

// !+crawl
func crawl(url_ string, cfg SaveCnfg) []string {
	err := savePage(url_, cfg)
	if err != nil {
		log.Print(err)
	}
	list, err := links.Extract(url_)
	if err != nil {
		log.Print(err)
	}
	var links []string
	for _, link := range list {
		parse, _ := url.Parse(link)
		if parse.Host == cfg.host {
			links = append(links, link)
		}
	}
	return links
}

func savePage(url string, cfg SaveCnfg) error {
	resp, err := http.Get(url)
	if err != nil {
		println(err)
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		fmt.Fprintf(os.Stderr, "getting %s: %s", url, resp.Status)
	}
	defer resp.Body.Close()
	if resp.Request.URL.Host == cfg.host {
		println(resp.Request.URL.Host)
		println(resp.Request.URL.Path)
		dir, name := urlParse(resp.Request.URL.Path)
		fmt.Printf("`%s`  `%s`\n", dir, name)
		err = createFile(dir, name, resp.Body, cfg)

		if err != nil {
			return err
		}
	}
	return nil
}

func createFile(dir string, name string, body io.ReadCloser, cfg SaveCnfg) error {
	if name == "" && dir == "" {
		name = "index.html"
	}
	if dir != "" {
		dir = path.Join(cfg.userPath, dir)
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			panic(err)
		}
	} else {
		dir = cfg.userPath
	}
	if name == "" {
		_, name = urlParse(dir)
	}
	if !strings.HasSuffix(name, ".html") {
		name = name + ".html"
	}
	file, err := os.OpenFile(
		path.Join(dir, name),
		os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		println(err)
		return err
	}
	defer file.Close()

	bytes, err := io.ReadAll(body)
	if err != nil {
		println(err)
		return err
	}
	bufWriter := bufio.NewWriter(file)
	_, err = bufWriter.Write(bytes)
	if err != nil {
		return err
	}
	err = bufWriter.Flush()
	if err != nil {
		return err
	}
	return nil
}

//!-crawl

// !+main
func main() {
	// Crawl the web breadth-first,
	// starting from the command-line arguments.
	parseUrl, err := url.Parse(os.Args[1:][0])
	if err != nil {
		log.Println(err)
		return
	}

	cfg := SaveCnfg{
		url:        os.Args[1:][0],
		host:       parseUrl.Host,
		userPath:   "",
		folderName: "site",
	}
	cfg.userPath = path.Join(projectDir(), cfg.folderName)

	err = makeDir(cfg.userPath)
	if err != nil {
		log.Println(err)
		return
	}
	breadthFirst(crawl, cfg)
	//содержимое папки
	// listDirByReadDir(".")
}

func urlParse(url string) (string, string) {
	index := strings.LastIndex(url, "/")
	if index == -1 {
		return url, ""
	}
	return url[:index], url[index+1:]
}

//!-main

func homeDir() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return usr.HomeDir
}

func projectDir() string {
	getwd, err := os.Getwd()
	if err != nil {
		log.Println(err)
		return ""
	}
	return getwd
}

func makeDir(name string) error {
	err := os.Mkdir(name, 0777)
	if err != nil {
		panic(err)
	}
	return nil
}

func listDirByReadDir(path string) {
	dir, err := os.ReadDir(path)
	if err != nil {
		panic(err)
	}
	for _, entry := range dir {
		if entry.IsDir() {
			fmt.Printf("[%s]\n", entry.Name())
		} else {
			fmt.Println(entry.Name())
		}
	}
}
