package models

type Product struct {
	Title  string `json:"title" bson:"title"`
	Vendor string `json:"vendor" bson:"vendor"`
	Price  string `json:"price" bson:"price"`
	Img    string `json:"img" bson:"img"`
	Link   string `json:"link" bson:"link"`
}

type Favorite struct {
	Product Product     `json:"product" bson:"product"`
	Vendor  WebsiteData `json:"vendor" bson:"vendor"`
}

type FavoriteQuery struct {
	Products []Product     `json:"products" bson:"products"`
	Vendors  []WebsiteData `json:"vendors" bson:"vendors"`
}
