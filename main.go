package main

import (
	"encoding/json"
	"fmt"
	"os"

	"golang.org/x/net/html"
)

type Bookmarks struct {
	Title    string
	URL      string
	IsDir    bool
	Children []*Bookmarks
}

func main() {

	file, err := os.Open("bookmarks_2022_3_18.html")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	defer file.Close()
	doc, err := html.Parse(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	var f func(*html.Node, *Bookmarks) *Bookmarks
	f = func(n *html.Node, b *Bookmarks) *Bookmarks {
		if n.Type == html.ElementNode && n.Data == "a" {
			var url, text string
			for _, a := range n.Attr {
				if a.Key == "href" {
					url = a.Val
					break
				}
			}
			if n.FirstChild != nil {
				text = n.FirstChild.Data
			}
			b.Children = append(b.Children, &Bookmarks{URL: url, Title: text})
			return b
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			if c.Type == html.ElementNode && c.Data == "h3" {
				var text string
				if c.FirstChild != nil {
					text = c.FirstChild.Data
				}
				dir := &Bookmarks{Title: text, IsDir: true}
				b.Children = append(b.Children, dir)
				b = dir
			} else if c.Type != html.TextNode && c.Type != html.DoctypeNode && c.Type != html.CommentNode {
				f(c, b)
			}
		}
		return b
	}
	res := f(doc, &Bookmarks{Title: "Bookmarks", IsDir: true})
	b, err := json.Marshal(*res)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(b))
	}
}
