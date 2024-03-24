package util

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
	"go-server/server/models"
	"regexp"
	"strings"
	"unicode"
)

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// GetPrice function that extracts price from the text by removing all non-numeric characters except for the dot
// this function is using the strings.Builder for efficiency
func GetPrice(text string) string {
	var builder strings.Builder
	for _, char := range text {
		if unicode.IsNumber(char) || char == '.' {
			builder.WriteRune(char)
		}
	}
	return builder.String()
}

func NewProduct(title string, vendor string, price string, img string, link string) models.Product {
	return models.Product{Title: title, Vendor: vendor, Price: price, Img: img, Link: link}
}

func CheckLinkByURL(link string) string {
	// if the link is from ksp , check if it is an item or a category
	if regexp.MustCompile(`^(https?://)?(www\.)?ksp`).MatchString(link) {
		if regexp.MustCompile(`^(https?://)?(www\.)?ksp\.co\.il/(web|mob)/item/`).MatchString(link) {
			return "ksp Item"
		} else if regexp.MustCompile(`^(https?://)?(www\.)?ksp\.co\.il/(web|mob)/cat/`).MatchString(link) {
			return "ksp Cat"
		} else {
			return "Err"
		}
	} else if regexp.MustCompile(`^(https?://)?(www\.)?1pc`).MatchString(link) {
		if regexp.MustCompile(`^(https?://)?(www\.)?1pc\.co\.il/he/product`).MatchString(link) {
			return "1pc Item"
		} else if regexp.MustCompile(`^(https?://)?(www\.)?1pc\.co\.il/he/cat`).MatchString(link) {
			return "1pc Cat"
		} else {
			return "Err"
		}
	} else if regexp.MustCompile(`^(https?://)?(www\.)?allincell`).MatchString(link) {
		if regexp.MustCompile(`^(https?://)?(www\.)?allincell\.co\.il/product/`).MatchString(link) {
			return "allincell Item"
		} else if regexp.MustCompile(`^(https?://)?(www\.)?allincell\.co\.il/product-category/`).MatchString(link) {
			return "allincell Cat"
		} else {
			return "Err"
		}
	} else if regexp.MustCompile(`^(https?://)?(www\.)?bug`).MatchString(link) {
		if regexp.MustCompile(`^(https?://)?(www\.)?bug\.co\.il/brand/`).MatchString(link) {
			return "bug Item"
		} else {
			return "bug Cat"
		} // bug does not have specific categories indicator in the link
	} else if regexp.MustCompile(`^(https?://)?(www\.)?ivory`).MatchString(link) {
		if regexp.MustCompile(`^(https?://)?(www\.)?ivory\.co\.il/catalog\.php\?id=`).MatchString(link) {
			return "ivory Item"
		} else {
			return "ivory Cat"
		} // ivory does not have specific categories indicator in the link
	} else if regexp.MustCompile(`^(https?://)?(www\.)?officedepot`).MatchString(link) {
		if regexp.MustCompile(`^(https?://)?(www\.)?officedepot\.co\.il/product/`).MatchString(link) {
			return "officedepot Item"
		} else if regexp.MustCompile(`^(https?://)?(www\.)?officedepot\.co\.il/product-category/`).MatchString(link) {
			return "officedepot Cat"
		} else {
			return "Err"
		}
	} else if regexp.MustCompile(`^(https?://)?(www\.)?kravitz`).MatchString(link) {
		if regexp.MustCompile(`^(https?://)?(www\.)?kravitz\.co\.il/\d\d\d`).MatchString(link) {
			return "kravitz Item"
		} else {
			return "kravitz Cat"
		} // kravitz does not have specific categories indicator in the link
	} else if regexp.MustCompile(`^(https?://)?(www\.)?terminalx`).MatchString(link) {
		if regexp.MustCompile(`^(https?://)?(www\.)?terminalx\.com[^\s]+color=\d+$`).MatchString(link) {
			return "terminalx Item"
		} else {
			return "terminalx Cat"
		} // terminalx does not have specific categories indicator in the link
	} else if regexp.MustCompile(`^(https?://)?(www\.)?pc365`).MatchString(link) {
		if regexp.MustCompile(`^(https?://)?(www\.)?pc365\.co\.il/product`).MatchString(link) {
			return "pc365 Item"
		} else if regexp.MustCompile(`^(https?://)?(www\.)?pc365\.co\.il/cat`).MatchString(link) {
			return "pc365 Cat"
		} else {
			return "Err"
		}
	}
	return "Err"
}

func CheckLinkByTags(link string) string {
	var isItem bool = false
	if regexp.MustCompile(`^(https?://)?(www\.)?1pc`).MatchString(link) {
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

		// 1pc.co.il -> Item
		collector.OnHTML("#product-details-form", func(e *colly.HTMLElement) {
			isItem = true
		})

		// 1pc.co.il -> Category
		collector.OnHTML("div.category-page", func(e *colly.HTMLElement) {
			isItem = false
		})

		err := collector.Visit(link)
		if err != nil {
			return ""
		}

		if isItem {
			return "1pc Item"
		} else {
			return "1pc Cat"
		}
	} else if regexp.MustCompile(`^(https?://)?(www\.)?officedepot`).MatchString(link) {
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

		// officedepot.co.il -> Item
		collector.OnHTML("#tab-label-description", func(e *colly.HTMLElement) {
			isItem = true
		})

		// officedepot.co.il -> Category
		collector.OnHTML("div.products-grid", func(e *colly.HTMLElement) {
			isItem = false
		})

		err := collector.Visit(link)
		if err != nil {
			return ""
		}

		if isItem {
			return "officedepot Item"
		} else {
			return "officedepot Cat"
		}
	} else if regexp.MustCompile(`^(https?://)?(www\.)?adcs`).MatchString(link) {
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

		// adcs.co.il -> Item
		collector.OnHTML("div.product-info-price", func(e *colly.HTMLElement) {
			isItem = true
		})

		// adcs.co.il -> Category
		collector.OnHTML("div.filter-content", func(e *colly.HTMLElement) {
			isItem = false
		})

		err := collector.Visit(link)
		if err != nil {
			return ""
		}

		if isItem {
			return "adcs Item"
		} else {
			return "adcs Cat"
		}
	}
	return "Err"
}

func CheckLink(link string) string {
	fastCheck := CheckLinkByURL(link)
	if fastCheck == "Err" {
		slowCheck := CheckLinkByTags(link)
		if slowCheck == "Err" {
			return "Err"
		} else {
			return slowCheck
		}
	} else {
		return fastCheck
	}
}
