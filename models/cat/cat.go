package cat

import (
	"errors"
	"flag"
	"fmt"
	"log"

	"github.com/camiloaromero23/cat-scraper-api/db"
	"github.com/camiloaromero23/cat-scraper-api/scrapes"
	"github.com/camiloaromero23/cat-scraper-api/types"
	"gorm.io/gorm"
)

func GetCat(id string) (*types.Cat, error) {
	defer db.CloseDB()
	db := db.GetDB()

	var cat types.Cat
	res := db.First(&cat, id)

	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, fmt.Errorf("Error getting cat %s", res.Error)
	}

	return &cat, nil
}

func GetCats() ([]types.Cat, error) {
	defer db.CloseDB()
	db := db.GetDB()

	var cats []types.Cat
	res := db.Find(&cats)

	if res.Error != nil {
		return nil, fmt.Errorf("Error getting cats %s", res.Error)
	}

	return cats, nil
}

func UpdateCats() ([]types.Cat, error) {
	defer db.CloseDB()
	sync := flag.Bool("sync", false, "Run synchronously")
	flag.Parse()

	const (
		domain = "www.expertoanimal.com"
		url    = "https://www.expertoanimal.com/"
	)

	var endpoints = []string{
		"razas-de-gatos.html",
		"razas-de-gatos_2.html",
	}

	params := types.ScrapeParams{
		Domain:    domain,
		Url:       url,
		Endpoints: endpoints,
	}

	var cats []types.Cat

	if *sync {
		cats = scrapes.SyncScrape(params)
	} else {
		cats = scrapes.GoroutinesScrape(params)
	}

	log.Println("Scraping cats finished")

	log.Printf("Cats found: %d\n", len(cats))

	db := db.GetDB()

	if err := db.AutoMigrate(&types.Cat{}); err != nil {
		log.Println("Error migrating cats db", err)
	}

	db.Exec("DELETE FROM cats")

	res := db.Create(&cats)

	if res.Error != nil {
		log.Println("Error saving cats", res.Error)
		return nil, fmt.Errorf("Error saving cats %s", res.Error)
	}

	log.Println("Cats saved successfully")
	return cats, nil
}
