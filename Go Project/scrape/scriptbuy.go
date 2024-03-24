package scrape

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
	"go-server/server/models"
	"go-server/util"
	"strings"
)

func WithColly(targetUrl string, cleanUrl string, vendor string, targetType string, selectors models.Selector) []models.Product {
	//create array of products
	var products []models.Product
	var link string

	collector := colly.NewCollector()
	extensions.RandomUserAgent(collector)
	extensions.Referer(collector)

	collector.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting:", r.URL)
	})
	collector.OnResponse(func(r *colly.Response) {
		fmt.Println("Status:", r.StatusCode)
	})
	collector.OnError(func(r *colly.Response, err error) {
		if err.Error() == "abort" {
			return
		} else {
			fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
		}
	})

	collector.OnHTML(selectors.MainSelector, func(e *colly.HTMLElement) {
		itemTitle := e.ChildText(selectors.TitleSelector)
		itemPrice := util.GetPrice(e.ChildText(selectors.PriceSelector))
		itemImg := selectors.ImgSelectorPrefix + e.ChildAttr(selectors.ImageSelector[0], selectors.ImageSelector[1])

		if targetType == "Item" {
			link = targetUrl
		} else {
			extractedUrl := e.ChildAttr(selectors.UrlSelector[0], selectors.UrlSelector[1])
			// if the link is a full link (http:// or https://) then use it as is else add the cleanUrl prefix
			if strings.HasPrefix(strings.ToLower(extractedUrl), "http") {
				link = extractedUrl
			} else {
				link = cleanUrl + extractedUrl
			}
		}

		//util.PrintProduct(util.NewProduct(itemTitle, vendor, itemPrice, itemImg, targetUrl))
		products = append(products, util.NewProduct(itemTitle, vendor, itemPrice, itemImg, link))
	})

	err := collector.Visit(targetUrl)
	if err != nil {
		return nil
	}
	return products
}
