package controllers

import (
	"backend/database"
	"backend/helpers"
	"backend/models"
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "User")
var validate = validator.New()

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}

	return string(bytes)
}

func VerifyPassword(userPassword string, providedPassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(userPassword))
	check := true
	msg := ""

	if err != nil {
		msg = fmt.Sprintf("login or password is incorrect")
		check = false
	}

	return check, msg
}

func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var user models.User

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validationErr := validate.Struct(user)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		count, err := userCollection.CountDocuments(ctx, bson.M{"username": user.Username})
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while checking for the Username"})
			log.Panic(err)
			return
		}

		password := HashPassword(*user.Password)
		user.Password = &password

		if count > 0 {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "this Username already exists"})
			return
		}

		user.ID = primitive.NewObjectID()
		token, refreshToken, _ := helpers.GenerateAllTokens(*user.Email, *user.Name, *user.Username, *user.UserType)
		user.Group = nil
		user.Token = &token
		user.RefreshToken = &refreshToken

		resultInsertionNumber, insertErr := userCollection.InsertOne(ctx, user)
		if insertErr != nil {
			msg := fmt.Sprintf("User item was not created")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		defer cancel()

		c.JSON(http.StatusOK, resultInsertionNumber)

	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var user models.User
		var foundUser models.User

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err := userCollection.FindOne(ctx, bson.M{"username": user.Username}).Decode(&foundUser)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "login or password is incorrect"})
			return
		}

		passwordIsValid, msg := VerifyPassword(*user.Password, *foundUser.Password)
		defer cancel()
		if passwordIsValid != true {
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		if foundUser.Username == nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "user not found"})
			return
		}
		token, refreshToken, _ := helpers.GenerateAllTokens(*foundUser.Email, *foundUser.Name, *foundUser.Username, *foundUser.UserType)

		helpers.UpdateAllTokens(token, refreshToken, *foundUser.Username)
		err = userCollection.FindOne(ctx, bson.M{"username": foundUser.Username}).Decode(&foundUser)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, foundUser)

	}
}

func GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := helpers.CheckUserType(c, "ADMIN"); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var role = c.Param("type")
		defer cancel()

		var users []models.User
		cursor, err := userCollection.Find(ctx, bson.M{"usertype": strings.ToUpper(role)})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		for cursor.Next(ctx) {
			var user models.User
			err := cursor.Decode(&user)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			users = append(users, user)
		}
		if len(users) > 0 {
			c.JSON(http.StatusOK, users)
		} else {
			c.JSON(http.StatusOK, "No users of this type available!")
		}

	}
}

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := helpers.CheckUserType(c, "ADMIN"); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		username := c.Param("username")

		if err := helpers.MatchUserTypeToUid(c, username); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var user models.User

		err := userCollection.FindOne(ctx, bson.M{"username": username}).Decode(&user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, user)

	}
}

func UpdateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := helpers.CheckUserType(c, "ADMIN"); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		username := c.Param("username")
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var user models.User

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var foundUser models.User
		err := userCollection.FindOne(ctx, bson.M{"username": username}).Decode(&foundUser)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "user was not found"})
			return
		}

		validationErr := validate.Struct(user)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}
		password := HashPassword(*user.Password)
		user.Password = &password
		count, err := userCollection.CountDocuments(ctx, bson.M{"username": user.Username})
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if *user.Username != *foundUser.Username && count > 0 {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "this Username already exists"})
			return
		}
		update := bson.M{"username": user.Username, "email": user.Email, "password": user.Password, "usertype": user.UserType, "name": user.Name, "profiledescription": user.ProfileDescription}
		result, err := userCollection.UpdateOne(ctx, bson.M{"username": username}, bson.M{"$set": update})

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		var updatedUser models.User
		if result.MatchedCount == 1 {
			err := userCollection.FindOne(ctx, bson.M{"username": user.Username}).Decode(&updatedUser)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
		}
		c.JSON(http.StatusOK, updatedUser)
	}
}

func DeleteUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := helpers.CheckUserType(c, "ADMIN"); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		username := c.Param("username")
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		result, err := userCollection.DeleteOne(ctx, bson.M{"username": username})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		if result.DeletedCount < 1 {
			c.JSON(http.StatusNotFound, gin.H{"error": "User was not found!"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"success": "User was deleted!"})
	}
}

func GenerateUser(name string, CNP string, c *gin.Context) models.User {

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	var generatedUser models.User

	generatedUser.Name = &name

	generatedUser.ID = primitive.NewObjectID()

	var role = c.Param("type")
	generatedUser.UserType = &role

	var username = strings.ReplaceAll(name, " ", "") + "." + CNP
	generatedUser.Username = &username

	var email = username + "@sleepybears.com"
	generatedUser.Email = &email

	var password = "password" + CNP[7:]
	generatedUser.Password = &password
	validationErr := validate.Struct(&generatedUser)
	if validationErr != nil {
		var badUser models.User
		var error = "USER VALIDATION ERROR"
		badUser.Name = &error
		return badUser
	}
	hashedPassword := HashPassword(password)
	generatedUser.Password = &hashedPassword

	token, refreshToken, _ := helpers.GenerateAllTokens(*generatedUser.Email, *generatedUser.Name, *generatedUser.Username, *generatedUser.UserType)
	generatedUser.Token = &token
	generatedUser.RefreshToken = &refreshToken

	count, err := userCollection.CountDocuments(ctx, bson.M{"username": generatedUser.Username})
	if err != nil {
		var badUser models.User
		var error = "Error making user w/ " + CNP
		badUser.Name = &error
		return badUser
	}

	if count > 0 {
		var badUser models.User
		var error = "CNP ALREADY EXISTS"
		badUser.Name = &error
		return badUser
	}

	_, insertErr := userCollection.InsertOne(ctx, generatedUser)
	if insertErr != nil {
		var badUser models.User
		var error = "USER GENERATION ERROR"
		badUser.Name = &error
		return badUser
	}
	generatedUser.Password = &password
	return generatedUser
}

func GenerateUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := helpers.CheckUserType(c, "ADMIN"); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var userDTOList models.UserDTOList
		if err := c.BindJSON(&userDTOList); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var generatedUsers []models.User
		for _, user := range userDTOList.UserDTOs {
			generatedUsers = append(generatedUsers, GenerateUser(*user.Name, *user.CNP, c))
		}
		c.JSON(http.StatusOK, generatedUsers)
	}
}
