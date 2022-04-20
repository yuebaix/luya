package util

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
	"net/http"
)

func DefaultCollector() *colly.Collector {
	c := colly.NewCollector(
	//colly.Async(true),
	//colly.AllowedDomains("hackerspaces.org", "wiki.hackerspaces.org"),
	)
	c.WithTransport(&http.Transport{
		DisableKeepAlives: true,
	})
	extensions.RandomUserAgent(c)

	setDefaultCallback(c)
	return c
}

func setDefaultCallback(c *colly.Collector) {
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		fmt.Printf("Link found: %q -> %s\n", e.Text, link)
		//c.Visit(e.Request.AbsoluteURL(link))
	})
}
