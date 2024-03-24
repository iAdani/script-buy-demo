package models

type Selector struct {
	Type              string    `json:"type" bson:"type"`
	MainSelector      string    `json:"main_selector" bson:"main_selector"`
	TitleSelector     string    `json:"title_selector" bson:"title_selector"`
	PriceSelector     string    `json:"price_selector" bson:"price_selector"`
	ImageSelector     [2]string `json:"image_selector" bson:"image_selector"`
	ImgSelectorPrefix string    `json:"img_selector_prefix" bson:"img_selector_prefix"`
	UrlSelector       [2]string `json:"url_selector" bson:"url_selector"`
}

type Website struct {
	Name      string      `json:"name" bson:"_name"`
	Url       string      `json:"url" bson:"url"`
	LogoUrl   string      `json:"logo_url" bson:"logo_url"`
	Selectors [2]Selector `json:"selectors" bson:"selectors"`
}

type Category struct {
	Cat      string    `json:"cat" bson:"_cat"`
	Websites []Website `json:"websites" bson:"websites"`
	Trending []Product `json:"trending" bson:"trending"`
}

type WebsiteData struct {
	Name    string `json:"name" bson:"_name"`
	LogoUrl string `json:"logo_url" bson:"logo_url"`
}
