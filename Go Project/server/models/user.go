package models

// User is a struct that represents a user in the database
type User struct {
	Uid              string        `json:"uid" bson:"_id"`
	Name             string        `json:"name" bson:"name"`
	Email            string        `json:"email" bson:"email"`
	RecentSearches   []string      `json:"searches" bson:"searches"`
	Favorites        []Product     `json:"favorites" bson:"favorites"`
	FavoritesVendors []WebsiteData `json:"favorites_vendors" bson:"favorites_vendors"`
}

type UserUID struct {
	Uid string `json:"uid" bson:"_uid"`
}
