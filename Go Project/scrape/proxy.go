package scrape

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func main2() {
	// Call the GetProxyList function to retrieve the list of proxies
	proxies := GetProxyList()

	//// Print the list of proxies
	//for _, proxy := range proxies {
	//	fmt.Println(proxy)
	//}

	// randomly shuffle the proxies
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(proxies), func(i, j int) { proxies[i], proxies[j] = proxies[j], proxies[i] })

	target := "https://www.pc365.co.il/product-23068-Dell_UltraSharp_25_USB_C_Monitor_U2520D_3Y_%D7%91%D7%9E%D7%9C%D7%90%D7%99"

	// Loop through the proxies and make a request
	for _, proxy := range proxies {
		// Create a new HTTP client and set a timeout
		client := &http.Client{
			Timeout: 5 * time.Second,
		}

		// Create a new request using http
		request, err := http.NewRequest("GET", target, nil)
		if err != nil {
			fmt.Println(err)
			continue
		}

		// Set the proxy URL
		proxyUrl, err := url.Parse("https://" + proxy)
		if err != nil {
			fmt.Println(err)
			continue
		}

		// Set the proxy URL in the client
		client.Transport = &http.Transport{
			Proxy: http.ProxyURL(proxyUrl),
		}

		// Make the request to the URL
		response, err := client.Do(request)
		if err != nil {
			fmt.Println(err)
			continue
		}

		// Close the response body
		response.Body.Close()

		// Check the status code
		if response.StatusCode == 200 {
			fmt.Println("Success with proxy: " + proxy)
			break
		}
	}

}

func GetProxyList() []string {
	MainUrl := "https://free-proxy-list.net/"

	response, err := http.Get(MainUrl)
	if err != nil {
		panic(err)
	}

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		panic(err)
	}

	var proxies []string

	doc.Find("tbody tr").Each(func(i int, s *goquery.Selection) {
		isHttps := strings.Contains(s.Find("td:nth-child(7)").Text(), "yes")
		if isHttps {
			ip := s.Find("td:nth-child(1)").Text()
			port := s.Find("td:nth-child(2)").Text()
			proxy := ip + ":" + port
			proxies = append(proxies, proxy)
		}
	})

	return proxies[:20]
}
