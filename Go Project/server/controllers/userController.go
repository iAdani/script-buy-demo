package controllers

import (
	"github.com/gin-gonic/gin"
	"go-server/server/database"
	"go-server/server/models"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
)

func findUser(uid string) *models.User {
	db := database.Connect()
	collection := db.Client.Database("ScriptBuy").Collection("users")
	var userFound models.User
	filter := bson.D{{"_id", uid}}
	err := collection.FindOne(db.Ctx, filter).Decode(&userFound)
	if err != nil {
		return nil
	} else {
		return &userFound
	}
}

// AddUser - Post function : /user
func AddUser(ctx *gin.Context) {
	var user models.User
	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"Error": "Please provide a valid User"})
		return
	}
	db := database.Connect()
	collection := db.Client.Database("ScriptBuy").Collection("users")
	_, err = collection.InsertOne(db.Ctx, user)
	if err != nil {
		ctx.IndentedJSON(http.StatusFailedDependency, gin.H{"Error": "User already exists"})
		return
	}
	ctx.IndentedJSON(http.StatusCreated, gin.H{"message": "User " + user.Name + " added successfully"})
}

// GetUser - Get function : /user/:uid
func GetUser(ctx *gin.Context) {
	uid := ctx.Param("uid")
	userFound := findUser(uid)
	if userFound == nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"Error": "User not found"})
		return
	}
	ctx.IndentedJSON(http.StatusOK, userFound)
}

// RemoveUser - Delete function : /user/:uid
func RemoveUser(ctx *gin.Context) {
	uid := ctx.Param("uid")
	if uid == "" {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"Error": "Please provide a valid UID"})
		return
	}
	db := database.Connect()
	collection := db.Client.Database("ScriptBuy").Collection("users")
	filter := bson.D{{"_id", uid}}
	_, err := collection.DeleteOne(db.Ctx, filter)
	if err != nil {
		ctx.IndentedJSON(http.StatusFailedDependency, gin.H{"Error": "User not found / User already deleted"})
		return
	}
	ctx.IndentedJSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

// this helper function will add a new vendor to the user's favorites only if the vendor is not already in the list
func addFavoriteVendor(userFavoritesVendors *[]models.WebsiteData, newVendor models.WebsiteData) {
	for _, vendor := range *userFavoritesVendors {
		if vendor.Name == newVendor.Name {
			return
		}
	}
	*userFavoritesVendors = append(*userFavoritesVendors, newVendor)
}

// GetFavorites - Get function : /user/:uid/favorites
func GetFavorites(ctx *gin.Context) {
	uid := ctx.Param("uid")
	user := findUser(uid)
	if user == nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"Error": "User not found"})
		return
	}
	ctx.IndentedJSON(http.StatusOK, models.FavoriteQuery{Products: user.Favorites, Vendors: user.FavoritesVendors})
}

// AddFavorite - Post function : /user/:uid/favorites
func AddFavorite(ctx *gin.Context) {
	uid := ctx.Param("uid")
	user := findUser(uid)
	if user == nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"Error": "User not found"})
		return
	}
	var favorite models.Favorite
	if err := ctx.BindJSON(&favorite); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"Error": "Please provide a valid favorite"})
		return
	}
	favoriteProduct := favorite.Product
	favoriteVendor := favorite.Vendor

	if favoriteProduct.Title == "" || favoriteVendor.Name == "" || favoriteProduct.Vendor != favoriteVendor.Name {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"Error": "Please provide a valid favorite"})
		return
	}

	for _, v := range user.Favorites {
		if v == favoriteProduct {
			ctx.IndentedJSON(http.StatusNoContent, gin.H{"Message": "Favorite already added"})
			return
		}
	}

	// add the product to the user's favorites
	user.Favorites = append(user.Favorites, favoriteProduct)
	// add the vendor to the user's favorites_vendors
	addFavoriteVendor(&user.FavoritesVendors, favoriteVendor)

	db := database.Connect()
	collection := db.Client.Database("ScriptBuy").Collection("users")
	filter := bson.D{{"_id", uid}}
	update := bson.D{
		{"$set", bson.D{
			{"favorites", user.Favorites},
			{"favorites_vendors", user.FavoritesVendors},
		}},
	}
	_, err := collection.UpdateOne(db.Ctx, filter, update)
	if err != nil {
		ctx.IndentedJSON(http.StatusFailedDependency, gin.H{"Error": "Favorite not added"})
		return
	}
	ctx.IndentedJSON(http.StatusCreated, gin.H{"message": "Favorite added successfully"})
}

// this helper function will remove a vendor from the user's favorites only if the vendor is in the list and
// the user has no more favorites products from this vendor in his favorites
func removeFavoriteVendor(userFavorites []models.Product, userFavoritesVendors *[]models.WebsiteData, vendorToRemove models.WebsiteData) {
	// check if there is a product from this vendor in the user's favorites
	for _, product := range userFavorites {
		if product.Vendor == vendorToRemove.Name {
			return
		}
	}
	// if there is no product from this vendor in the user's favorites, remove the vendor from the user's favorites_vendors
	for i, vendor := range *userFavoritesVendors {
		if vendor.Name == vendorToRemove.Name {
			*userFavoritesVendors = append((*userFavoritesVendors)[:i], (*userFavoritesVendors)[i+1:]...)
			break
		}
	}
}

// RemoveFavorite - Delete function : /user/:uid/favorites
func RemoveFavorite(ctx *gin.Context) {
	uid := ctx.Param("uid")
	user := findUser(uid)
	if user == nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"Error": "User not found"})
		return
	}
	var favoriteToRemove models.Favorite
	if err := ctx.BindJSON(&favoriteToRemove); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"Error": "Please provide a valid favoriteToRemove"})
		return
	}

	favoriteProduct := favoriteToRemove.Product
	favoriteVendor := favoriteToRemove.Vendor

	if favoriteProduct.Title == "" || favoriteVendor.Name == "" || favoriteProduct.Vendor != favoriteVendor.Name {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"Error": "Please provide a valid favorite"})
		return
	}

	for i, v := range user.Favorites {
		if v == favoriteProduct {
			user.Favorites = append(user.Favorites[:i], user.Favorites[i+1:]...)
			removeFavoriteVendor(user.Favorites, &user.FavoritesVendors, favoriteToRemove.Vendor)
			db := database.Connect()
			collection := db.Client.Database("ScriptBuy").Collection("users")
			filter := bson.D{{"_id", uid}}
			update := bson.D{
				{"$set", bson.D{
					{"favorites", user.Favorites},
					{"favorites_vendors", user.FavoritesVendors},
				}},
			}
			_, err := collection.UpdateOne(db.Ctx, filter, update)
			if err != nil {
				ctx.IndentedJSON(http.StatusFailedDependency, gin.H{"Error": "Favorite not removed"})
				return
			}
			ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Favorite removed successfully"})
			return
		}
	}
	ctx.IndentedJSON(http.StatusNotFound, gin.H{"Error": "Favorite not found"})
}
