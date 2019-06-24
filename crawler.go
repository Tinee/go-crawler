package crawler

import (
	"fmt"
	"net/http"
	"net/url"

	"golang.org/x/net/html"
)

type Crawler struct {
	work    chan *url.URL
	filters []FilterFunc
}

func New() *Crawler {
	return &Crawler{
		work: make(chan *url.URL, 1),
	}
}

func (c *Crawler) Crawl(target *url.URL) error {
	c.work <- target
	count := 0

	for {
		select {
		case u := <-c.work:
			count++
			if !c.passesFilters(u) {
				continue
			}

			fmt.Println(u.String())

			go func(u *url.URL) {
				for _, l := range getLinks(u) {
					c.work <- l
				}
			}(u)
		}
	}
}

func (c *Crawler) AddFilters(funcs ...FilterFunc) {
	c.filters = funcs
}

func (c *Crawler) passesFilters(u *url.URL) bool {
	for _, f := range c.filters {
		if !f(u) {
			return false
		}
	}
	return true
}

func getLinks(target *url.URL) []*url.URL {
	var result []*url.URL

	res, err := http.Get(target.String())
	if err != nil {
		return result
	}
	defer res.Body.Close()

	doc, err := html.Parse(res.Body)
	if err != nil {
		return result
	}
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					u, err := url.Parse(a.Val)
					if err != nil {
						continue
					}

					if u.Hostname() == "" {
						u = target.ResolveReference(u)
					}

					result = append(result, u)
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	return result
}
