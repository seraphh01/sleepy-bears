package controllers

import (
  "backend/database"
  "backend/helpers"
  "backend/models"
  "context"
  "fmt"
  "log"
  "net/http"
  "strconv"
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
    user.UserId = *user.Username
    token, refreshToken, _ := helpers.GenerateAllTokens(*user.Email, *user.Name, *user.Username, *user.UserType, *&user.UserId)
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
    fmt.Println(*foundUser.Username)
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
    token, refreshToken, _ := helpers.GenerateAllTokens(*foundUser.Email, *foundUser.Name, *foundUser.Username, *foundUser.UserType, foundUser.UserId)

    helpers.UpdateAllTokens(token, refreshToken, foundUser.UserId)
    err = userCollection.FindOne(ctx, bson.M{"userid": foundUser.UserId}).Decode(&foundUser)

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

    // recordPerPage := 10
    recordPerPage, err := strconv.Atoi(c.Query("recordPerPage"))
    if err != nil || recordPerPage < 1 {
      recordPerPage = 10
    }

    page, err1 := strconv.Atoi(c.Query("page"))
    if err1 != nil || page < 1 {
      page = 1
    }

    startIndex := (page - 1) * recordPerPage
    startIndex, err = strconv.Atoi(c.Query("startIndex"))

    matchStage := bson.D{{"$match", bson.D{{}}}}
    groupStage := bson.D{{"$group", bson.D{{"_id", bson.D{{"_id", "null"}}}, {"total_count", bson.D{{"$sum", 1}}}, {"data", bson.D{{"$push", "$$ROOT"}}}}}}
    projectStage := bson.D{
      {"$project", bson.D{
        {"_id", 0},
        {"total_count", 1},
        {"user_items", bson.D{{"$slice", []interface{}{"$data", startIndex, recordPerPage}}}},
      }}}

    result, err := userCollection.Aggregate(ctx, mongo.Pipeline{
      matchStage, groupStage, projectStage})
    defer cancel()
    if err != nil {
      c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while listing user items"})
    }
    var allusers []bson.M
    if err = result.All(ctx, &allusers); err != nil {
      log.Fatal(err)
    }
    c.JSON(http.StatusOK, allusers[0])

  }
}

//GetUser is the api used to get a single user
func GetUser() gin.HandlerFunc {
  return func(c *gin.Context) {
    userId := c.Param("userid")

    if err := helpers.MatchUserTypeToUid(c, userId); err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
      return
    }
    var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

    var user models.User

    err := userCollection.FindOne(ctx, bson.M{"userid": userId}).Decode(&user)
    defer cancel()
    if err != nil {
      c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
      return
    }

    c.JSON(http.StatusOK, user)

  }
}
