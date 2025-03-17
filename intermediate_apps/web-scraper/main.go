package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"sync"

	"github.com/gocolly/colly"
)

type Product struct {
	Url   string
	Image string
	Name  string
	Price string
}

func main() {

	var products []Product
	var visitedUrls sync.Map

	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
		colly.AllowedDomains("www.scrapingcourse.com"),
	)

	// TODO: Add opportunity to avoid anti-bots algorithms there

	c.OnHTML("li.product", func(e *colly.HTMLElement) {
		product := Product{
			Url:   e.ChildAttr("a", "href"),
			Image: e.ChildAttr("img", "src"),
			Name:  e.ChildText(".product-name"),
			Price: e.ChildText(".price"),
		}

		products = append(products, product)
	})

	c.OnHTML("a.next", func(e *colly.HTMLElement) {
		nextPage := e.Attr("href")

		if _, found := visitedUrls.Load(nextPage); !found {
			fmt.Println("Scrapping: ", nextPage)

			visitedUrls.Store(nextPage, struct{}{})

			e.Request.Visit(nextPage)
		}
	})

	c.OnScraped(func(r *colly.Response) {
		file, err := os.Create("output.csv")
		if err != nil {
			fmt.Println("Error when try to create output.csv file: ", err)
		}
		defer file.Close()

		writer := csv.NewWriter(file)

		header := []string{
			"Url",
			"Image",
			"Name",
			"Price",
		}

		writer.Write(header)

		for _, product := range products {
			record := []string{
				product.Url,
				product.Image,
				product.Name,
				product.Price,
			}

			writer.Write(record)
		}

		defer writer.Flush()
	})

	c.Visit("https://www.scrapingcourse.com/ecommerce")
}
