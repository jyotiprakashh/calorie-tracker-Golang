package routes

import (
	"context"
	"fmt"
	"net/http"
	"time"
	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin"
	"github.com/jyotiprakashh/calorie-tracker/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var validate = validator.New()

// var entryCollection *mongo.Collection = database.OpenCollection(database.Client, "entry")
var entryCollection *mongo.Collection = OpenCollection(Client, "calories")

func AddEntry(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var entry models.Entry

	if err := c.BindJSON(&entry); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	validationErr := validate.Struct(entry)
	if validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		fmt.Print("error occured while validating entry", validationErr)
		return
	}
	entry.ID = primitive.NewObjectID()
	result, insertErr := entryCollection.InsertOne(ctx, entry)
	if insertErr != nil {
		msg := fmt.Sprintf("entry item was not created")
		c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		fmt.Print(insertErr)
		return
	}
	defer cancel()
	c.JSON(http.StatusOK, result)
}

func GetEntries(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	var entries []bson.M
	cursor, err := entryCollection.Find(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Print("error occured while getting entries", err)
		return
	}

	if err = cursor.All(ctx, &entries); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Print("error occured while getting entries", err)
		return
	}
	defer cancel()
	fmt.Print(entries)
	c.JSON(http.StatusOK, entries)
}

func GetEntryById(c *gin.Context) {
	EntryID := c.Params.ByName("id")
	docID, _ := primitive.ObjectIDFromHex(EntryID)

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	var entry bson.M
	err := entryCollection.FindOne(ctx, bson.M{"_id": docID}).Decode(&entry)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Print("error occured while getting entry", err)
		return
	}
	defer cancel()
	fmt.Print(entry)
	c.JSON(http.StatusOK, entry)
}

func GetEntryByIngridient(c *gin.Context) {
	ingredients := c.Params.ByName("id")
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	var entries []bson.M

	cursor, err := entryCollection.Find(ctx, bson.M{"ingredients": ingredients})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Print("error occured while getting entries", err)
		return
	}

	if err = cursor.All(ctx, &entries); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Print("error occured while getting entries", err)
		return
	}

	defer cancel()
	fmt.Print(entries)

	c.JSON(http.StatusOK, entries)
}

func UpdateEntry(c *gin.Context) {
	entryID := c.Params.ByName("id")
	docID, _ := primitive.ObjectIDFromHex(entryID)

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	var entry models.Entry
	if err := c.BindJSON(&entry); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validationErr := validate.Struct(entry)
	if validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		fmt.Print("error occured while validating entry", validationErr)
		return
	}

	result, err := entryCollection.ReplaceOne(ctx, bson.M{"_id": docID}, bson.M{
		"dish":        entry.Dish,
		"fat":         entry.Fat,
		"ingredients": entry.Ingredients,
		"calories":    entry.Calories,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Print("error occured while updating entry", err)
		return
	}

	defer cancel()
	c.JSON(http.StatusOK, result.ModifiedCount)

}

func UpdateIngredient(c *gin.Context) {
	entryID := c.Params.ByName("id")
	docID, _ := primitive.ObjectIDFromHex(entryID)

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	type Ingredient struct {
		Ingredients *string `json:"ingredients,omitempty" `
	}

	var ingredient Ingredient

	if err := c.BindJSON(&ingredient); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Print("error occured while updating entry", err)
		return
	}
// 	result, err := entryCollection.UpdateOne(ctx, bson.M{"_id": docID}, bson.M{"_id": docID},
// 	bson.D{{"$set", bson.D{
// 		{"ingredients", ingredient.Ingredients}}}},
// )


//     if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		fmt.Print("error occured while updating entry", err)
// 		return
// 	}


update := bson.D{{"$set", bson.D{{"ingredients", ingredient.Ingredients}}}}

options := options.UpdateOptions{}
options.SetUpsert(true) // This will insert the document if it doesn't exist
// options.SetReturnDocument(options.After) // This will return the updated document

result, err := entryCollection.UpdateOne(ctx, bson.M{"_id": docID}, update, &options)

    if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Print("error occured while updating entry", err)
		return
	}

	defer cancel()
    c.JSON(http.StatusOK, result.ModifiedCount)
}

func DeleteEntry(c *gin.Context) {
	entryId := c.Params.ByName("id")
	docID, _ := primitive.ObjectIDFromHex(entryId)

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	result, err := entryCollection.DeleteOne(ctx, bson.M{"_id": docID})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Print("error occured while deleting entry", err)
		return
	}

	defer cancel()
	c.JSON(http.StatusOK, result.DeletedCount)
}
