package scrape

import (
	"context"
	"fmt"
	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/chromedp"
	"go-server/server/models"
	"go-server/util"
	"log"
	"math"
	"strconv"
	"time"
	"unicode"
)

// ApplyOffset function that extract only numbers from string
// if the function fails to extract a number it will return -1
func ApplyOffset(divName string, offset int) string {
	var num string
	for _, char := range divName {
		if unicode.IsNumber(char) {
			num += string(char)
		}
	}
	result, err := strconv.Atoi(num)
	if err != nil {
		return "nil"
	}
	result = result + offset
	return "jss" + strconv.Itoa(result)
}

// function that will scrape Catalog page according to the websiteID given
func scrapeCatalog(target string, websiteID string) []models.Product {
	// create chrome instance
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()

	// create a timeout
	ctx, cancel = context.WithTimeout(ctx, 200*time.Second)
	defer cancel()

	linkCleanPrefix := ""
	imgCleanPrefix := ""
	var priceArray []string
	var titleArray []string
	var imgArray []string
	var linkArray []string

	if websiteID == "ksp" {
		linkCleanPrefix = "https://ksp.co.il"
		var generatedNum string

		// this part will extract the randomly generated class name from the page
		err := chromedp.Run(ctx,
			emulation.SetUserAgentOverride("WebScraper 1.0"),
			chromedp.Navigate(target),

			chromedp.WaitVisible(`ul.slider.animated`, chromedp.ByQuery),

			// get the randomly generated offset
			chromedp.Evaluate(`document.querySelector('[aria-label="הוספה לפריטים שאהבתי"] > div').getAttribute("class")`, &generatedNum),
		)
		if err != nil {
			log.Fatal(err)
		}

		// generate the class name of the price and image
		targetClass := ApplyOffset(generatedNum, 5)
		priceSelector := fmt.Sprintf(`[...document.querySelectorAll("div.%s")].map((e) => e.innerText)`, targetClass)
		targetClass = ApplyOffset(generatedNum, -15)
		imgSelector := fmt.Sprintf(`[...document.querySelectorAll("div.%s > img")].map((e) => e.getAttribute("src"))`, targetClass)

		// get the arrays of the price, title and image
		err = chromedp.Run(ctx,
			chromedp.Evaluate(imgSelector, &imgArray),
			chromedp.Evaluate(`[...document.querySelectorAll("h3 a")].map((e) => e.innerText)`, &titleArray),
			//chromedp.Evaluate(`[...document.querySelectorAll("a.MuiTypography-root.MuiTypography-body1")].map((e) => e.getAttribute("aria-label"))`, &titleArray),
			chromedp.Evaluate(priceSelector, &priceArray),
			chromedp.Evaluate(`[...document.querySelectorAll("h3 a")].map((e) => e.getAttribute("href"))`, &linkArray),
		)
		if err != nil {
			log.Fatal(err)
		}
	} else if websiteID == "kravitz" {
		// get the arrays of the price, title and image
		err := chromedp.Run(ctx,
			emulation.SetUserAgentOverride("WebScraper 1.0"),
			chromedp.Navigate(target),

			chromedp.Evaluate(`[...document.querySelectorAll("img.product-image-photo")].map((e) => e.getAttribute("src"))`, &imgArray),
			chromedp.Evaluate(`[...document.querySelectorAll("h3 a")].map((e) => e.innerText)`, &titleArray),
			chromedp.Evaluate(`[...document.querySelectorAll("div.price-box > span.price-container.price-final_price.tax.weee")].map((e) => e.innerText)`, &priceArray),
			chromedp.Evaluate(`[...document.querySelectorAll("h3 a")].map((e) => e.getAttribute("href"))`, &linkArray),
		)
		if err != nil {
			log.Fatal(err)
		}
	} else if websiteID == "pc365" {
		linkCleanPrefix = "https://pc365.co.il/"
		// get the arrays of the price, title and image
		err := chromedp.Run(ctx,
			emulation.SetUserAgentOverride("WebScraper 1.0"),
			chromedp.Navigate(target),

			chromedp.Evaluate(`[...document.querySelectorAll("#s_prodDesktop div.s_prod_img img")].map((e) => e.getAttribute("src"))`, &imgArray),
			chromedp.Evaluate(`[...document.querySelectorAll("#s_prodDesktop div.s_prod_cont")].map((e) =>  e.firstChild.nodeValue)`, &titleArray),
			chromedp.Evaluate(`[...document.querySelectorAll("#s_prodDesktop div.s_price.lineprice > :first-child")].map((e) => e.innerText)`, &priceArray),
			chromedp.Evaluate(`[...document.querySelectorAll("div.prod_th_cont > a")].map((e) => e.getAttribute("href"))`, &linkArray),
		)
		if err != nil {
			log.Fatal(err)
		}
	} else if websiteID == "ivory" {
		imgCleanPrefix = "https://www.ivory.co.il/"

		opts := []chromedp.ExecAllocatorOption{
			//chromedp.ExecPath(`/Applications/Chromium.app/Contents/MacOS/Chromium`),
			chromedp.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/77.0.3830.0 Safari/537.36"),
			chromedp.WindowSize(1920, 1080),
			chromedp.NoFirstRun,
			chromedp.NoDefaultBrowserCheck,
			chromedp.Headless,
		}
		ctx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
		defer cancel()

		ctx, cancel = chromedp.NewContext(ctx)
		defer cancel()

		err := chromedp.Run(ctx,
			//emulation.SetUserAgentOverride("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/77.0.3830.0 Safari/537.36"),
			chromedp.Navigate(target),

			chromedp.Evaluate(`[...document.querySelectorAll("div.image-d-wrapper > img.img-fluid")].map((e) => e.getAttribute("data-src"))`, &imgArray),
			chromedp.Evaluate(`[...document.querySelectorAll("div.col-md-12.col-12.title_product_catalog.mb-md-1.main-text-area")].map((e) => e.innerText)`, &titleArray),
			chromedp.Evaluate(`[...document.querySelectorAll("div.text-right.right-position.regular-price.pricing-col.col-12.col-md-12  span.price-area span.price")].map((e) => e.innerText)`, &priceArray),
			chromedp.Evaluate(`[...document.querySelectorAll("a.product-anchor")].map((e) => e.getAttribute("href"))`, &linkArray),
		)
		if err != nil {
			log.Fatal(err)
		}
	}
	if len(priceArray) == 0 {
		return nil
	}
	//get the min length of the arrays
	minLength := math.Min(float64(len(titleArray)), float64(len(priceArray)))
	minLength = math.Min(minLength, float64(len(imgArray)))
	var minLen = int(minLength)

	products := make([]models.Product, minLen)

	// create array of products
	for i := 0; i < minLen; i++ {
		products[i] = util.NewProduct(titleArray[i], websiteID, util.GetPrice(priceArray[i]), imgCleanPrefix+imgArray[i], linkCleanPrefix+linkArray[i])
	}

	return products
}

// function that will scrape Item page according to the websiteID given
func scrapeItem(target string, websiteID string) models.Product {
	// create chrome instance
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()

	// create a timeout
	ctx, cancel = context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	// navigate to a page, wait for an element, click
	var price string
	var title string
	var img string

	if websiteID == "ksp" {
		var generatedNum string

		// this part will extract the randomly generated class name from the page
		err := chromedp.Run(ctx,
			emulation.SetUserAgentOverride("WebScraper 1.0"),
			chromedp.Navigate(target),

			chromedp.WaitVisible(`#review-section`, chromedp.ByQuery),

			// get the randomly generated offset
			chromedp.Evaluate(`document.querySelector("[data-id='top-product-page']").parentElement.className`, &generatedNum),
		)
		if err != nil {
			log.Fatal(err)
		}

		// generate the class name of the price and image
		targetClass := ApplyOffset(generatedNum, 33)
		priceSelector := fmt.Sprintf(`div.%s`, targetClass)
		targetClass = ApplyOffset(generatedNum, -11)
		titleSelector := fmt.Sprintf(`h1.MuiTypography-root.%s.MuiTypography-h1 > span`, targetClass)

		err = chromedp.Run(ctx,
			chromedp.Text(priceSelector, &price, chromedp.ByQuery),
			chromedp.Text(titleSelector, &title, chromedp.ByQuery),
			chromedp.AttributeValue(`li.slide.selected.previous > div > img`, "src", &img, nil, chromedp.ByQuery),
		)
		if err != nil {
			log.Fatal(err)
		}
	} else if websiteID == "kravitz" {

		err := chromedp.Run(ctx,
			emulation.SetUserAgentOverride("WebScraper 1.0"),
			chromedp.Navigate(target),

			chromedp.Text(`h1.page-title span`, &title, chromedp.ByQuery),
			chromedp.Text(`span.price`, &price, chromedp.ByQuery),

			// wait for image to load
			chromedp.WaitVisible(`#zoom1`, chromedp.ByQuery),

			chromedp.AttributeValue(`#zoom1`, "href", &img, nil, chromedp.ByQuery),
		)
		if err != nil {
			log.Fatal(err)
		}
	} else if websiteID == "pc365" {
		err := chromedp.Run(ctx,
			emulation.SetUserAgentOverride("WebScraper 1.0"),
			chromedp.Navigate(target),

			chromedp.Text(`#product-title`, &title, chromedp.ByQuery),
			chromedp.Text(`div.r b`, &price, chromedp.ByQuery),

			chromedp.AttributeValue(`#mainpic img`, "src", &img, nil, chromedp.ByQuery),
		)
		if err != nil {
			log.Fatal(err)
		}
	}
	return util.NewProduct(title, websiteID, util.GetPrice(price), img, target)
}

func WithDriver(link string, vendor string, targetType string) []models.Product {
	if targetType == "Item" {
		return []models.Product{scrapeItem(link, vendor)}
	} else {
		return scrapeCatalog(link, vendor)
	}
}
