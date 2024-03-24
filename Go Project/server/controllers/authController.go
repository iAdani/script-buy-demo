package controllers

//this controller is for the authentication of the user

import (
	"context"
	"fmt"
	"go-server/server/database"
	"go-server/server/models"
	"google.golang.org/api/idtoken"
	"log"
	"net/http"
	"strings"

	firebase "firebase.google.com/go"

	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
)

func googleAuth(ctx *gin.Context, uid string) *models.User {
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"Error": "Please provide a valid token"})
		return nil
	}
	token := strings.Replace(authHeader, "Bearer ", "", 1)
	payload, err := idtoken.Validate(context.Background(), token, "")
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"Error": "Token is expired or invalid"})
		return nil
	}
	ctx.IndentedJSON(http.StatusOK, gin.H{"message": "User authenticated successfully"})
	return &models.User{Uid: uid, Name: payload.Claims["name"].(string), Email: payload.Claims["email"].(string), RecentSearches: []string{}, Favorites: []models.Product{}}
}

func InitAuth() (*firebase.App, error) {
	opt := option.WithCredentialsFile("C:\\CodingProjects\\script-buy\\Go Project\\server\\controllers\\scriptbuy-83438-firebase-adminsdk-ej0ez-a21e2ae32d.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, fmt.Errorf("error initializing app: %v", err)
	}
	return app, nil
}

// AuthJWT this function will check if the given uid from firebase is valid or not
func AuthJWT(app *firebase.App) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		client, err := app.Auth(ctx)
		if err != nil {
			log.Fatalf("error getting Auth client: %v\n", err)
		}
		authHeader := ctx.GetHeader("Authorization")
		log.Println("authHeader", authHeader)
		if authHeader == "" {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{"Error": "Please provide a valid token"})
			return
		}
		token := strings.Replace(authHeader, "Bearer ", "", 1)
		idToken, err := client.VerifyIDToken(ctx, token)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"Error": "Invalid token"})
		}
		log.Println("UID", idToken.UID)
		ctx.Set("uid", idToken.UID)
		ctx.Next()
	}
}

// AuthUID - Post function : /auth
func AuthUID(ctx *gin.Context) {
	var user models.UserUID
	if err := ctx.BindJSON(&user); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"Error": "Please provide a valid user"})
		return
	} else if user.Uid == "" {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"Error": "Please provide a valid user ID"})
		return
	}
	userFound := findUser(user.Uid)
	if userFound == nil {
		// logic for inserting new user into the DB :
		// 1. validate google token ID
		// 2. insert user into the DB with token payload info (uid, name, email)
		NewUser := googleAuth(ctx, user.Uid)
		db := database.Connect()
		collection := db.Client.Database("ScriptBuy").Collection("users")
		_, err := collection.InsertOne(db.Ctx, NewUser)
		if err != nil {
			ctx.IndentedJSON(http.StatusFailedDependency, gin.H{"Error": "Failed to create user"})
			return
		}
	} else {
		// UID found on the DB
		ctx.IndentedJSON(http.StatusOK, gin.H{"message": "User authenticated successfully"})
	}
}
