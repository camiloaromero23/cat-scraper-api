package scrapes

import (
	"fmt"
	"log"
	"sync"

	"github.com/camiloaromero23/cat-scraper-api/types"
	"github.com/gocolly/colly/v2"
)

func GoroutinesScrape(params types.ScrapeParams) []types.Cat {

	wg := sync.WaitGroup{}

	c := colly.NewCollector(colly.AllowedDomains(params.Domain))

	c.OnHTML(".resultado.link", scrapeCats)

	for _, endpoint := range params.Endpoints {
		wg.Add(1)
		go func(endpoint string) {
			err := c.Visit(fmt.Sprintf("%s%s", params.Url, endpoint))
			if err != nil {
				log.Printf("Scraping failed - %s\n", err)
			}
			wg.Done()
		}(endpoint)
	}
	wg.Wait()

	d := colly.NewCollector(colly.AllowedDomains(params.Domain))
	d.OnHTML("article.columna-post", scrapeCatsDetails)

	for _, cat := range cats {
		wg.Add(1)
		go func(cat types.Cat) {
			err := d.Visit(cat.Link)
			if err != nil {
				fmt.Printf("Scraping failed - %s\n", err)
			}
			wg.Done()
		}(cat)
	}
	wg.Wait()

	return cats
}
