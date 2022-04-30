package parser

import (
	"fmt"
	"log"
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

var links []Link

func Parsehtml(htmlString string) {
	doc, err := html.Parse(strings.NewReader(htmlString))
	if err != nil {
		log.Fatal(err)
	}
	iterateOverTree(doc)

	fmt.Printf("%+v\n", links)

}

func iterateOverTree(ele *html.Node) {
	if ele.Type == html.ElementNode && ele.Data == "a" {
		getLink(ele)
		return
	}
	for c := ele.FirstChild; c != nil; c = c.NextSibling {
		iterateOverTree(c)
	}
}

func getLink(ele *html.Node) {
	link := Link{}
	var text string
	for _, attr := range ele.Attr {
		if attr.Key == "href" {
			link.Href = attr.Val
		}
	}
	link.Text = strings.TrimSpace(getTextFromEle(ele, &text))
	links = append(links, link)
}

func getTextFromEle(ele *html.Node, text *string) string {

	if ele.Type == html.TextNode {
		*text += ele.Data
	}
	for c := ele.FirstChild; c != nil; c = c.NextSibling {
		getTextFromEle(c, text)
	}
	return *text
}
