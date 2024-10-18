package api

import (
	"GYM-App/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetChallenge(context *gin.Context) {
	var name models.ChallangeName
	if err := context.ShouldBindJSON(&name); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := models.SetRandomChallenge(name.Name)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Saved", "Challange": name})
}

func GetAllChallenge(context *gin.Context) { //not in use as I render on server
	allChallanges, err := models.GetAllChallenges()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if allChallanges == nil {
		context.JSON(http.StatusOK, gin.H{"error": "empty DB"})
		return
	}

	context.JSON(http.StatusOK, allChallanges)

}

type ChallangeID struct {
	ID string
}

func RemoveChallenge(context *gin.Context) {
	var id ChallangeID
	if err := context.ShouldBindJSON(&id); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := models.RemoveChallengesByID(id.ID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Updated", "Challange": id.ID})

}
