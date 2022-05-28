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

  "github.com/gin-gonic/gin"
  "go.mongodb.org/mongo-driver/bson"
  "go.mongodb.org/mongo-driver/bson/primitive"
)

var groupCollection = database.OpenCollection(database.Client, "Group")

func GetGroupsByAcademicYear() gin.HandlerFunc {
  return func(c *gin.Context) {
    var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
    defer cancel()

    var groups []models.Group
    yearId := c.Param("academic_year_id")
    realYearId, err := primitive.ObjectIDFromHex(yearId)
    if err != nil {
      c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
      return
    }
    cursor, err := groupCollection.Find(ctx, bson.M{"academicyear._id": realYearId})
    if err != nil {
      c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
      return
    }
    for cursor.Next(ctx) {
      var group models.Group
      err := cursor.Decode(&group)
      if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
      }
      groups = append(groups, group)
    }
    if len(groups) > 0 {
      c.JSON(http.StatusOK, groups)
    } else {
      c.JSON(http.StatusOK, "No groups available!")
    }
  }
}

func AddGroup() gin.HandlerFunc {
  return func(c *gin.Context) {
    if err := helpers.CheckUserType(c, "ADMIN"); err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
      return
    }
    var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
    var group models.Group
    defer cancel()

    if err := c.BindJSON(&group); err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
      return
    }

    academicYear := getCurrentAcademicYear()
    group.AcademicYear = &academicYear

    validationErr := validate.Struct(group)
    if validationErr != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
      return
    }

    count, err := groupCollection.CountDocuments(ctx, bson.M{"number": group.Number, "academicyear._id": academicYear.ID})
    defer cancel()
    if err != nil {
      c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while checking for the group number"})
      log.Panic(err)
      return
    }

    if count > 0 {
      c.JSON(http.StatusInternalServerError, gin.H{"error": "this group number already exists"})
      return
    }

    group.ID = primitive.NewObjectID()

    resultInsertionNumber, insertErr := groupCollection.InsertOne(ctx, group)
    if insertErr != nil {
      msg := fmt.Sprintf("Group was not created")
      c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
      return
    }
    c.JSON(http.StatusOK, resultInsertionNumber)
  }
}

func AddStudentToGroup() gin.HandlerFunc {
  return func(c *gin.Context) {
    if err := helpers.CheckUserType(c, "ADMIN"); err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
      return
    }
    var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
    defer cancel()

    groupNumber, err := strconv.Atoi(c.Param("groupnumber"))
    if err != nil {
      c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
      return
    }
    username := c.Param("username")

    var student models.User
    err = userCollection.FindOne(ctx, bson.M{"username": username}).Decode(&student)
    if err != nil {
      c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
      return
    }
    allowedType := "STUDENT"
    if *student.UserType != allowedType {
      c.JSON(http.StatusInternalServerError, gin.H{"error": "User is not of student type"})
      return
    }
    var group models.Group
    err = groupCollection.FindOne(ctx, bson.M{"number": groupNumber}).Decode(&group)
    if err != nil {
      c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
      return
    }
    student.Group = &group
    update := bson.M{"group": student.Group}
    result, err := userCollection.UpdateOne(ctx, bson.M{"username": student.Username}, bson.M{"$set": update})

    _, err = enrollmentCollection.UpdateMany(ctx, bson.M{"user._id": student.ID}, bson.M{"$set": bson.M{"user.group": group}})

    if err != nil {
      c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
      return
    }

    if err != nil {
      c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
      return
    }
    var updatedUser models.User
    if result.MatchedCount == 1 {
      err := userCollection.FindOne(ctx, bson.M{"username": student.Username}).Decode(&updatedUser)
      if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
      }
    }
    c.JSON(http.StatusOK, updatedUser)
  }
}

func getAllGroupsForStatistics() []models.Group {
  var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
  defer cancel()

  var groups []models.Group
  cursor, err := groupCollection.Find(ctx, bson.M{})
  if err != nil {
    return []models.Group{}
  }
  for cursor.Next(ctx) {
    var group models.Group
    err := cursor.Decode(&group)
    if err != nil {
      fmt.Println("error for some reason")
      return []models.Group{}
    }
    groups = append(groups, group)
  }
  if len(groups) > 0 {
    return groups
  } else {
    return []models.Group{}
  }
}
