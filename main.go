package main

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"log"
)

// Film contains the fields that a film should have.
// Some of these should probably be broken down into slices e.g. Cast
type Film struct {
	Artist           string
	Cast             string
	Cinematographers string
	Countries        string
	Director         string
	Editors          string
	Language         string
	Name             string
	Pitch            string
	Premiere         string
	Producers        string
	Production       string
	Runtime          string
	Score            string
	Screenplay       string
	Sound            string
	Year             string
}

// Selectors is a bit of overkill, but it makes updating for 2017 and beyond pretty easy
var Selectors = map[string]string{
	"Artist":           "#artist .credit-content",
	"Cast":             "#cast .credit-content",
	"Cinematographers": "#cinematographers .credit-content",
	"Countries":        "span.quick-info .countries",
	"Director":         "#director .credit-content",
	"Editors":          "#editors .credit-content",
	"Language":         "span.quick-info .language",
	"Name":             ".container h1",
	"Pitch":            ".pitch",
	"Premiere":         "span.quick-info .premiere",
	"Producers":        "#producers .credit-content",
	"Production":       "#productionCompany .credit-content",
	"Runtime":          "span.quick-info .runtime",
	"Score":            "#originalScore .credit-content",
	"Screenplay":       "#screenplay .credit-content",
	"Sound":            "#sound .credit-content",
	"Year":             "span.quick-info .year",
}

// getUrls reads a urls.json from disk and unmarshals it into a slice of strings
func getUrls() []string {
	// setup the slice to hold resulting list from json
	var urls []string

	// read the file
	raw, err := ioutil.ReadFile("./urls.json")
	if err != nil {
		log.Fatal(err.Error())
	}

	// unmarshal the json into the urls slice
	json.Unmarshal(raw, &urls)
	return urls
}

// writeJSON just chucks some bytes to disk
func writeJSON(b []byte) {
	err := ioutil.WriteFile("films/films.json", b, 0644)
	if err != nil {
		panic(err)
	}
}

// scrapeFilm loads the remote URL and parses it using goquery to populate the fields of a Film struct
func scrapeFilm(url string, ch chan Film, chFinished chan bool) {
	var f Film

	defer func() {
		// Notify that scrapeFilm call is done when this function ends
		chFinished <- true
	}()

	// fetch the URL and tokenize/parse it using goquery
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}

	// Extract the #wrap element, and parse the contents for required fields
	doc.Find("#wrap").Each(func(_ int, s *goquery.Selection) {
		f = Film{
			Artist:           s.Find(Selectors["Artist"]).Text(),
			Cast:             s.Find(Selectors["Cast"]).Text(),
			Cinematographers: s.Find(Selectors["Cinematographers"]).Text(),
			Countries:        s.Find(Selectors["Countries"]).Text(),
			Director:         s.Find(Selectors["Director"]).Text(),
			Editors:          s.Find(Selectors["Editors"]).Text(),
			Language:         s.Find(Selectors["Language"]).Text(),
			Name:             s.Find(Selectors["Name"]).Text(),
			Pitch:            s.Find(Selectors["Pitch"]).Text(),
			Premiere:         s.Find(Selectors["Premiere"]).Text(),
			Producers:        s.Find(Selectors["Producers"]).Text(),
			Production:       s.Find(Selectors["Production"]).Text(),
			Runtime:          s.Find(Selectors["Runtime"]).Text(),
			Score:            s.Find(Selectors["Score"]).Text(),
			Screenplay:       s.Find(Selectors["Screenplay"]).Text(),
			Sound:            s.Find(Selectors["Sound"]).Text(),
			Year:             s.Find(Selectors["Year"]).Text(),
		}
		// throw the result into the channel
		ch <- f
	})
}

func main() {
	// setup an empty slice of Films
	var films []Film

	// open the list of films in urls.json
	urls := getUrls()

	// setup channels for concurrent parsing
	chFilms := make(chan Film)
	chFinished := make(chan bool)

	// scrape each URL and put the result in chFilms channel
	for _, url := range urls {
		go scrapeFilm(url, chFilms, chFinished)
	}

	// read channels until all film urls have been parsed
	for c := 0; c < len(urls); {
		select {
		case film := <-chFilms:
			films = append(films, film)
		case <-chFinished:
			c++
		}
	}

	// tidy up
	close(chFilms)

	// format the resulting list of films as json
	f, err := json.MarshalIndent(films, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	// write the json to output.json
	writeJSON(f)

	// print the json to the screen
	fmt.Println(string(f))

}
