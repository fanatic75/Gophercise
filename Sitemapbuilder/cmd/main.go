package main

import (
	"encoding/xml"
	"excercises/htmlParser/parser"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type handler struct{}

var urlProvided string
var maxDepth int

var levelTraversed int = 0

var linksIns map[string]string = make(map[string]string)

type loc struct {
	Loc string `xml:"loc"`
}

type urlset struct {
	Xmlns string `xml:"xmlns,attr"`
	Url   []loc  `xml:"url"`
}

func init() {

	flag.StringVar(&urlProvided, "Domain", "https://gophercises.com/", "Enter the domain to scrape all the anchor tag links")
	flag.IntVar(&maxDepth, "Max Depth", 3, "How deep should the parser go?")
	flag.Parse()
	linksIns[urlProvided] = urlProvided
}

func main() {
	getHtmlAndAppendUrl(urlProvided)
	uset := urlset{Xmlns: "http://www.sitemaps.org/schemas/sitemap/0.9"}
	for _, value := range linksIns {
		uset.Url = append(uset.Url, loc{value})

	}
	output, err := xml.MarshalIndent(uset, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(xml.Header)

	os.Stdout.Write(output)
}

func getHtmlAndAppendUrl(url string) {
	if maxDepth == levelTraversed {

		return
	}

	html, URL := getHTMLFromUrl(url)
	links := parser.Parsehtml(html)
	for _, link := range links {

		if strings.HasPrefix(link.Href, "#") {
			continue
		}
		if strings.HasPrefix(link.Href, "/") {
			if _, ok := linksIns[URL.Host+link.Href]; ok {
				continue
			}
			linksIns[URL.Host+link.Href] = "https://" + URL.Host + link.Href

			levelTraversed++
			getHtmlAndAppendUrl(linksIns[URL.Host+link.Href])
		} else {
			if !strings.HasPrefix(link.Href, URL.Host) {
				continue
			}
			if _, ok := linksIns[link.Href]; ok {
				continue
			}
			linksIns[link.Href] = link.Href

			levelTraversed++
			getHtmlAndAppendUrl(linksIns[link.Href])

		}

	}
}

func getHTMLFromUrl(url string) (string, *url.URL) {

	r, err := http.Get(url)
	if err != nil {
		return "", nil
	}
	defer r.Body.Close()

	data, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("unable to read data from get body of url")
		return "", nil
	}
	return string(data), r.Request.URL

}
