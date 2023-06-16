package scrapes

import (
  "fmt"
	"log"

	"github.com/camiloaromero23/cat-scraper-api/types"
	"github.com/gocolly/colly/v2"
)

func SyncScrape(params types.ScrapeParams) []types.Cat {

	c := colly.NewCollector(colly.AllowedDomains(params.Domain))

	c.OnHTML(".resultado.link", scrapeCats)

	for _, endpoint := range params.Endpoints {
		err := c.Visit(fmt.Sprintf("%s%s", params.Url, endpoint))
		if err != nil {
			log.Printf("Scraping failed - %s\n", err)
		}
	}

	d := colly.NewCollector(colly.AllowedDomains(params.Domain))
	d.OnHTML("article.columna-post", scrapeCatsDetails)

	for _, cat := range cats {
		err := d.Visit(cat.Link)
		if err != nil {
			log.Printf("Scraping failed - %s\n", err)
		}
	}

  return cats
}
