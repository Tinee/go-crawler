package main

import (
	"log"
	"net/url"

	"github.com/Tinee/crawler"
)

func main() {
	var (
		// answer = ""
		c = crawler.New()
	)

	// fmt.Println("Where do you want to crawl?")
	// fmt.Scanln(&answer)

	u, err := url.Parse("https://blog.marcuskarlsson.com/")
	if err != nil {
		log.Fatalln(err)
	}

	c.AddFilters(
		crawler.Unique(),
		crawler.SameHost(u),
	)

	c.Crawl(u)
}
