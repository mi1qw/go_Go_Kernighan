package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
)

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

var movies = []Movie{
	{Title: "Casablanca", Year: 1942, Color: false,
		Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
	{Title: "Cool Hand Luke", Year: 1967, Color: true,
		Actors: []string{"Paul Newman"}},
	{Title: "Bullitt", Year: 1968, Color: true,
		Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
}

func main() {
	// создать json
	data, err := json.Marshal(movies)
	if err != nil {
		log.Fatalf("Сбой маршалинга JSON: %s", err)
	}
	fmt.Printf("%s\n", data)

	// создать json в читаемом виде
	dataCmfrt, err := json.MarshalIndent(movies, "", "   ")
	fmt.Printf("%s\n-конец dataCmfrt\n\n", dataCmfrt)
	//-----------------------------------------------------
	// создать struct из json
	var titles []struct {
		Title string
		Color bool
	}
	if err := json.Unmarshal(data, &titles); err != nil {
		log.Fatalf("Сбой демаршалинга JSON: %s", err)
	}
	fmt.Println(titles)

	//-----------------------------------------------------
	// используется потоковый декодер json.Decoder , который
	// может декодировать несколько последовательных объектов JSON
	// из одного и того же потока
	ts := httptest.NewServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				_, _ = w.Write(data)
			}))
	defer ts.Close()
	res, _ := http.Get(ts.URL)
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(res.Body)
	var moviesDec []Movie
	d := json.NewDecoder(res.Body)
	err = d.Decode(&moviesDec)
	if err != nil {
		log.Fatalf("Decode: %v", err)
	}
	fmt.Printf("\nпотоковый декодер json.Decoder\n")
	for _, movie := range moviesDec {
		fmt.Printf("%s %d   цветной-%-10t\n    актёры: %s\n",
			movie.Title, movie.Year, movie.Color, movie.Actors)
	}
}
