package api

import (
	"GYM-App/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type W struct {
	We string
}

func AddBodyWeight(context *gin.Context) {
	var weight W
	if err := context.ShouldBindJSON(&weight); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := models.InsertWeight(weight.We)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Saved", "weight": weight})
}

func GetAllWeight(context *gin.Context) {
	allWeight, err := models.GetAllWeightDB()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if allWeight == nil {
		context.JSON(http.StatusOK, gin.H{"error": "empty DB"})
		return
	}

	context.JSON(http.StatusOK, allWeight)

}

func RemoveWeight(context *gin.Context) {
	var id models.WeightEntry
	if err := context.ShouldBindJSON(&id.ID); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := models.RemoveYTLinkByID(id.ID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Updated", "Weight": id.ID})

}

// func AddBodyWeight(context *gin.Context) {
// 	var bodyWeight models.BodyWeight
// 	var parseWeight Weight

// 	if err := context.ShouldBindJSON(&parseWeight); err != nil {
// 		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	bodyWeight.Weight, _ = strconv.ParseInt(parseWeight.Weight, 10, 64)

// 	bodyWeight.UserId = context.GetInt64("userId")

// 	bodyWeight.Date = time.Now()
// 	err := bodyWeight.InsertWeight()

// 	if err != nil {
// 		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save"})
// 		return
// 	}

// 	context.JSON(http.StatusOK, gin.H{"message": "Saved", "exercise": bodyWeight})

// }

// func GetBodyWeight(context *gin.Context) {
// 	var allBodyWeight []models.BodyWeight

// 	userId := context.GetInt64("userId") //middleware shit

// 	allBodyWeight, err := models.GetWeight(userId)
// 	if err != nil {
// 		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 	}

// 	context.JSON(http.StatusOK, allBodyWeight)

// }
