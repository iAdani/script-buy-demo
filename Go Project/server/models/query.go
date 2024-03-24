package models

type Query struct {
	Query string `json:"query" bson:"query"`
}

type QueryResult struct {
	Products []Product     `json:"products" bson:"products"`
	Vendors  []WebsiteData `json:"vendors" bson:"vendors"`
}
