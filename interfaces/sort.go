package main

import (
	"fmt"
	"sort"
	"time"
)

type Book struct {
	Name string
	Date time.Time
}

func (b Book) String() string {
	return fmt.Sprintf("Name: %s, Date: %s",
		b.Name,
		b.Date.Format("2006-01-02"))
}

type BookSlice []*Book

func (b BookSlice) Len() int {
	return len(b)
}

/*
	func (b BookSlice) Less(i, j int) bool {
		return b[i].Name < b[j].Name
	}
*/
func (b BookSlice) Less(i, j int) bool {
	if b[i].Date != b[j].Date {
		return b[i].Date.Before(b[j].Date)
	}
	if b[i].Name != b[j].Name {
		return b[i].Name < b[j].Name
	}
	return false
}

func (b BookSlice) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func main() {
	books := []*Book{
		{"Learning Algorithms", time.Now().AddDate(0, 0, -1)},
		{"JavaScript. Guide", time.Now().AddDate(0, -2, -5)},
		{"Effective Java", time.Date(2014, 10, 15, 0, 0, 0, 0, time.UTC)},
		{"Clean Code", time.Date(2009, 1, 1, 0, 0, 0, 0, time.UTC)},
	}
	//screen(books)
	//fmt.Printf("\n\n")
	sort.Sort(BookSlice(books))
	screen(books)

	println()

	sort.Sort(sort.Reverse(BookSlice(books)))
	screen(books)
}

func screen(books []*Book) {
	for _, book := range books {
		//fmt.Printf("%s \n", book)
		fmt.Printf("%-20s %s \n", book.Name, book.Date.Format("02.01.2006"))
	}
}
