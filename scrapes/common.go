package scrapes

import (
	"github.com/camiloaromero23/cat-scraper-api/types"
	"github.com/gocolly/colly/v2"
)

var cats []types.Cat

func scrapeCats(e *colly.HTMLElement) {
	name := e.ChildText("a.titulo.titulo--resultado")
	link := e.ChildAttr("a.titulo.titulo--resultado", "href")

	cats = append(cats, types.Cat{
		Name:        name,
		Image:       "",
		Description: "",
		Link:        link,
	})
}

func scrapeCatsDetails(e *colly.HTMLElement) {
	description := e.ChildText("p")
	catName := e.ChildText("h1.titulo.titulo--articulo")
	image := e.ChildAttr(".intro .imagen_wrap .imagen picture > img", "src")
	for i, cat := range cats {
		if cat.Name != catName {
			continue
		}

		cats[i].Description = description
		cats[i].Image = image
	}
}
