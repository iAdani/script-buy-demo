package scrape

import (
	"context"
	"fmt"
	"os"
	"strconv"

	g "github.com/serpapi/google-search-results-golang"
	"google.golang.org/api/customsearch/v1"
	"google.golang.org/api/option"
)

// GetResultsGoogleCustom This function will preform a Google Search with the given query and number of results and return the list of results
// using the Google Custom Search API
func GetResultsGoogleCustom(query string, numOfResults int) []string {
	ApiKey := os.Getenv("CUSTOM_SEARCH_API_KEY")
	SearchEngineID := os.Getenv("CUSTOM_SEARCH_ENGINE_ID")
	ctx := context.Background()
	customSearch, err := customsearch.NewService(ctx, option.WithAPIKey(ApiKey))
	if err != nil {
		panic(err)
	}
	cse := customsearch.NewCseService(customSearch)
	var output []string
	var numOfSet int
	left := numOfResults
	start := 0
	// Google Custom Search API allows only 10 results per request, so we need to send multiple requests
	// to get the desired number of results
	for i := 0; i < (numOfResults/10)+1; i++ {
		if left > 10 {
			numOfSet = 10
		} else {
			numOfSet = left
		}

		do, err := cse.List().Cx(SearchEngineID).Cr("il").Gl("il").Q(query).Num(int64(numOfSet)).Start(int64(start)).Do()
		if err != nil {
			panic(err)
		}
		for _, item := range do.Items {
			output = append(output, item.Link)
		}

		if i == 0 {
			start += 1
		}
		left = left - numOfSet
		start = start + 10
	}
	return output
}

// GetResultsSerpAPI This function will preform a Google Search with the given query and number of results and return the list of results
// using SerpApi API
func GetResultsSerpAPI(query string, numOfResults int) []string {
	//Create GoogleSearch object with your SERP API key
	parameter := map[string]string{
		"api_key":       os.Getenv("SERP_API_KEY"),
		"engine":        "google",
		"q":             query,
		"location":      "Israel",
		"google_domain": "google.co.il",
		"gl":            "il",
		"hl":            "iw",
		"num":           strconv.Itoa(numOfResults),
	}
	search := g.NewGoogleSearch(parameter, "5cb20f0ce5ddc69addec943627d813df45289b5436b8b9b60fad95d5dd436949")
	//// Run search with the specified query
	results, err := search.GetJSON()
	if err != nil {
		fmt.Printf("%s/n", results)
	}

	//results := map[string]interface{}{
	//	"search_metadata": map[string]interface{}{
	//		"id":               "63ac166c34ff95ca7aa8f914",
	//		"status":           "Success",
	//		"json_endpoint":    "https://serpapi.com/searches/3c9275752f45755f/63ac166c34ff95ca7aa8f914.json",
	//		"created_at":       "2022-12-28 10:11:56 UTC",
	//		"processed_at":     "2022-12-28 10:11:57 UTC",
	//		"google_url":       "https://www.google.co.il/search?q=Mx+Master+3&oq=Mx+Master+3&uule=w+CAIQICIGSXNyYWVs&hl=iw&gl=il&num=15&sourceid=chrome&ie=UTF-8",
	//		"raw_html_file":    "https://serpapi.com/searches/3c9275752f45755f/63ac166c34ff95ca7aa8f914.html",
	//		"total_time_taken": 3.02,
	//	},
	//	"search_parameters": map[string]interface{}{
	//		"engine":             "google",
	//		"q":                  "Mx Master 3",
	//		"location_requested": "Israel",
	//		"location_used":      "Israel",
	//		"google_domain":      "google.co.il",
	//		"hl":                 "iw",
	//		"gl":                 "il",
	//		"num":                "15",
	//		"device":             "desktop",
	//	},
	//	"search_information": map[string]interface{}{
	//		"organic_results_state": "Results for exact spelling",
	//		"query_displayed":       "Mx Master 3",
	//		"total_results":         184000000,
	//		"time_taken_displayed":  0.43,
	//		"menu_items": []interface{}{
	//			map[string]interface{}{
	//				"position": 1,
	//				"title":    "הכול",
	//			},
	//			map[string]interface{}{
	//				"position":     2,
	//				"title":        "תמונות",
	//				"link":         "https://www.google.co.il/search?q=Mx+Master+3&num=15&ucbcb=1&gl=il&hl=iw&source=lnms&tbm=isch&sa=X&ved=2ahUKEwi2mrTKiZz8AhUkuJUCHRcRDGcQ_AUoAXoECAEQAw",
	//				"serpapi_link": "https://serpapi.com/search.json?device=desktop&engine=google&gl=il&google_domain=google.co.il&hl=iw&location=Israel&num=15&q=Mx+Master+3&tbm=isch",
	//			},
	//			map[string]interface{}{
	//				"position":     3,
	//				"title":        "שופינג",
	//				"link":         "https://www.google.co.il/search?q=Mx+Master+3&num=15&ucbcb=1&gl=il&hl=iw&source=lnms&tbm=shop&sa=X&ved=2ahUKEwi2mrTKiZz8AhUkuJUCHRcRDGcQ_AUoAnoECAEQBA",
	//				"serpapi_link": "https://serpapi.com/search.json?device=desktop&engine=google&gl=il&google_domain=google.co.il&hl=iw&location=Israel&num=15&q=Mx+Master+3&tbm=shop",
	//			},
	//			map[string]interface{}{
	//				"position":     4,
	//				"title":        "סרטונים",
	//				"link":         "https://www.google.co.il/search?q=Mx+Master+3&num=15&ucbcb=1&gl=il&hl=iw&source=lnms&tbm=vid&sa=X&ved=2ahUKEwi2mrTKiZz8AhUkuJUCHRcRDGcQ_AUoA3oECAEQBQ",
	//				"serpapi_link": "https://serpapi.com/search.json?device=desktop&engine=google&gl=il&google_domain=google.co.il&hl=iw&location=Israel&num=15&q=Mx+Master+3&tbm=vid",
	//			},
	//			map[string]interface{}{
	//				"position":     5,
	//				"title":        "חדשות",
	//				"link":         "https://www.google.co.il/search?q=Mx+Master+3&num=15&ucbcb=1&gl=il&hl=iw&source=lnms&tbm=nws&sa=X&ved=2ahUKEwi2mrTKiZz8AhUkuJUCHRcRDGcQ_AUoBHoECAEQBg",
	//				"serpapi_link": "https://serpapi.com/search.json?device=desktop&engine=google&gl=il&google_domain=google.co.il&hl=iw&location=Israel&num=15&q=Mx+Master+3&tbm=nws",
	//			},
	//		},
	//	},
	//	"ads": []interface{}{
	//		map[string]interface{}{
	//			"position":       1,
	//			"block_position": "top",
	//			"title":          "מה ההבדל בין העכבר MX Master 3 ל- MX Master 3 for Mac |...",
	//			"link":           "https://www.adcs.co.il/difference-between-mx-master-3-and-mx-master-for-mac",
	//			"displayed_link": "https://www.adcs.co.il/",
	//			"tracking_link":  "https://www.google.co.il/aclk?sa=l&ai=DChcSEwiIj7rKiZz8AhVEwpEKHcTsAKUYABAFGgJjZQ&sig=AOD64_3H4uHX_uxseAzBRAto3Wu90y1DAg&q&adurl",
	//			"description":    "בלעדי באמירים! 3 שנות אחריות על מחשבי המקבוק פרו, כולל שירות איסוף והחזרה ומחשב חלופי. מבצעי מקבוק פרו במפרטים גבוהים במחירים סופר משתלמים. מגוון רחב בזמינות מיידית.",
	//			"sitelinks": []interface{}{
	//				map[string]interface{}{
	//					"title": "מבצעים",
	//					"link":  "https://www.adcs.co.il/deals",
	//					"snippets": []interface{}{
	//						"מחשבים, מסכים, טאבלטים וציוד היקפי המוצרים שנמצאים במבצע באתר אמירים",
	//					},
	//				},
	//				map[string]interface{}{
	//					"title": "המחלקה העסקית של אמירים",
	//					"link":  "https://www.adcs.co.il/amirim-for-business",
	//					"snippets": []interface{}{
	//						"לקוחות עסקיים מקבלים יותר אמירים נותנת לעסק שלך הטבות בלעדיות",
	//					},
	//				},
	//			},
	//		},
	//		map[string]interface{}{
	//			"position":       2,
	//			"block_position": "top",
	//			"title":          "עכבר אלחוטי Logitech MX Master 3S צבע | KSP",
	//			"link":           "https://ksp.co.il/mob/item/211980",
	//			"displayed_link": "https://www.ksp.co.il/",
	//			"tracking_link":  "https://www.google.co.il/aclk?sa=l&ai=DChcSEwiIj7rKiZz8AhVEwpEKHcTsAKUYABABGgJjZQ&sig=AOD64_2xOZVAhsZbmOvcpLF5BxZMG6eDbw&q&adurl",
	//			"extensions": []interface{}{
	//				"‏מוצרי מציאון ועודפי מלאי · ‏עולם הפארם של קיי.אס.פי · ‏עולם הבשמים של קיי.אס.פי",
	//			},
	//			"description": "קונים הכל במקום אחד! הבאנו אלפי מוצרים חדשים ומפתיעים לאתר KSP. פראייר מי שלא משווה! השוו אותנו מול כל המתחרים. זמין במלאי. המחירים הזולים ביותר. משלוח חינם לסניף הקרוב. סניפי KSP בפריסה ארצית. שירותי המקום: מחירי KSP, פריסת תשלומים נוחה, מגוון קטגוריות רחב.",
	//			"sitelinks": []interface{}{
	//				map[string]interface{}{
	//					"title": "מוצרי מציאון ועודפי מלאי",
	//					"link":  "https://ksp.co.il/web/outlet?utm_source=google&utm_medium=cpc&utm_campaign=dsa_dynamic_search_display_all_segments&utm_content=st",
	//				},
	//				map[string]interface{}{
	//					"title": "עולם הפארם של קיי.אס.פי",
	//					"link":  "https://ksp.co.il/web/world/13?utm_source=google&utm_medium=cpc&utm_campaign=dsa_dynamic_search_display_all_segments&utm_content=st",
	//				},
	//				map[string]interface{}{
	//					"title": "עולם הבשמים של קיי.אס.פי",
	//					"link":  "https://ksp.co.il/web/world/5?utm_source=google&utm_medium=cpc&utm_campaign=dsa_dynamic_search_display_all_segments&utm_content=st",
	//				},
	//			},
	//		},
	//		map[string]interface{}{
	//			"position":       3,
	//			"block_position": "top",
	//			"title":          "עכבר אלחוטי LogiTech MX Master 3 לוגיטק - זאפ השוואת מחירים",
	//			"link":           "https://www.zap.co.il/model.aspx?modelid=1046492",
	//			"displayed_link": "https://www.zap.co.il/",
	//			"tracking_link":  "https://www.google.co.il/aclk?sa=l&ai=DChcSEwiIj7rKiZz8AhVEwpEKHcTsAKUYABAAGgJjZQ&sig=AOD64_1CkeQwH3ntfKAneLAkAVp8NA6jDQ&q&adurl",
	//			"description":    "כל המחירים וחוות הדעת במקום אחד! היכנסו לזאפ השוואת מחירים עכשיו. חוות דעת על חנויות. חוות דעת על מוצרים. המחירים הכי אטרקטיביים. השוואת מחירי דגמים וסוגים.",
	//		},
	//	},
	//	"shopping_results": []interface{}{
	//		map[string]interface{}{
	//			"position":        1,
	//			"block_position":  "right",
	//			"title":           "עכבר אלחוטי MX Master 2S של ...",
	//			"price":           "‏245.68 ‏₪ + מס",
	//			"extracted_price": 245.68,
	//			"link":            "https://www.amazon.com/-/he/%D7%A2%D7%9B%D7%91%D7%A8-%D7%90%D7%9C%D7%97%D7%95%D7%98%D7%99-Master-%D7%9C%D7%95%D7%92%D7%99%D7%98%D7%A7-%D7%90%D7%A8%D7%92%D7%95%D7%A0%D7%95%D7%9E%D7%99%D7%AA/dp/B071YZJ1G1?source=ps-sl-shoppingads-lpcontext&psc=1&smid=A3ULK995NAZ7K2",
	//			"source":          "Amazon.com",
	//			"thumbnail":       "https://serpapi.com/searches/63ac166c34ff95ca7aa8f914/images/86eb6b893e16986314d9510a53873bc8eb0360fbfc2048c08ceb7764ceee3a5d.png",
	//		},
	//		map[string]interface{}{
	//			"position":        2,
	//			"block_position":  "right",
	//			"title":           "עכבר אלחוטי מקצועי Logitech MX ...",
	//			"price":           "‏399.00 ‏₪",
	//			"extracted_price": 399.0,
	//			"link":            "https://shipi.co.il/product/%D7%A2%D7%9B%D7%91%D7%A8-%D7%90%D7%9C%D7%97%D7%95%D7%98%D7%99-%D7%9E%D7%A7%D7%A6%D7%95%D7%A2%D7%99?utm_source=Google+Shopping&utm_medium=cpc&utm_campaign=feed_il_new",
	//			"source":          "שיפי - רשת סלולר",
	//			"thumbnail":       "https://encrypted-tbn0.gstatic.com/shopping?q=tbn:ANd9GcSHT2njhBIT-K82rrVctDJqygwZnORqXN8nSo45YfcG_0Ki6KG_cresCheSwWoGiXe1d9zyrYUGH6Q&usqp=CAc",
	//		},
	//		map[string]interface{}{
	//			"position":        3,
	//			"block_position":  "right",
	//			"title":           "‏עכבר ‏אלחוטי LogiTech MX ...",
	//			"price":           "‏269.00 ‏₪",
	//			"extracted_price": 269.0,
	//			"link":            "https://www.pompa.co.il/product/%E2%80%8F%D7%A2%D7%9B%D7%91%D7%A8-%E2%80%8F%D7%90%D7%9C%D7%97%D7%95%D7%98%D7%99-logitech-mx-master-2s-%D7%91%D7%9E%D7%9C%D7%90%D7%99-0",
	//			"source":          "אלקטריק פומפה",
	//			"thumbnail":       "https://serpapi.com/searches/63ac166c34ff95ca7aa8f914/images/86eb6b893e16986314d9510a53873bc82d582043bd51b5dd1fb9ff5fbaa1137a.png",
	//		},
	//		map[string]interface{}{
	//			"position":        4,
	//			"block_position":  "right",
	//			"title":           "עכבר אלחוטי לוג׳יטק Logitech ...",
	//			"price":           "‏499.00 ‏₪",
	//			"extracted_price": 499.0,
	//			"link":            "https://www.adcs.co.il/logitech-mx-master-3s-performance-wireless-mouse-graphite-910-006559.html",
	//			"source":          "אמירים הפצה",
	//			"thumbnail":       "https://encrypted-tbn3.gstatic.com/shopping?q=tbn:ANd9GcQ05r-1lgFXc0peQtrX0zP6UA3yXbTUDL3ZXxUBlHJLKHsTjaN7ofV_7OQ0emGYWcgO56W3eM8BfQ&usqp=CAc",
	//		},
	//	},
	//	"inline_videos": []interface{}{
	//		map[string]interface{}{
	//			"position":  1,
	//			"title":     "העכבר הטוב ביותר ליוצרי תוכן וגרפיקאים - Logitech MX Master 3",
	//			"link":      "https://www.youtube.com/watch?v=7pYgia7Wb9I",
	//			"thumbnail": "https://i.ytimg.com/vi/7pYgia7Wb9I/mqdefault.jpg?sqp=-oaymwEECHwQRg&rs=AMzJL3m8ExHXeZT5u_4SCfr-i8nGUMoJEw",
	//			"channel":   "מהיר ומחשובי",
	//			"duration":  "5:40",
	//			"platform":  "YouTube",
	//			"date":      "18 בדצמ׳ 2021",
	//		},
	//		map[string]interface{}{
	//			"position":  2,
	//			"title":     "?! העכבר שתרצו ב- 2021| logitech mx master 3 סקירה",
	//			"link":      "https://www.youtube.com/watch?v=lGlJdsAneuc",
	//			"thumbnail": "https://i.ytimg.com/vi/lGlJdsAneuc/mqdefault.jpg?sqp=-oaymwEECHwQRg&rs=AMzJL3mpJbAtXfH2u7AlNcLBiX9ijH61VQ",
	//			"channel":   "Yam Olisker",
	//			"duration":  "10:08",
	//			"platform":  "YouTube",
	//			"date":      "24 בפבר׳ 2021",
	//		},
	//		map[string]interface{}{
	//			"position":  3,
	//			"title":     "Logitech MX Master 3 Review - a Productivity Beast!",
	//			"link":      "https://www.youtube.com/watch?v=c2vJ7cS3Sh8",
	//			"thumbnail": "https://i.ytimg.com/vi/c2vJ7cS3Sh8/mqdefault.jpg?sqp=-oaymwEECHwQRg&rs=AMzJL3mnc9U1RgE-ERgMEADsfZ5R21vSgg",
	//			"channel":   "Created Tech",
	//			"duration":  "11:54",
	//			"platform":  "YouTube",
	//			"date":      "17 בינו׳ 2022",
	//			"key_moments": []interface{}{
	//				map[string]interface{}{
	//					"time":      "00:00",
	//					"title":     "MX Master 3 Review",
	//					"link":      "https://www.youtube.com/watch?v=c2vJ7cS3Sh8&t=0",
	//					"thumbnail": "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTWWSHXs8gjgKBhF1uA_w-3hd3XGpm4hjSiXxqClfcOXg&s",
	//				},
	//				map[string]interface{}{
	//					"time":      "00:48",
	//					"title":     "Form Factor",
	//					"link":      "https://www.youtube.com/watch?v=c2vJ7cS3Sh8&t=48",
	//					"thumbnail": "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRVr9fHdQOkUkF-85dOOHSeCcn69HbFVUImFQJpr5QA_g&s",
	//				},
	//				map[string]interface{}{
	//					"time":      "02:54",
	//					"title":     "Usability",
	//					"link":      "https://www.youtube.com/watch?v=c2vJ7cS3Sh8&t=174",
	//					"thumbnail": "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTohkA9mhem6SQhHQ7fB6KgQY6pJKJ0efiBee3uBRaj6g&s",
	//				},
	//				map[string]interface{}{
	//					"time":      "04:39",
	//					"title":     "Software and Shortcuts",
	//					"link":      "https://www.youtube.com/watch?v=c2vJ7cS3Sh8&t=279",
	//					"thumbnail": "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRW4baE4eK7yhxSIRAYXtCxF2XU5AStQHk3cN3AeMFhQQ&s",
	//				},
	//				map[string]interface{}{
	//					"time":      "06:31",
	//					"title":     "MagSpeed Scroll Wheel",
	//					"link":      "https://www.youtube.com/watch?v=c2vJ7cS3Sh8&t=391",
	//					"thumbnail": "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcT82TPeiDwriyM_73OmBzzZSQWEujb1nIXJsDsn4wqI6Q&s",
	//				},
	//				map[string]interface{}{
	//					"time":      "07:40",
	//					"title":     "Charging and Battery Life",
	//					"link":      "https://www.youtube.com/watch?v=c2vJ7cS3Sh8&t=460",
	//					"thumbnail": "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcR3jsVe69Byr8bGjQ_d2LiyqBir7nFTlraJRyw9z6kmDg&s",
	//				},
	//				map[string]interface{}{
	//					"time":      "09:17",
	//					"title":     "Pricing",
	//					"link":      "https://www.youtube.com/watch?v=c2vJ7cS3Sh8&t=557",
	//					"thumbnail": "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRxETSTaYs5rqCaIUEP5Gpq9zLp2C-T7sQzRKhpf0o-Og&s",
	//				},
	//				map[string]interface{}{
	//					"time":      "10:07",
	//					"title":     "MX Master 3 vs MX Master 3 For Mac",
	//					"link":      "https://www.youtube.com/watch?v=c2vJ7cS3Sh8&t=607",
	//					"thumbnail": "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRSQdByAyBs9uBFi66uE1iBgW5ZizcUm6CIXB6eRCXl7g&s",
	//				},
	//				map[string]interface{}{
	//					"time":      "10:21",
	//					"title":     "MX Master 3 vs MX Master 2S",
	//					"link":      "https://www.youtube.com/watch?v=c2vJ7cS3Sh8&t=621",
	//					"thumbnail": "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTNvnO9JdHfMNWC3ryoOS-xsR4qCIox5qVVtG4a9Cundg&s",
	//				},
	//				map[string]interface{}{
	//					"time":      "11:14",
	//					"title":     "Should You Buy the MX Master 3?",
	//					"link":      "https://www.youtube.com/watch?v=c2vJ7cS3Sh8&t=674",
	//					"thumbnail": "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTB2V0zQzRD16oHqL29c4v-j-7xxr-uy_HdfcUh0dOTAA&s",
	//				},
	//			},
	//		},
	//	},
	//	"organic_results": []interface{}{
	//		map[string]interface{}{
	//			"position":       1,
	//			"title":          "עכבר אלחוטי Logitech MX Master 3S צבע | KSP",
	//			"link":           "https://ksp.co.il/mob/item/211980",
	//			"displayed_link": "https://ksp.co.il › מחשבים וסלולר › עכבר › Logitech",
	//			"snippet":        "עכבר אלחוטי MX Master 3S מבית Logitech בגרסה מחודשת לעכבר האייקוני עם חיישן שעובד על כל משטח, כולל לחצנים שקטים, גלגלת שקטה ומהירה, מאפשר עבודה רציפה בין ...",
	//			"snippet_highlighted_words": []interface{}{
	//				"MX Master 3S",
	//			},
	//			"cached_page_link": "https://webcache.googleusercontent.com/search?q=cache:BQGnX37EYBAJ:https://ksp.co.il/mob/item/211980&cd=2&hl=iw&ct=clnk&gl=il",
	//			"source":           "ksp.co.il",
	//		},
	//		map[string]interface{}{
	//			"position":       2,
	//			"title":          "השוואת מחירים ל‏עכבר ‏אלחוטי LogiTech MX Master 3S לוגיטק",
	//			"link":           "https://www.zap.co.il/model.aspx?modelid=1157073",
	//			"displayed_link": "https://www.zap.co.il › ... › השוואת מחירים עכברים",
	//			"snippet":        "MX Master 3S לוגיטק החל מ - 363‏₪. רק בזאפ תמצאו השוואת מחירים של ‏עכבר ‏אלחוטי LogiTech MX Master 3S לוגיטק, מידע על חנויות קרובות, 4 חוות דעת, מפרט טכני ועוד.",
	//			"snippet_highlighted_words": []interface{}{
	//				"MX Master 3",
	//				"3",
	//			},
	//			"cached_page_link": "https://webcache.googleusercontent.com/search?q=cache:_PFFX1cJNC8J:https://ksp.co.il/mob/item/72907&cd=2&hl=iw&ct=clnk&gl=il",
	//		},
	//		map[string]interface{}{
	//			"position":       3,
	//			"title":          "MX Master 3S Wireless Performance Mouse - Logitech",
	//			"link":           "https://www.logitech.com/en-us/products/mice/mx-master-3s.html",
	//			"displayed_link": "https://www.logitech.com › products › mice",
	//			"snippet":        "Meet MX Master 3S – an iconic mouse remastered. Feel every moment of your workflow with even more precision, tactility, and performance, thanks to Quiet Clicks ...",
	//			"snippet_highlighted_words": []interface{}{
	//				"MX Master",
	//			},
	//			"cached_page_link": "https://webcache.googleusercontent.com/search?q=cache:BQGnX37EYBAJ:https://ksp.co.il/mob/item/211980&cd=3&hl=iw&ct=clnk&gl=il",
	//		},
	//		map[string]interface{}{
	//			"position":       4,
	//			"title":          "MX Master 3S Wireless Mouse | Logitech Europe",
	//			"link":           "https://www.logitech.com/en-eu/products/mice/mx-master-3s.html",
	//			"displayed_link": "https://www.logitech.com › products › mice",
	//			"snippet":        "Shop MX Master 3S Wireless Mouse. Features precision tracking, MagSpeed scroll and thumb wheel, app customization, flow between devices, and more.",
	//			"snippet_highlighted_words": []interface{}{
	//				"MX Master",
	//				"three",
	//			},
	//			"rich_snippet": map[string]interface{}{
	//				"top": map[string]interface{}{
	//					"detected_extensions": map[string]interface{}{
	//						"דירוג":    4.2,
	//						"‏_סקירות": 236,
	//						"‏‏_‏":     99.99,
	//					},
	//					"extensions": []interface{}{
	//						"דירוג: 4.2",
	//						"‏236 סקירות",
	//						"‏‏99.99 ‏$",
	//						"‏במלאי",
	//					},
	//				},
	//			},
	//		},
	//		map[string]interface{}{
	//			"position":       5,
	//			"title":          "עכבר אלחוטי Logitech Mx master 3s Wireless Bluetooth ... - ...",
	//			"link":           "https://www.ivory.co.il/catalog.php?id=45275",
	//			"displayed_link": "https://www.ivory.co.il › עכברים",
	//			"snippet":        "- MX Master 3 חיי סוללה בטעינה מלאה עד 70 ימים בתנאים אופטימליים. - הטענה מהירה- בדקה תקבל שלוש שעות של שימוש - כבל הטעינה USB-C כלול באריזה. - התאם את ה- MX ...",
	//			"snippet_highlighted_words": []interface{}{
	//				"MX Master 3",
	//				"MX",
	//			},
	//			"cached_page_link": "https://webcache.googleusercontent.com/search?q=cache:wrRQ7H7ZJxAJ:https://www.bug.co.il/brand/logitech/mouse/color/darkgray/mx/master/3&cd=5&hl=iw&ct=clnk&gl=il",
	//		},
	//		map[string]interface{}{
	//			"position":       6,
	//			"title":          "עכבר אלחוטי Logitech Mx master 3s Wireless Bluetooth בצבע ...",
	//			"link":           "https://www.ivory.co.il/catalog.php?id=45278",
	//			"displayed_link": "https://www.ivory.co.il › עכברים",
	//			"snippet":        "היכנסו לרכישת עכבר אלחוטי MX Master 3S Logitech ברשת באג, באג מציעה לכם את מוצרים מהמותגים המובילים עם אחריות מלאה במחירים משתלמים ומשלוח מהיר עד הבית מהיום ...",
	//			"snippet_highlighted_words": []interface{}{
	//				"MX Master",
	//			},
	//			"cached_page_link": "https://webcache.googleusercontent.com/search?q=cache:KmNlTwMUtlcJ:https://www.bug.co.il/brand/logitech/mx/master/3s/white&cd=6&hl=iw&ct=clnk&gl=il",
	//		},
	//		map[string]interface{}{
	//			"position":       7,
	//			"title":          "עכבר ‏אלחוטי LogiTech MX Master 3S לוגיטק - ALL In Cell",
	//			"link":           "https://allincell.co.il/product/%E2%80%8F%D7%A2%D7%9B%D7%91%D7%A8-%E2%80%8F%D7%90%D7%9C%D7%97%D7%95%D7%98%D7%99-logitech-mx-master-3s-%D7%9C%D7%95%D7%92%D7%99%D7%98%D7%A7/",
	//			"displayed_link": "https://allincell.co.il › למשרד › עכברים",
	//			"snippet":        "עכבר Logitech MX Master 3 לוגיטק - לפני הקנייה השווה מחירים וקרא סקירות מומחים, מפרטים טכניים וחוות דעת ב- Wisebuy.",
	//			"snippet_highlighted_words": []interface{}{
	//				"MX Master 3",
	//			},
	//			"rich_snippet": map[string]interface{}{
	//				"top": map[string]interface{}{
	//					"detected_extensions": map[string]interface{}{
	//						"דירוג":         5,
	//						"‏_הצבעות":      27,
	//						"‏‏_‏₪_עד_‏_‏₪": 310.0,
	//					},
	//					"extensions": []interface{}{
	//						"דירוג: 5",
	//						"‏27 הצבעות",
	//						"‏‏310.00 ‏₪ עד ‏536.00 ‏₪",
	//					},
	//				},
	//				"bottom": map[string]interface{}{
	//					"extensions": []interface{}{
	//						"סוג: עכבר",
	//						"חיבור: אלחוטי",
	//						"סוגי אלחוט: Bluetooth",
	//						"לחצנים: גלגלת צדית, 6 לחצנים",
	//					},
	//					"detected_extensions": map[string]interface{}{
	//						"לחצנים_גלגלת_צדית_לחצנים": nil,
	//					},
	//				},
	//			},
	//			"cached_page_link": "https://webcache.googleusercontent.com/search?q=cache:Ye9taKcjIrkJ:https://www.wisebuy.co.il/product.aspx%3Fcategory%3Dc-mouse%26productid%3D1046492&cd=7&hl=iw&ct=clnk&gl=il",
	//		},
	//		map[string]interface{}{
	//			"position":       8,
	//			"title":          "Amazon.com: Logitech MX Master 3S - Graphite : Electronics",
	//			"link":           "https://www.amazon.com/Logitech-MX-Master-3S-Graphite/dp/B09HM94VDS",
	//			"displayed_link": "https://www.amazon.com › Logitech-MX-Master-3S-Gra...",
	//			"snippet":        "Mx master 3 is the most advanced master series mouse yet. It has been designed for designers and engineered for coders – to create, make, and do faster and more ...",
	//			"snippet_highlighted_words": []interface{}{
	//				"Mx master 3",
	//			},
	//			"rich_snippet": map[string]interface{}{
	//				"top": map[string]interface{}{
	//					"detected_extensions": map[string]interface{}{
	//						"דירוג":    4.7,
	//						"‏_סקירות": 15353,
	//					},
	//					"extensions": []interface{}{
	//						"דירוג: 4.7",
	//						"‏15,353 סקירות",
	//					},
	//				},
	//				"bottom": map[string]interface{}{
	//					"detected_extensions": map[string]interface{}{
	//						"series_logitech_mx_master_advanced_wirel": 3,
	//						"operating_system_windows_or_later_li":     10,
	//					},
	//					"extensions": []interface{}{
	//						"Series: ‎Logitech MX Master 3 Advanced Wirel...",
	//						"Compatible Devices: ‎Laptop, Personal Compu...",
	//						"Operating System: ‎Windows 10, 11 or later, Li...",
	//						"Recommended Uses For Product: ‎Laptop",
	//					},
	//				},
	//			},
	//		},
	//		map[string]interface{}{
	//			"position":       9,
	//			"title":          "Logitech MX Master 3S - השוואת מחירים וסקירות מומחים - ...",
	//			"link":           "https://www.wisebuy.co.il/product.aspx?category=c-mouse&productid=1157073",
	//			"displayed_link": "https://www.wisebuy.co.il › מחשבים ותוכנות › עכברים",
	//			"snippet":        "היכנסו לקניית עכבר אלחוטי Logitech MX Master 3 Wireless Bluetooth בצבע שחור ברשת אייבורי מחשבים וסלולר המציעה את מיטב המותגים במחירים משתלמים.",
	//			"snippet_highlighted_words": []interface{}{
	//				"MX Master 3",
	//			},
	//		},
	//		map[string]interface{}{
	//			"position":         10,
	//			"title":            "עכבר אלחוטי לוגיטק שחור logitech mx master 3s wireless mouse",
	//			"link":             "https://www.espir.co.il/product/%D7%A2%D7%9B%D7%91%D7%A8-%D7%90%D7%9C%D7%97%D7%95%D7%98%D7%99-%D7%9C%D7%95%D7%92%D7%99%D7%98%D7%A7-logitech-mx-master-3s-91000-655-90",
	//			"displayed_link":   "https://www.espir.co.il › ... › עכברי Logitech",
	//			"snippet":          "יצרן: Logitech, סוג מוצר: עכבר, חיבור: אלחוטי / Bluetooth, תקופת האחריות: שנתיים, נותן השירות: היבואן הרשמי.",
	//			"cached_page_link": "https://webcache.googleusercontent.com/search?q=cache:pEWPD_vE7OUJ:https://www.pc365.co.il/product-26931-%25D7%25A2%25D7%259B%25D7%2591%25D7%25A8_%25D7%2590%25D7%259C%25D7%2597%25D7%2595%25D7%2598%25D7%2599_Logitech_MX_Master_3_Bluetooth_Mid_Grey&cd=13&hl=iw&ct=clnk&gl=il",
	//		},
	//		map[string]interface{}{
	//			"position":       11,
	//			"title":          "עכבר אלחוטי Logitech - MX Master 3S Performance 91000 ...",
	//			"link":           "https://www.pc365.co.il/product-33166-%D7%A2%D7%9B%D7%91%D7%A8_%D7%90%D7%9C%D7%97%D7%95%D7%98%D7%99_Logitech_MX_Master_3S_Performance_Pale_Gray_%D7%91%D7%9E%D7%9C%D7%90%D7%99",
	//			"displayed_link": "https://www.pc365.co.il › ציוד היקפי › עכברים",
	//			"snippet":        "עכבר אלחוטי Logitech MX Master 3 המתקדם ביותר של סדרת המאסטר. עם גלילה אלקטרומגנטית חדשה של ™ MagSpeed; מדויקת ומהירה מספיק כדי לגלול 1,000 שורות בשנייה ...",
	//			"snippet_highlighted_words": []interface{}{
	//				"MX Master 3",
	//			},
	//			"cached_page_link": "https://webcache.googleusercontent.com/search?q=cache:-U3AmR8I8NYJ:https://www.benda.co.il/product/logitech-mx-master-3-black/&cd=14&hl=iw&ct=clnk&gl=il",
	//		},
	//		map[string]interface{}{
	//			"position":       12,
	//			"title":          "Logitech עכבר אלחוטי MX Master 3S For Mac - באג",
	//			"link":           "https://www.bug.co.il/brand/logitech/mx/master/3s/for/mac/space/grey",
	//			"displayed_link": "https://www.bug.co.il › מקלדות ועכברים",
	//			"snippet":        "עכבר אלחוטי לוגיטק שחור Logitech MX Master 3S Wireless Mouse. ( חוות דעת גולשים) ... Logitech MX Master 3 MX Master 3 Logitech MX Master 3.",
	//			"snippet_highlighted_words": []interface{}{
	//				"MX Master",
	//				"MX Master 3 MX Master 3",
	//				"MX Master 3",
	//			},
	//			"rich_snippet": map[string]interface{}{
	//				"top": map[string]interface{}{
	//					"detected_extensions": map[string]interface{}{
	//						"‏_‏₪": 465.0,
	//					},
	//					"extensions": []interface{}{
	//						"‏465.00 ‏₪",
	//						"‏במלאי",
	//					},
	//				},
	//				"bottom": map[string]interface{}{
	//					"detected_extensions": map[string]interface{}{
	//						"מק_ט_חנות": 91000,
	//						"unknown":   655,
	//						"price":     90,
	//						"currency":  "מחיר משלוח ₪ זמן אספקה עדימי עסקים",
	//					},
	//					"extensions": []interface{}{
	//						"מק\"ט חנות: 91000",
	//						"655",
	//						"90מחיר משלוח: ₪ 55זמן אספקה: עד3ימי עסקים",
	//					},
	//				},
	//			},
	//			"cached_page_link": "https://webcache.googleusercontent.com/search?q=cache:S2-C0L09dzUJ:https://www.espir.co.il/product/%25D7%25A2%25D7%259B%25D7%2591%25D7%25A8-%25D7%2590%25D7%259C%25D7%2597%25D7%2595%25D7%2598%25D7%2599-%25D7%259C%25D7%2595%25D7%2592%25D7%2599%25D7%2598%25D7%25A7-logitech-mx-master-3s-91000-655-90&cd=15&hl=iw&ct=clnk&gl=il",
	//		},
	//		map[string]interface{}{
	//			"position":       13,
	//			"title":          "עכבר אלחוטי Logitech MX Master 3S - עכברים למחשב - King ...",
	//			"link":           "https://www.king-games.co.il/products/item/15983",
	//			"displayed_link": "https://www.king-games.co.il › ... › עכברים למחשב",
	//			"snippet":        "MX Master 3 סוג אלחוטי ממשק. USB Bluetooth שונות מאפשר לנווט בין שני מחשבים ולהעתיק ולהדביק קבצים מאחד לשני ערך נומינלי: 1000DPI",
	//			"snippet_highlighted_words": []interface{}{
	//				"MX Master 3",
	//			},
	//			"cached_page_link": "https://webcache.googleusercontent.com/search?q=cache:vLpXuMYLO1MJ:https://allincell.co.il/product/%25E2%2580%258F%25D7%25A2%25D7%259B%25D7%2591%25D7%25A8-%25E2%2580%258F%25D7%2590%25D7%259C%25D7%2597%25D7%2595%25D7%2598%25D7%2599-logitech-mx-master-3-%25D7%259C%25D7%2595%25D7%2592%25D7%2599%25D7%2598%25D7%25A7/&cd=16&hl=iw&ct=clnk&gl=il",
	//		},
	//		map[string]interface{}{
	//			"position":       14,
	//			"title":          "עכבר אלחוטי לוג׳יטק Logitech MX Master 3S for Mac ...",
	//			"link":           "https://www.adcs.co.il/mx-master-3s-for-mac-performance-wireless-mouse-pale-gray-910-006572.html",
	//			"displayed_link": "https://www.adcs.co.il › mx-master-3s-for-mac-perfor...",
	//			"snippet":        "עכבר אלחוטי Logitech MX Master 3 מק\"ט: 910-005647 אחריות: שלוש שנים יצרן: Logitech קישור לאתר יצרן.",
	//			"snippet_highlighted_words": []interface{}{
	//				"MX Master 3",
	//			},
	//			"rich_snippet": map[string]interface{}{
	//				"top": map[string]interface{}{
	//					"detected_extensions": map[string]interface{}{
	//						"דירוג": 5,
	//					},
	//					"extensions": []interface{}{
	//						"דירוג: 5",
	//						"‏הצבעה אחת",
	//					},
	//				},
	//			},
	//			"cached_page_link": "https://webcache.googleusercontent.com/search?q=cache:JoeLYEh440IJ:https://1pc.co.il/he/product-100376-%25D7%25A2%25D7%259B%25D7%2591%25D7%25A8_%25D7%2590%25D7%259C%25D7%2597%25D7%2595%25D7%2598%25D7%2599_logitech_mx_master_3_%25D7%259C%25D7%2595%25D7%2592%25D7%2599%25D7%2598%25D7%25A7&cd=17&hl=iw&ct=clnk&gl=il",
	//		},
	//		map[string]interface{}{
	//			"position": 14,
	//			"title":    "עכבר אלחוטי לוג׳יטק Logitech MX Master 3S for Mac ...",
	//			"link":     "https://www.adcs.co.il/mx-master-3s-for-mac-performance-wireless-mouse-pale-gray-910-006572.html",
	//		},
	//	},
	//	"related_searches": []interface{}{
	//		map[string]interface{}{
	//			"query": "Logitech mx master 3 זאפ",
	//			"link":  "https://www.google.co.il/search?num=15&ucbcb=1&gl=il&hl=iw&q=Logitech+mx+master+3+%D7%96%D7%90%D7%A4&sa=X&ved=2ahUKEwi2mrTKiZz8AhUkuJUCHRcRDGcQ1QJ6BAgoEAE",
	//		},
	//		map[string]interface{}{
	//			"query": "Logitech mx master 3 מחיר",
	//			"link":  "https://www.google.co.il/search?num=15&ucbcb=1&gl=il&hl=iw&q=Logitech+mx+master+3+%D7%9E%D7%97%D7%99%D7%A8&sa=X&ved=2ahUKEwi2mrTKiZz8AhUkuJUCHRcRDGcQ1QJ6BAgmEAE",
	//		},
	//		map[string]interface{}{
	//			"query": "עכבר logitech mx master 3",
	//			"link":  "https://www.google.co.il/search?num=15&ucbcb=1&gl=il&hl=iw&q=%D7%A2%D7%9B%D7%91%D7%A8+logitech+mx+master+3&sa=X&ved=2ahUKEwi2mrTKiZz8AhUkuJUCHRcRDGcQ1QJ6BAglEAE",
	//		},
	//		map[string]interface{}{
	//			"query": "Mx master 2s",
	//			"link":  "https://www.google.co.il/search?num=15&ucbcb=1&gl=il&hl=iw&q=Mx+master+2s&sa=X&ved=2ahUKEwi2mrTKiZz8AhUkuJUCHRcRDGcQ1QJ6BAghEAE",
	//		},
	//		map[string]interface{}{
	//			"query": "Logitech MX Master 4",
	//			"link":  "https://www.google.co.il/search?num=15&ucbcb=1&gl=il&hl=iw&q=Logitech+MX+Master+4&sa=X&ved=2ahUKEwi2mrTKiZz8AhUkuJUCHRcRDGcQ1QJ6BAgeEAE",
	//		},
	//		map[string]interface{}{
	//			"query": "Mx master 3 bug",
	//			"link":  "https://www.google.co.il/search?num=15&ucbcb=1&gl=il&hl=iw&q=Mx+master+3+bug&sa=X&ved=2ahUKEwi2mrTKiZz8AhUkuJUCHRcRDGcQ1QJ6BAgdEAE",
	//		},
	//		map[string]interface{}{
	//			"query": "Logitech MX Master 3 driver",
	//			"link":  "https://www.google.co.il/search?num=15&ucbcb=1&gl=il&hl=iw&q=Logitech+MX+Master+3+driver&sa=X&ved=2ahUKEwi2mrTKiZz8AhUkuJUCHRcRDGcQ1QJ6BAgbEAE",
	//		},
	//		map[string]interface{}{
	//			"query": "Logitech MX Master 3 Amazon",
	//			"link":  "https://www.google.co.il/search?num=15&ucbcb=1&gl=il&hl=iw&q=Logitech+MX+Master+3+Amazon&sa=X&ved=2ahUKEwi2mrTKiZz8AhUkuJUCHRcRDGcQ1QJ6BAgaEAE",
	//		},
	//	},
	//	"pagination": map[string]interface{}{
	//		"current": 1,
	//		"next":    "https://www.google.co.il/search?q=Mx+Master+3&num=15&ucbcb=1&gl=il&hl=iw&ei=bhasY7arNKTw1sQPl6KwuAY&start=15&sa=N&ved=2ahUKEwi2mrTKiZz8AhUkuJUCHRcRDGcQ8NMDegQICBAW",
	//		"other_pages": map[string]interface{}{
	//			"2":  "https://www.google.co.il/search?q=Mx+Master+3&num=15&ucbcb=1&gl=il&hl=iw&ei=bhasY7arNKTw1sQPl6KwuAY&start=15&sa=N&ved=2ahUKEwi2mrTKiZz8AhUkuJUCHRcRDGcQ8tMDegQICBAE",
	//			"3":  "https://www.google.co.il/search?q=Mx+Master+3&num=15&ucbcb=1&gl=il&hl=iw&ei=bhasY7arNKTw1sQPl6KwuAY&start=30&sa=N&ved=2ahUKEwi2mrTKiZz8AhUkuJUCHRcRDGcQ8tMDegQICBAG",
	//			"4":  "https://www.google.co.il/search?q=Mx+Master+3&num=15&ucbcb=1&gl=il&hl=iw&ei=bhasY7arNKTw1sQPl6KwuAY&start=45&sa=N&ved=2ahUKEwi2mrTKiZz8AhUkuJUCHRcRDGcQ8tMDegQICBAI",
	//			"5":  "https://www.google.co.il/search?q=Mx+Master+3&num=15&ucbcb=1&gl=il&hl=iw&ei=bhasY7arNKTw1sQPl6KwuAY&start=60&sa=N&ved=2ahUKEwi2mrTKiZz8AhUkuJUCHRcRDGcQ8tMDegQICBAK",
	//			"6":  "https://www.google.co.il/search?q=Mx+Master+3&num=15&ucbcb=1&gl=il&hl=iw&ei=bhasY7arNKTw1sQPl6KwuAY&start=75&sa=N&ved=2ahUKEwi2mrTKiZz8AhUkuJUCHRcRDGcQ8tMDegQICBAM",
	//			"7":  "https://www.google.co.il/search?q=Mx+Master+3&num=15&ucbcb=1&gl=il&hl=iw&ei=bhasY7arNKTw1sQPl6KwuAY&start=90&sa=N&ved=2ahUKEwi2mrTKiZz8AhUkuJUCHRcRDGcQ8tMDegQICBAO",
	//			"8":  "https://www.google.co.il/search?q=Mx+Master+3&num=15&ucbcb=1&gl=il&hl=iw&ei=bhasY7arNKTw1sQPl6KwuAY&start=105&sa=N&ved=2ahUKEwi2mrTKiZz8AhUkuJUCHRcRDGcQ8tMDegQICBAQ",
	//			"9":  "https://www.google.co.il/search?q=Mx+Master+3&num=15&ucbcb=1&gl=il&hl=iw&ei=bhasY7arNKTw1sQPl6KwuAY&start=120&sa=N&ved=2ahUKEwi2mrTKiZz8AhUkuJUCHRcRDGcQ8tMDegQICBAS",
	//			"10": "https://www.google.co.il/search?q=Mx+Master+3&num=15&ucbcb=1&gl=il&hl=iw&ei=bhasY7arNKTw1sQPl6KwuAY&start=135&sa=N&ved=2ahUKEwi2mrTKiZz8AhUkuJUCHRcRDGcQ8tMDegQICBAU",
	//		},
	//	},
	//	"serpapi_pagination": map[string]interface{}{
	//		"current":   1,
	//		"next_link": "https://serpapi.com/search.json?device=desktop&engine=google&gl=il&google_domain=google.co.il&hl=iw&location=Israel&num=15&q=Mx+Master+3&start=15",
	//		"next":      "https://serpapi.com/search.json?device=desktop&engine=google&gl=il&google_domain=google.co.il&hl=iw&location=Israel&num=15&q=Mx+Master+3&start=15",
	//		"other_pages": map[string]interface{}{
	//			"2":  "https://serpapi.com/search.json?device=desktop&engine=google&gl=il&google_domain=google.co.il&hl=iw&location=Israel&num=15&q=Mx+Master+3&start=15",
	//			"3":  "https://serpapi.com/search.json?device=desktop&engine=google&gl=il&google_domain=google.co.il&hl=iw&location=Israel&num=15&q=Mx+Master+3&start=30",
	//			"4":  "https://serpapi.com/search.json?device=desktop&engine=google&gl=il&google_domain=google.co.il&hl=iw&location=Israel&num=15&q=Mx+Master+3&start=45",
	//			"5":  "https://serpapi.com/search.json?device=desktop&engine=google&gl=il&google_domain=google.co.il&hl=iw&location=Israel&num=15&q=Mx+Master+3&start=60",
	//			"6":  "https://serpapi.com/search.json?device=desktop&engine=google&gl=il&google_domain=google.co.il&hl=iw&location=Israel&num=15&q=Mx+Master+3&start=75",
	//			"7":  "https://serpapi.com/search.json?device=desktop&engine=google&gl=il&google_domain=google.co.il&hl=iw&location=Israel&num=15&q=Mx+Master+3&start=90",
	//			"8":  "https://serpapi.com/search.json?device=desktop&engine=google&gl=il&google_domain=google.co.il&hl=iw&location=Israel&num=15&q=Mx+Master+3&start=105",
	//			"9":  "https://serpapi.com/search.json?device=desktop&engine=google&gl=il&google_domain=google.co.il&hl=iw&location=Israel&num=15&q=Mx+Master+3&start=120",
	//			"10": "https://serpapi.com/search.json?device=desktop&engine=google&gl=il&google_domain=google.co.il&hl=iw&location=Israel&num=15&q=Mx+Master+3&start=135",
	//		},
	//	},
	//}

	//results := map[string]interface{}{
	//	"organic_results": []interface{}{
	//		map[string]interface{}{
	//			"position":       1,
	//			"title":          "עכבר אלחוטי Logitech MX Master 3S צבע | KSP",
	//			"link":           "https://ksp.co.il/mob/item/211980",
	//			"displayed_link": "https://ksp.co.il › מחשבים וסלולר › עכבר › Logitech",
	//			"snippet":        "עכבר אלחוטי MX Master 3S מבית Logitech בגרסה מחודשת לעכבר האייקוני עם חיישן שעובד על כל משטח, כולל לחצנים שקטים, גלגלת שקטה ומהירה, מאפשר עבודה רציפה בין ...",
	//			"snippet_highlighted_words": []interface{}{
	//				"MX Master 3S",
	//			},
	//			"cached_page_link": "https://webcache.googleusercontent.com/search?q=cache:BQGnX37EYBAJ:https://ksp.co.il/mob/item/211980&cd=2&hl=iw&ct=clnk&gl=il",
	//			"source":           "ksp.co.il",
	//		},
	//		map[string]interface{}{
	//			"position":       2,
	//			"title":          "השוואת מחירים ל‏עכבר ‏אלחוטי LogiTech MX Master 3S לוגיטק",
	//			"link":           "https://www.zap.co.il/model.aspx?modelid=1157073",
	//			"displayed_link": "https://www.zap.co.il › ... › השוואת מחירים עכברים",
	//			"snippet":        "MX Master 3S לוגיטק החל מ - 363‏₪. רק בזאפ תמצאו השוואת מחירים של ‏עכבר ‏אלחוטי LogiTech MX Master 3S לוגיטק, מידע על חנויות קרובות, 4 חוות דעת, מפרט טכני ועוד.",
	//			"snippet_highlighted_words": []interface{}{
	//				"MX Master 3",
	//				"3",
	//			},
	//			"cached_page_link": "https://webcache.googleusercontent.com/search?q=cache:_PFFX1cJNC8J:https://ksp.co.il/mob/item/72907&cd=2&hl=iw&ct=clnk&gl=il",
	//		},
	//		map[string]interface{}{
	//			"position":       3,
	//			"title":          "MX Master 3S Wireless Performance Mouse - Logitech",
	//			"link":           "https://www.logitech.com/en-us/products/mice/mx-master-3s.html",
	//			"displayed_link": "https://www.logitech.com › products › mice",
	//			"snippet":        "Meet MX Master 3S – an iconic mouse remastered. Feel every moment of your workflow with even more precision, tactility, and performance, thanks to Quiet Clicks ...",
	//			"snippet_highlighted_words": []interface{}{
	//				"MX Master",
	//			},
	//			"cached_page_link": "https://webcache.googleusercontent.com/search?q=cache:BQGnX37EYBAJ:https://ksp.co.il/mob/item/211980&cd=3&hl=iw&ct=clnk&gl=il",
	//		},
	//		map[string]interface{}{
	//			"position":       4,
	//			"title":          "MX Master 3S Wireless Mouse | Logitech Europe",
	//			"link":           "https://www.logitech.com/en-eu/products/mice/mx-master-3s.html",
	//			"displayed_link": "https://www.logitech.com › products › mice",
	//			"snippet":        "Shop MX Master 3S Wireless Mouse. Features precision tracking, MagSpeed scroll and thumb wheel, app customization, flow between devices, and more.",
	//			"snippet_highlighted_words": []interface{}{
	//				"MX Master",
	//				"three",
	//			},
	//			"rich_snippet": map[string]interface{}{
	//				"top": map[string]interface{}{
	//					"detected_extensions": map[string]interface{}{
	//						"דירוג":    4.2,
	//						"‏_סקירות": 236,
	//						"‏‏_‏":     99.99,
	//					},
	//					"extensions": []interface{}{
	//						"דירוג: 4.2",
	//						"‏236 סקירות",
	//						"‏‏99.99 ‏$",
	//						"‏במלאי",
	//					},
	//				},
	//			},
	//		},
	//		map[string]interface{}{
	//			"position":       5,
	//			"title":          "עכבר אלחוטי Logitech Mx master 3s Wireless Bluetooth ... - ...",
	//			"link":           "https://www.ivory.co.il/catalog.php?id=45275",
	//			"displayed_link": "https://www.ivory.co.il › עכברים",
	//			"snippet":        "- MX Master 3 חיי סוללה בטעינה מלאה עד 70 ימים בתנאים אופטימליים. - הטענה מהירה- בדקה תקבל שלוש שעות של שימוש - כבל הטעינה USB-C כלול באריזה. - התאם את ה- MX ...",
	//			"snippet_highlighted_words": []interface{}{
	//				"MX Master 3",
	//				"MX",
	//			},
	//			"cached_page_link": "https://webcache.googleusercontent.com/search?q=cache:wrRQ7H7ZJxAJ:https://www.bug.co.il/brand/logitech/mouse/color/darkgray/mx/master/3&cd=5&hl=iw&ct=clnk&gl=il",
	//		},
	//		map[string]interface{}{
	//			"position":       6,
	//			"title":          "עכבר אלחוטי Logitech Mx master 3s Wireless Bluetooth בצבע ...",
	//			"link":           "https://www.ivory.co.il/catalog.php?id=45278",
	//			"displayed_link": "https://www.ivory.co.il › עכברים",
	//			"snippet":        "היכנסו לרכישת עכבר אלחוטי MX Master 3S Logitech ברשת באג, באג מציעה לכם את מוצרים מהמותגים המובילים עם אחריות מלאה במחירים משתלמים ומשלוח מהיר עד הבית מהיום ...",
	//			"snippet_highlighted_words": []interface{}{
	//				"MX Master",
	//			},
	//			"cached_page_link": "https://webcache.googleusercontent.com/search?q=cache:KmNlTwMUtlcJ:https://www.bug.co.il/brand/logitech/mx/master/3s/white&cd=6&hl=iw&ct=clnk&gl=il",
	//		},
	//		map[string]interface{}{
	//			"position":       7,
	//			"title":          "עכבר ‏אלחוטי LogiTech MX Master 3S לוגיטק - ALL In Cell",
	//			"link":           "https://allincell.co.il/product/%E2%80%8F%D7%A2%D7%9B%D7%91%D7%A8-%E2%80%8F%D7%90%D7%9C%D7%97%D7%95%D7%98%D7%99-logitech-mx-master-3s-%D7%9C%D7%95%D7%92%D7%99%D7%98%D7%A7/",
	//			"displayed_link": "https://allincell.co.il › למשרד › עכברים",
	//			"snippet":        "עכבר Logitech MX Master 3 לוגיטק - לפני הקנייה השווה מחירים וקרא סקירות מומחים, מפרטים טכניים וחוות דעת ב- Wisebuy.",
	//			"snippet_highlighted_words": []interface{}{
	//				"MX Master 3",
	//			},
	//			"rich_snippet": map[string]interface{}{
	//				"top": map[string]interface{}{
	//					"detected_extensions": map[string]interface{}{
	//						"דירוג":         5,
	//						"‏_הצבעות":      27,
	//						"‏‏_‏₪_עד_‏_‏₪": 310.0,
	//					},
	//					"extensions": []interface{}{
	//						"דירוג: 5",
	//						"‏27 הצבעות",
	//						"‏‏310.00 ‏₪ עד ‏536.00 ‏₪",
	//					},
	//				},
	//				"bottom": map[string]interface{}{
	//					"extensions": []interface{}{
	//						"סוג: עכבר",
	//						"חיבור: אלחוטי",
	//						"סוגי אלחוט: Bluetooth",
	//						"לחצנים: גלגלת צדית, 6 לחצנים",
	//					},
	//					"detected_extensions": map[string]interface{}{
	//						"לחצנים_גלגלת_צדית_לחצנים": nil,
	//					},
	//				},
	//			},
	//			"cached_page_link": "https://webcache.googleusercontent.com/search?q=cache:Ye9taKcjIrkJ:https://www.wisebuy.co.il/product.aspx%3Fcategory%3Dc-mouse%26productid%3D1046492&cd=7&hl=iw&ct=clnk&gl=il",
	//		},
	//		map[string]interface{}{
	//			"position":       8,
	//			"title":          "Amazon.com: Logitech MX Master 3S - Graphite : Electronics",
	//			"link":           "https://www.amazon.com/Logitech-MX-Master-3S-Graphite/dp/B09HM94VDS",
	//			"displayed_link": "https://www.amazon.com › Logitech-MX-Master-3S-Gra...",
	//			"snippet":        "Mx master 3 is the most advanced master series mouse yet. It has been designed for designers and engineered for coders – to create, make, and do faster and more ...",
	//			"snippet_highlighted_words": []interface{}{
	//				"Mx master 3",
	//			},
	//			"rich_snippet": map[string]interface{}{
	//				"top": map[string]interface{}{
	//					"detected_extensions": map[string]interface{}{
	//						"דירוג":    4.7,
	//						"‏_סקירות": 15353,
	//					},
	//					"extensions": []interface{}{
	//						"דירוג: 4.7",
	//						"‏15,353 סקירות",
	//					},
	//				},
	//				"bottom": map[string]interface{}{
	//					"detected_extensions": map[string]interface{}{
	//						"series_logitech_mx_master_advanced_wirel": 3,
	//						"operating_system_windows_or_later_li":     10,
	//					},
	//					"extensions": []interface{}{
	//						"Series: ‎Logitech MX Master 3 Advanced Wirel...",
	//						"Compatible Devices: ‎Laptop, Personal Compu...",
	//						"Operating System: ‎Windows 10, 11 or later, Li...",
	//						"Recommended Uses For Product: ‎Laptop",
	//					},
	//				},
	//			},
	//		},
	//		map[string]interface{}{
	//			"position":       9,
	//			"title":          "Logitech MX Master 3S - השוואת מחירים וסקירות מומחים - ...",
	//			"link":           "https://www.wisebuy.co.il/product.aspx?category=c-mouse&productid=1157073",
	//			"displayed_link": "https://www.wisebuy.co.il › מחשבים ותוכנות › עכברים",
	//			"snippet":        "היכנסו לקניית עכבר אלחוטי Logitech MX Master 3 Wireless Bluetooth בצבע שחור ברשת אייבורי מחשבים וסלולר המציעה את מיטב המותגים במחירים משתלמים.",
	//			"snippet_highlighted_words": []interface{}{
	//				"MX Master 3",
	//			},
	//		},
	//		map[string]interface{}{
	//			"position":         10,
	//			"title":            "עכבר אלחוטי לוגיטק שחור logitech mx master 3s wireless mouse",
	//			"link":             "https://www.espir.co.il/product/%D7%A2%D7%9B%D7%91%D7%A8-%D7%90%D7%9C%D7%97%D7%95%D7%98%D7%99-%D7%9C%D7%95%D7%92%D7%99%D7%98%D7%A7-logitech-mx-master-3s-91000-655-90",
	//			"displayed_link":   "https://www.espir.co.il › ... › עכברי Logitech",
	//			"snippet":          "יצרן: Logitech, סוג מוצר: עכבר, חיבור: אלחוטי / Bluetooth, תקופת האחריות: שנתיים, נותן השירות: היבואן הרשמי.",
	//			"cached_page_link": "https://webcache.googleusercontent.com/search?q=cache:pEWPD_vE7OUJ:https://www.pc365.co.il/product-26931-%25D7%25A2%25D7%259B%25D7%2591%25D7%25A8_%25D7%2590%25D7%259C%25D7%2597%25D7%2595%25D7%2598%25D7%2599_Logitech_MX_Master_3_Bluetooth_Mid_Grey&cd=13&hl=iw&ct=clnk&gl=il",
	//		},
	//		map[string]interface{}{
	//			"position":       11,
	//			"title":          "עכבר אלחוטי Logitech - MX Master 3S Performance 91000 ...",
	//			"link":           "https://www.pc365.co.il/product-33166-%D7%A2%D7%9B%D7%91%D7%A8_%D7%90%D7%9C%D7%97%D7%95%D7%98%D7%99_Logitech_MX_Master_3S_Performance_Pale_Gray_%D7%91%D7%9E%D7%9C%D7%90%D7%99",
	//			"displayed_link": "https://www.pc365.co.il › ציוד היקפי › עכברים",
	//			"snippet":        "עכבר אלחוטי Logitech MX Master 3 המתקדם ביותר של סדרת המאסטר. עם גלילה אלקטרומגנטית חדשה של ™ MagSpeed; מדויקת ומהירה מספיק כדי לגלול 1,000 שורות בשנייה ...",
	//			"snippet_highlighted_words": []interface{}{
	//				"MX Master 3",
	//			},
	//			"cached_page_link": "https://webcache.googleusercontent.com/search?q=cache:-U3AmR8I8NYJ:https://www.benda.co.il/product/logitech-mx-master-3-black/&cd=14&hl=iw&ct=clnk&gl=il",
	//		},
	//		map[string]interface{}{
	//			"position":       12,
	//			"title":          "Logitech עכבר אלחוטי MX Master 3S For Mac - באג",
	//			"link":           "https://www.bug.co.il/brand/logitech/mx/master/3s/for/mac/space/grey",
	//			"displayed_link": "https://www.bug.co.il › מקלדות ועכברים",
	//			"snippet":        "עכבר אלחוטי לוגיטק שחור Logitech MX Master 3S Wireless Mouse. ( חוות דעת גולשים) ... Logitech MX Master 3 MX Master 3 Logitech MX Master 3.",
	//			"snippet_highlighted_words": []interface{}{
	//				"MX Master",
	//				"MX Master 3 MX Master 3",
	//				"MX Master 3",
	//			},
	//			"rich_snippet": map[string]interface{}{
	//				"top": map[string]interface{}{
	//					"detected_extensions": map[string]interface{}{
	//						"‏_‏₪": 465.0,
	//					},
	//					"extensions": []interface{}{
	//						"‏465.00 ‏₪",
	//						"‏במלאי",
	//					},
	//				},
	//				"bottom": map[string]interface{}{
	//					"detected_extensions": map[string]interface{}{
	//						"מק_ט_חנות": 91000,
	//						"unknown":   655,
	//						"price":     90,
	//						"currency":  "מחיר משלוח ₪ זמן אספקה עדימי עסקים",
	//					},
	//					"extensions": []interface{}{
	//						"מק\"ט חנות: 91000",
	//						"655",
	//						"90מחיר משלוח: ₪ 55זמן אספקה: עד3ימי עסקים",
	//					},
	//				},
	//			},
	//			"cached_page_link": "https://webcache.googleusercontent.com/search?q=cache:S2-C0L09dzUJ:https://www.espir.co.il/product/%25D7%25A2%25D7%259B%25D7%2591%25D7%25A8-%25D7%2590%25D7%259C%25D7%2597%25D7%2595%25D7%2598%25D7%2599-%25D7%259C%25D7%2595%25D7%2592%25D7%2599%25D7%2598%25D7%25A7-logitech-mx-master-3s-91000-655-90&cd=15&hl=iw&ct=clnk&gl=il",
	//		},
	//		map[string]interface{}{
	//			"position":       13,
	//			"title":          "עכבר אלחוטי Logitech MX Master 3S - עכברים למחשב - King ...",
	//			"link":           "https://www.king-games.co.il/products/item/15983",
	//			"displayed_link": "https://www.king-games.co.il › ... › עכברים למחשב",
	//			"snippet":        "MX Master 3 סוג אלחוטי ממשק. USB Bluetooth שונות מאפשר לנווט בין שני מחשבים ולהעתיק ולהדביק קבצים מאחד לשני ערך נומינלי: 1000DPI",
	//			"snippet_highlighted_words": []interface{}{
	//				"MX Master 3",
	//			},
	//			"cached_page_link": "https://webcache.googleusercontent.com/search?q=cache:vLpXuMYLO1MJ:https://allincell.co.il/product/%25E2%2580%258F%25D7%25A2%25D7%259B%25D7%2591%25D7%25A8-%25E2%2580%258F%25D7%2590%25D7%259C%25D7%2597%25D7%2595%25D7%2598%25D7%2599-logitech-mx-master-3-%25D7%259C%25D7%2595%25D7%2592%25D7%2599%25D7%2598%25D7%25A7/&cd=16&hl=iw&ct=clnk&gl=il",
	//		},
	//		map[string]interface{}{
	//			"position":       14,
	//			"title":          "עכבר אלחוטי לוג׳יטק Logitech MX Master 3S for Mac ...",
	//			"link":           "https://www.adcs.co.il/mx-master-3s-for-mac-performance-wireless-mouse-pale-gray-910-006572.html",
	//			"displayed_link": "https://www.adcs.co.il › mx-master-3s-for-mac-perfor...",
	//			"snippet":        "עכבר אלחוטי Logitech MX Master 3 מק\"ט: 910-005647 אחריות: שלוש שנים יצרן: Logitech קישור לאתר יצרן.",
	//			"snippet_highlighted_words": []interface{}{
	//				"MX Master 3",
	//			},
	//			"rich_snippet": map[string]interface{}{
	//				"top": map[string]interface{}{
	//					"detected_extensions": map[string]interface{}{
	//						"דירוג": 5,
	//					},
	//					"extensions": []interface{}{
	//						"דירוג: 5",
	//						"‏הצבעה אחת",
	//					},
	//				},
	//			},
	//			"cached_page_link": "https://webcache.googleusercontent.com/search?q=cache:JoeLYEh440IJ:https://1pc.co.il/he/product-100376-%25D7%25A2%25D7%259B%25D7%2591%25D7%25A8_%25D7%2590%25D7%259C%25D7%2597%25D7%2595%25D7%2598%25D7%2599_logitech_mx_master_3_%25D7%259C%25D7%2595%25D7%2592%25D7%2599%25D7%2598%25D7%25A7&cd=17&hl=iw&ct=clnk&gl=il",
	//		},
	//	},
	//}

	// Output: for each result returned by the API, return the "link" field
	var output []string
	for _, result := range results["organic_results"].([]interface{}) {
		output = append(output, result.(map[string]interface{})["link"].(string))
	}
	return output
}
