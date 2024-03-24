package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-server/scrape"
	"go-server/server/database"
	"go-server/server/models"
	"go-server/util"
	"go.mongodb.org/mongo-driver/bson"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

// InitDB - Initialize Database with default values
func InitDB(ctx *gin.Context) {
	db := database.Connect()
	collection := db.Client.Database("ScriptBuy").Collection("scrape")
	err := collection.Drop(db.Ctx)
	if err != nil {
		ctx.IndentedJSON(http.StatusFailedDependency, gin.H{"Error": "Failed to drop collection"})
		return
	}
	_, _ = collection.InsertMany(db.Ctx, []interface{}{
		bson.D{
			{"_cat", "Electronics"},
			{"websites", []interface{}{
				bson.D{
					{"_name", "1pc"},
					{"url", "https://1pc.co.il"},
					{"logo_url", "https://1pc.co.il/images/thumbs/0076482_logo.e8a255e0.png"},
					{"selectors", []interface{}{
						bson.D{
							{"type", "Item"},
							{"main_selector", "#product-details-form"},
							{"title_selector", "div.product-name > h1"},
							{"price_selector", "div.product-price > span"},
							{"image_selector", []interface{}{"#cloudZoomImage", "src"}},
						},
						bson.D{
							{"type", "Cat"},
							{"main_selector", "div.product-grid > div.item-grid > div.item-box"},
							{"title_selector", "h2.product-title > a"},
							{"price_selector", "span.actual-price"},
							{"image_selector", []interface{}{"div.picture > a > img", "data-lazyloadsrc"}},
							{"url_selector", []interface{}{"h2.product-title > a", "href"}},
						},
					}},
				},
				bson.D{
					{"_name", "ksp"},
					{"url", "https://ksp.co.il"},
					{"logo_url", "https://ksp.co.il/meNew/img/logos/computers_logo.png"},
					{"selectors", []interface{}{}},
				},
				bson.D{
					{"_name", "kravitz"},
					{"url", "https://kravitz.co.il"},
					{"logo_url", "https://www.kravitz.co.il/media/logo/stores/1/LOGO.webp"},
					{"selectors", []interface{}{}},
				},
				bson.D{
					{"_name", "pc365"},
					{"url", "https://pc365.co.il"},
					{"logo_url", "https://www.pc365.co.il/img/logo.png"},
					{"selectors", []interface{}{}},
				},
				bson.D{
					{"_name", "allincell"},
					{"url", "https://allincell.co.il"},
					{"logo_url", "https://allincell.co.il/wp-content/uploads/2022/10/aic-logo-black.png"},
					{"selectors", []interface{}{
						bson.D{
							{"type", "Item"},
							{"main_selector", "div.row.top-image-content"},
							{"title_selector", "h3.title"},
							{"price_selector", "div.custom_price span.woocommerce-Price-amount.amount"},
							{"image_selector", []interface{}{"div.woocommerce-product-gallery__image > a", "href"}},
						},
						bson.D{
							{"type", "Cat"},
							{"main_selector", "div.product_div"},
							{"title_selector", "h5 > a"},
							{"price_selector", "span.woocommerce-Price-amount.amount"},
							{"image_selector", []interface{}{"img.img-responsive", "src"}},
							{"url_selector", []interface{}{"h5 > a", "href"}},
						},
					}},
				},
				bson.D{
					{"_name", "bug"},
					{"url", "https://bug.co.il"},
					{"logo_url", "https://www.bug.co.il/images/logoBug.png"},
					{"selectors", []interface{}{
						bson.D{
							{"type", "Item"},
							{"main_selector", "#product-page"},
							{"title_selector", "#product-page-prodict-title"},
							{"price_selector", "#product-price-container ins"},
							{"image_selector", []interface{}{"img.pm-gaee", "src"}},
							{"img_selector_prefix", "https://www.bug.co.il"},
						},
						bson.D{
							{"type", "Cat"},
							{"main_selector", "div.span_4_of_12.col.float-right-col.product-cube.pc-gaee"},
							{"title_selector", "h4 > a.tpurl"},
							{"price_selector", "div.price > span"},
							{"image_selector", []interface{}{"a.image > img.img-lazy-load", "data-original"}},
							{"img_selector_prefix", "https://www.bug.co.il"},
							{"url_selector", []interface{}{"a.image", "href"}},
						},
					}},
				},
				bson.D{
					{"_name", "ivory"},
					{"url", "https://ivory.co.il"},
					{"logo_url", "https://img.zap.co.il/pics/imgs/nsite/newui/newssite-1546.gif"},
					{"selectors", []interface{}{
						bson.D{
							{"type", "Item"},
							{"main_selector", "#productslider"},
							{"title_selector", "#titleProd"},
							{"price_selector", "span.print-actual-price , table.d-none tr:nth-child(2) td.discount-price"},
							{"image_selector", []interface{}{"#img_zoom_inout", "src"}},
							{"img_selector_prefix", "https://www.ivory.co.il/"},
						},
						bson.D{
							{"type", "Cat"},
							{"main_selector", "div.row.p-1.entry-wrapper"},
							{"title_selector", "div.col-md-12.col-12.title_product_catalog.mb-md-1.main-text-area"},
							{"price_selector", "div.regular-price span.price-area"},
							{"image_selector", []interface{}{"div.image-d-wrapper > img.img-fluid", "data-src"}},
							{"img_selector_prefix", "https://www.ivory.co.il/splendid_images/cache/"},
							{"url_selector", []interface{}{"a.product-anchor", "href"}},
						},
					}},
				},
				bson.D{
					{"_name", "officedepot"},
					{"url", "https://officedepot.co.il"},
					{"logo_url", "https://www.officedepot.co.il/media/logo/stores/1/logo-b2c.png"},
					{"selectors", []interface{}{
						bson.D{
							{"type", "Item"},
							{"main_selector", "div.officedepot-product"},
							{"title_selector", "h1.page-title"},
							{"price_selector", "span.price-wrapper span.price"},
							{"image_selector", []interface{}{"#amasty-main-image", "src"}},
						},
						bson.D{
							{"type", "Cat"},
							{"main_selector", "li.product-item"},
							{"title_selector", "li.product-item a.product-item-link"},
							{"price_selector", "span.price-wrapper span.price"},
							{"image_selector", []interface{}{"img.product-image-photo", "src"}},
							{"url_selector", []interface{}{"a.product-item-photo", "href"}},
						},
					}},
				},
				bson.D{
					{"_name", "adcs"},
					{"url", "https://www.adcs.co.il/"},
					{"logo_url", "https://automations.co.il/wp-content/uploads/thegem-logos/logo_f7af2914bbdc5d1e72b35dc126346f3b_1x.png"},
					{"selectors", []interface{}{
						bson.D{
							{"type", "Item"},
							{"main_selector", "#maincontent"},
							{"title_selector", "h1.page-title > span.base"},
							{"price_selector", "span[data-price-type='finalPrice'] > span"},
							{"image_selector", []interface{}{"img.gallery-placeholder__image", "src"}},
						},
						bson.D{
							{"type", "Cat"},
							{"main_selector", "li.product-item"},
							{"title_selector", "a.product-item-link"},
							{"price_selector", "span[data-price-type='finalPrice'] > span"},
							{"image_selector", []interface{}{"img.product-image-photo", "src"}},
							{"url_selector", []interface{}{"a.product-item-link", "href"}},
						},
					}},
				},
			},
			},
			{"trending", []interface{}{}},
		},
		bson.D{
			{"_cat", "Clothing, Shoes & Accessories"},
			{"websites", []interface{}{
				bson.D{
					{"_name", "terminalx"},
					{"url", "https://terminalx.com"},
					{"logo_url", "https://upload.wikimedia.org/wikipedia/he/4/44/Terminal_X_logo.png"},
					{"selectors", []interface{}{
						bson.D{
							{"type", "Item"},
							{"main_selector", "div.top_1-oZ.rtl_3OXU"},
							{"title_selector", "h1.name_20R6"},
							{"price_selector", "div.row_2tcG.bold_2wBM.prices-final_1R9x"},
							{"image_selector", []interface{}{"img.image_3k9y.image-element_22jc", "src"}},
						},
						bson.D{
							{"type", "Cat"},
							{"main_selector", "li.listing-product_3mjp"},
							{"title_selector", "a.title_3ZxJ.tx-link_29YD.underline-hover_3GkV"},
							{"price_selector", "div.row_2tcG.bold_2wBM.final-price_8CiX"},
							{"image_selector", []interface{}{"img.image_3k9y", "src"}},
							{"url_selector", []interface{}{"a.tx-link_29YD", "href"}},
						},
					}},
				},
				bson.D{
					{"_name", "ksp"},
					{"url", "https://ksp.co.il"},
					{"logo_url", "https://ksp.co.il/meNew/img/logos/computers_logo.png"},
					{"selectors", []interface{}{}},
				},
				bson.D{
					{"_name", "kravitz"},
					{"url", "https://kravitz.co.il"},
					{"logo_url", "https://www.kravitz.co.il/media/logo/stores/1/LOGO.webp"},
					{"selectors", []interface{}{}},
				},
				bson.D{
					{"_name", "pc365"},
					{"url", "https://pc365.co.il"},
					{"logo_url", "https://www.pc365.co.il/img/logo.png"},
					{"selectors", []interface{}{}},
				},
			},
			},
			{"trending", []interface{}{}},
		},
		bson.D{
			{"_cat", "Home & Patio"},
			{"websites", []interface{}{
				bson.D{
					{"_name", "officedepot"},
					{"url", "https://officedepot.co.il"},
					{"logo_url", "https://www.officedepot.co.il/media/logo/stores/1/logo-b2c.png"},
					{"selectors", []interface{}{
						bson.D{
							{"type", "Item"},
							{"main_selector", "div.officedepot-product"},
							{"title_selector", "h1.page-title"},
							{"price_selector", "span.price-wrapper span.price"},
							{"image_selector", []interface{}{"#amasty-main-image", "src"}},
						},
						bson.D{
							{"type", "Cat"},
							{"main_selector", "li.product-item"},
							{"title_selector", "li.product-item a.product-item-link"},
							{"price_selector", "span.price-wrapper span.price"},
							{"image_selector", []interface{}{"img.product-image-photo", "src"}},
							{"url_selector", []interface{}{"a.product-item-photo", "href"}},
						},
					}},
				},
				bson.D{
					{"_name", "ksp"},
					{"url", "https://ksp.co.il"},
					{"logo_url", "https://ksp.co.il/meNew/img/logos/computers_logo.png"},
					{"selectors", []interface{}{}},
				},
				bson.D{
					{"_name", "kravitz"},
					{"url", "https://kravitz.co.il"},
					{"logo_url", "https://www.kravitz.co.il/media/logo/stores/1/LOGO.webp"},
					{"selectors", []interface{}{}},
				},
				bson.D{
					{"_name", "pc365"},
					{"url", "https://pc365.co.il"},
					{"logo_url", "https://www.pc365.co.il/img/logo.png"},
					{"selectors", []interface{}{}},
				},
			},
			},
			{"trending", []interface{}{}},
		},
		bson.D{
			{"_cat", "Baby"},
			{"websites", []interface{}{
				bson.D{
					{"_name", "ksp"},
					{"url", "https://ksp.co.il"},
					{"logo_url", "https://ksp.co.il/meNew/img/logos/computers_logo.png"},
					{"selectors", []interface{}{}},
				},
				bson.D{
					{"_name", "kravitz"},
					{"url", "https://kravitz.co.il"},
					{"logo_url", "https://www.kravitz.co.il/media/logo/stores/1/LOGO.webp"},
					{"selectors", []interface{}{}},
				},
				bson.D{
					{"_name", "pc365"},
					{"url", "https://pc365.co.il"},
					{"logo_url", "https://www.pc365.co.il/img/logo.png"},
					{"selectors", []interface{}{}},
				},
			}},
			{"trending", []interface{}{}},
		},
		bson.D{
			{"_cat", "School & Office"},
			{"websites", []interface{}{
				bson.D{
					{"_name", "ksp"},
					{"url", "https://ksp.co.il"},
					{"logo_url", "https://ksp.co.il/meNew/img/logos/computers_logo.png"},
					{"selectors", []interface{}{}},
				},
				bson.D{
					{"_name", "kravitz"},
					{"url", "https://kravitz.co.il"},
					{"logo_url", "https://www.kravitz.co.il/media/logo/stores/1/LOGO.webp"},
					{"selectors", []interface{}{}},
				},
				bson.D{
					{"_name", "pc365"},
					{"url", "https://pc365.co.il"},
					{"logo_url", "https://www.pc365.co.il/img/logo.png"},
					{"selectors", []interface{}{}},
				},
			}},
			{"trending", []interface{}{}},
		},
		bson.D{
			{"_cat", "Toys"},
			{"websites", []interface{}{
				bson.D{
					{"_name", "ksp"},
					{"url", "https://ksp.co.il"},
					{"logo_url", "https://ksp.co.il/meNew/img/logos/computers_logo.png"},
					{"selectors", []interface{}{}},
				},
				bson.D{
					{"_name", "kravitz"},
					{"url", "https://kravitz.co.il"},
					{"logo_url", "https://www.kravitz.co.il/media/logo/stores/1/LOGO.webp"},
					{"selectors", []interface{}{}},
				},
				bson.D{
					{"_name", "pc365"},
					{"url", "https://pc365.co.il"},
					{"logo_url", "https://www.pc365.co.il/img/logo.png"},
					{"selectors", []interface{}{}},
				},
			}},
			{"trending", []interface{}{}},
		},
		bson.D{
			{"_cat", "Sports, Fitness & Outdoors"},
			{"websites", []interface{}{
				bson.D{
					{"_name", "ksp"},
					{"url", "https://ksp.co.il"},
					{"logo_url", "https://ksp.co.il/meNew/img/logos/computers_logo.png"},
					{"selectors", []interface{}{}},
				},
				bson.D{
					{"_name", "kravitz"},
					{"url", "https://kravitz.co.il"},
					{"logo_url", "https://www.kravitz.co.il/media/logo/stores/1/LOGO.webp"},
					{"selectors", []interface{}{}},
				},
				bson.D{
					{"_name", "pc365"},
					{"url", "https://pc365.co.il"},
					{"logo_url", "https://www.pc365.co.il/img/logo.png"},
					{"selectors", []interface{}{}},
				},
			}},
			{"trending", []interface{}{}},
		},
		bson.D{
			{"_cat", "Entertainment"},
			{"websites", []interface{}{
				bson.D{
					{"_name", "ksp"},
					{"url", "https://ksp.co.il"},
					{"logo_url", "https://ksp.co.il/meNew/img/logos/computers_logo.png"},
					{"selectors", []interface{}{}},
				},
				bson.D{
					{"_name", "kravitz"},
					{"url", "https://kravitz.co.il"},
					{"logo_url", "https://www.kravitz.co.il/media/logo/stores/1/LOGO.webp"},
					{"selectors", []interface{}{}},
				},
				bson.D{
					{"_name", "pc365"},
					{"url", "https://pc365.co.il"},
					{"logo_url", "https://www.pc365.co.il/img/logo.png"},
					{"selectors", []interface{}{}},
				},
			}},
			{"trending", []interface{}{}},
		},
		bson.D{
			{"_cat", "Beauty & Personal Care"},
			{"websites", []interface{}{
				bson.D{
					{"_name", "ksp"},
					{"url", "https://ksp.co.il"},
					{"logo_url", "https://ksp.co.il/meNew/img/logos/computers_logo.png"},
					{"selectors", []interface{}{}},
				},
				bson.D{
					{"_name", "kravitz"},
					{"url", "https://kravitz.co.il"},
					{"logo_url", "https://www.kravitz.co.il/media/logo/stores/1/LOGO.webp"},
					{"selectors", []interface{}{}},
				},
				bson.D{
					{"_name", "pc365"},
					{"url", "https://pc365.co.il"},
					{"logo_url", "https://www.pc365.co.il/img/logo.png"},
					{"selectors", []interface{}{}},
				},
			}},
			{"trending", []interface{}{}},
		},
		bson.D{
			{"_cat", "Health"},
			{"websites", []interface{}{
				bson.D{
					{"_name", "ksp"},
					{"url", "https://ksp.co.il"},
					{"logo_url", "https://ksp.co.il/meNew/img/logos/computers_logo.png"},
					{"selectors", []interface{}{}},
				},
				bson.D{
					{"_name", "kravitz"},
					{"url", "https://kravitz.co.il"},
					{"logo_url", "https://www.kravitz.co.il/media/logo/stores/1/LOGO.webp"},
					{"selectors", []interface{}{}},
				},
				bson.D{
					{"_name", "pc365"},
					{"url", "https://pc365.co.il"},
					{"logo_url", "https://www.pc365.co.il/img/logo.png"},
					{"selectors", []interface{}{}},
				},
			}},
			{"trending", []interface{}{}},
		},
		bson.D{
			{"_cat", "Household Essentials"},
			{"websites", []interface{}{
				bson.D{
					{"_name", "ksp"},
					{"url", "https://ksp.co.il"},
					{"logo_url", "https://ksp.co.il/meNew/img/logos/computers_logo.png"},
					{"selectors", []interface{}{}},
				},
				bson.D{
					{"_name", "kravitz"},
					{"url", "https://kravitz.co.il"},
					{"logo_url", "https://www.kravitz.co.il/media/logo/stores/1/LOGO.webp"},
					{"selectors", []interface{}{}},
				},
				bson.D{
					{"_name", "pc365"},
					{"url", "https://pc365.co.il"},
					{"logo_url", "https://www.pc365.co.il/img/logo.png"},
					{"selectors", []interface{}{}},
				},
			}},
			{"trending", []interface{}{}},
		},
		bson.D{
			{"_cat", "Pets"},
			{"websites", []interface{}{
				bson.D{
					{"_name", "ksp"},
					{"url", "https://ksp.co.il"},
					{"logo_url", "https://ksp.co.il/meNew/img/logos/computers_logo.png"},
					{"selectors", []interface{}{}},
				},
				bson.D{
					{"_name", "kravitz"},
					{"url", "https://kravitz.co.il"},
					{"logo_url", "https://www.kravitz.co.il/media/logo/stores/1/LOGO.webp"},
					{"selectors", []interface{}{}},
				},
				bson.D{
					{"_name", "pc365"},
					{"url", "https://pc365.co.il"},
					{"logo_url", "https://www.pc365.co.il/img/logo.png"},
					{"selectors", []interface{}{}},
				},
			}},
			{"trending", []interface{}{}},
		},
		bson.D{
			{"_cat", "Grocery"},
			{"websites", []interface{}{
				bson.D{
					{"_name", "ksp"},
					{"url", "https://ksp.co.il"},
					{"logo_url", "https://ksp.co.il/meNew/img/logos/computers_logo.png"},
					{"selectors", []interface{}{}},
				},
				bson.D{
					{"_name", "kravitz"},
					{"url", "https://kravitz.co.il"},
					{"logo_url", "https://www.kravitz.co.il/media/logo/stores/1/LOGO.webp"},
					{"selectors", []interface{}{}},
				},
				bson.D{
					{"_name", "pc365"},
					{"url", "https://pc365.co.il"},
					{"logo_url", "https://www.pc365.co.il/img/logo.png"},
					{"selectors", []interface{}{}},
				},
			}},
			{"trending", []interface{}{}},
		},
	})

	ctx.IndentedJSON(http.StatusCreated, gin.H{"message": "Database initialized successfully"})

}

// GetCategorySelectors returns the scrape information for a specific category of websites
func GetCategorySelectors(ctx *gin.Context, categoryName string, result *models.Category) {
	db := database.Connect()
	collection := db.Client.Database("ScriptBuy").Collection("scrape")
	filter := bson.M{"_cat": categoryName}
	err := collection.FindOne(db.Ctx, filter).Decode(result)
	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "No data found"})
		return
	}
}

func CreateVisitedMap(products []models.Product) map[string]bool {
	visited := make(map[string]bool)
	for _, product := range products {
		visited[product.Vendor] = true
	}
	return visited
}

// GetVendorsData returns the data of all the websites in a specific category
func GetVendorsData(category models.Category, visited map[string]bool) []models.WebsiteData {
	var websitesData []models.WebsiteData
	for _, website := range category.Websites {
		if visited[website.Name] {
			websitesData = append(websitesData, models.WebsiteData{Name: website.Name, LogoUrl: website.LogoUrl})
		}
	}
	return websitesData
}

// GetWebsiteSelectors returns the scrape information for a specific website
func GetWebsiteSelectors(category models.Category, websiteName string, websiteType string) models.Selector {
	for _, website := range category.Websites {
		if website.Name == websiteName {
			if websiteType == "Item" {
				return website.Selectors[0]
			} else {
				return website.Selectors[1]
			}
		}
	}
	return models.Selector{}
}

// GetCleanWebsiteUrl returns the clean URL of a website given its name
func GetCleanWebsiteUrl(category models.Category, websiteName string) string {
	for _, website := range category.Websites {
		if website.Name == websiteName {
			return website.Url
		}
	}
	return ""
}

// GetTrending get a list of result of a search and update the trending list of a category accordingly
func updateTrending(ctx *gin.Context, categoryName string, productsResult []models.Product) {
	db := database.Connect()
	collection := db.Client.Database("ScriptBuy").Collection("scrape")
	filter := bson.D{{"_cat", categoryName}}
	var category models.Category
	err := collection.FindOne(db.Ctx, filter).Decode(&category)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to find the category in the database"})
		return
	}
	trendingMap := make(map[string]bool)
	trendingLen := len(category.Trending)
	productsResultLen := len(productsResult)
	rand.Shuffle(productsResultLen, func(i, j int) { productsResult[i], productsResult[j] = productsResult[j], productsResult[i] })
	for _, product := range category.Trending {
		trendingMap[product.Link] = true
	}

	if trendingLen < 10 {
		remainingSpace := 10 - trendingLen
		maxProductsToAdd := util.Min(remainingSpace, productsResultLen)
		for i := 0; i < productsResultLen; i++ {
			if maxProductsToAdd == 0 {
				break
			}
			if !trendingMap[productsResult[i].Link] {
				category.Trending = append(category.Trending, productsResult[i])
				trendingMap[productsResult[i].Link] = true
				maxProductsToAdd--
			}
		}
	} else {
		numReplacements := util.Min(trendingLen, rand.Intn(3)+2)
		replaceIndices := rand.Perm(trendingLen)[:numReplacements]
		for i := 0; i < productsResultLen; i++ {
			if !trendingMap[productsResult[i].Link] {
				category.Trending[replaceIndices[numReplacements-1]] = productsResult[i]
				trendingMap[productsResult[i].Link] = true
				numReplacements--
			}
			if numReplacements == 0 {
				break
			}
		}
	}
	_, err = collection.UpdateOne(db.Ctx, filter, bson.D{{"$set", bson.D{{"trending", category.Trending}}}})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update the category in the database"})
		return
	}
}

func driverConcurrentScrape(times int, url, websiteName, websiteType string) []models.Product {
	var wg sync.WaitGroup
	resultChan := make(chan []models.Product, times)

	for i := 1; i <= 7; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			result := scrape.WithDriver(url, websiteName, websiteType)
			if result != nil {
				resultChan <- result
			}
		}()
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	for result := range resultChan {
		return result
	}

	return nil
}

// Search is the main function that starts the scraping process given a query
func Search(ctx *gin.Context) {
	var query models.Query
	if err := ctx.BindJSON(&query); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"Error": "Please provide a valid query"})
		return
	}

	//Step 1: Identify the website category using GPT
	Category := scrape.Classify(query.Query)
	if Category == "" || Category == "none" {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to classify the query"})
		return
	}

	//Step 2: Get URL lists from Google API Call
	SearchResults := scrape.GetResultsSerpAPI(query.Query, 30)
	// SearchResults := scrape.GetResultsGoogleCustom(query.Query, 30)

	//Step 3: fetch Category data from MongoDB
	var categoryData models.Category
	GetCategorySelectors(ctx, Category, &categoryData)

	var wg sync.WaitGroup
	websitesVisited := make(map[string]bool)
	results := make(chan []models.Product)

	// Worker function that performs scraping tasks
	worker := func(url string) {
		defer wg.Done()

		linkIdentifier := util.CheckLink(url)
		if linkIdentifier == "Err" {
			return
		}
		targetData := strings.Split(linkIdentifier, " ")
		websiteName := targetData[0]
		websitesVisited[websiteName] = true
		websiteType := targetData[1]

		var scrapedProducts []models.Product
		if websiteName == "ksp" || websiteName == "kravitz" || websiteName == "pc365" || (websiteName == "ivory" && websiteType == "Cat") {
			if websiteName == "ivory" {
				scrapedProducts = driverConcurrentScrape(7, url, websiteName, websiteType)
			} else {
				scrapedProducts = scrape.WithDriver(url, websiteName, websiteType)
			}
		} else {
			websiteSelectors := GetWebsiteSelectors(categoryData, websiteName, websiteType)
			cleanWebsiteUrl := GetCleanWebsiteUrl(categoryData, websiteName)
			scrapedProducts = scrape.WithColly(url, cleanWebsiteUrl, websiteName, websiteType, websiteSelectors)
		}
		results <- scrapedProducts
	}

	// Step 4 :Spawn worker goroutines to scrape URLs concurrently
	for _, url := range SearchResults {
		wg.Add(1)
		go worker(url)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	// Step 5: Collect the results
	var products []models.Product
	for p := range results {
		products = append(products, p...)
	}

	//Update the trending products
	updateTrending(ctx, Category, products)

	//Get a list of WebsiteData
	vendors := GetVendorsData(categoryData, websitesVisited)

	// return the products as json
	ctx.JSON(http.StatusOK, models.QueryResult{Products: products, Vendors: vendors})
	return
}

func Demo(ctx *gin.Context) {
	time.Sleep(2 * time.Second)
	var query models.Query
	//var products []models.Product
	if err := ctx.BindJSON(&query); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"Error": "Please provide a valid query"})
		return
	}

	//Read the JSON file
	jsonData, err := os.ReadFile("../assets/products.json")
	if err != nil {
		fmt.Println(err)
	}

	//Unmarshal the JSON data
	var result models.QueryResult
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		fmt.Println(err)
	}

	// return the products as json
	ctx.JSON(http.StatusOK, result)
	return
}

// GetTrending returns the trending products for a given category
// Returns :
// 200 OK with the trending products if successful
// 404 No Found if the category is not found
// 204 No Content if the category has no trending products
func GetTrending(ctx *gin.Context) {
	category := ctx.Param("cat")
	var categoryData models.Category
	GetCategorySelectors(ctx, category, &categoryData)
	if len(categoryData.Trending) == 0 {
		ctx.JSON(http.StatusNoContent, []models.Product{})
		return
	}
	visited := CreateVisitedMap(categoryData.Trending)
	vendors := GetVendorsData(categoryData, visited)
	ctx.JSON(http.StatusOK, models.QueryResult{Products: categoryData.Trending, Vendors: vendors})
	return
}
