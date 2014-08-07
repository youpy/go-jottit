package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/youpy/go-jottit"
)

func main() {
	var name, pageUrl string
	var showRev, list bool
	var revId int

	flag.StringVar(&name, "name", "", "site name(XXX.jottit.com)")
	flag.StringVar(&pageUrl, "page", "", "page url(http://XXX.jottit.com/YYY)")
	flag.BoolVar(&showRev, "revisions", false, "show revision ids of given page")
	flag.BoolVar(&list, "list", false, "list page urls")
	flag.IntVar(&revId, "revision", 1, "revision id")
	flag.Parse()

	if name != "" {
		site, err := jottit.NewSite("http://" + name)
		if err != nil {
			log.Fatal(err)
		}

		if list {
			pages, err := site.GetPages()
			if err != nil {
				log.Fatal(err)
			}

			for _, page := range pages {
				fmt.Println(page)
			}
		}
	}

	if pageUrl != "" {
		page, err := jottit.NewPage(pageUrl)
		if err != nil {
			log.Fatal(err)
		}

		if showRev {
			revs, err := page.GetRevisions()
			if err != nil {
				log.Fatal(err)
			}

			for _, rev := range revs {
				fmt.Println(rev.Id)
			}
		} else if revId != 0 {
			rev := jottit.NewRevision(page, revId)

			time, err := rev.GetPostTime()
			if err != nil {
				log.Fatal(err)
			}

			content, err := rev.GetContent()
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(time)
			fmt.Println(content)
		}
	}
}
