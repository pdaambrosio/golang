package main

import (
	"fmt"

	"github.com/gocolly/colly"
	"github.com/sclevine/agouti"
)

func Scraping(
	url string,
	selector string,
) {
	c := colly.NewCollector()

	c.OnHTML(selector, func(e *colly.HTMLElement) {
		fmt.Println(e.Text)
	})

	c.Visit(url)
}

func ScrapingWithAgouti(url string, selector string) {
    driver := agouti.ChromeDriver(
        agouti.ChromeOptions("args", []string{"--headless", "--no-sandbox", "--disable-dev-shm-usage"}),
    )

	driver.Start()
	page, _ := driver.NewPage()

	page.Navigate(url)
	element := page.Find(selector)
	text, _ := element.Text()
	fmt.Println(text)
}

func main() {
	Scraping("https://www.google.com", "title")
	ScrapingWithAgouti("https://www.google.com", "title")
}
