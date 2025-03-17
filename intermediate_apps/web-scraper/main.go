package main

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/gocolly/colly"
)

type Product struct {
	Url   string
	Image string
	Name  string
	Price string
}

func main() {

	pagesToScrape := []string{
		"https://www.scrapingcourse.com/ecommerce/page/1/",
		"https://www.scrapingcourse.com/ecommerce/page/2/",
		"https://www.scrapingcourse.com/ecommerce/page/3/",
		"https://www.scrapingcourse.com/ecommerce/page/4/",
		"https://www.scrapingcourse.com/ecommerce/page/5/",
		"https://www.scrapingcourse.com/ecommerce/page/6/",
		"https://www.scrapingcourse.com/ecommerce/page/7/",
		"https://www.scrapingcourse.com/ecommerce/page/8/",
		"https://www.scrapingcourse.com/ecommerce/page/9/",
		"https://www.scrapingcourse.com/ecommerce/page/10/",
		"https://www.scrapingcourse.com/ecommerce/page/11/",
		"https://www.scrapingcourse.com/ecommerce/page/12/",
	}

	var products []Product

	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
		colly.AllowedDomains("www.scrapingcourse.com"),
		colly.Async(true),
	)

	c.Limit(&colly.LimitRule{
		Parallelism: 12,
	})

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

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visited: ", r.URL)
	})

	for _, pageToScrape := range pagesToScrape {
		c.Visit(pageToScrape)

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
	}

	c.Wait()

}
