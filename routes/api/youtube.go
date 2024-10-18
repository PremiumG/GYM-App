package api

import (
	"GYM-App/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetYTLink(context *gin.Context) {
	var ytData models.YT
	if err := context.ShouldBindJSON(&ytData); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := models.SetYTLink(ytData.Title, ytData.Link)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Saved", "ytData": ytData})
}

type id struct {
	ID string
}

func RemoveYTLink(context *gin.Context) {
	var id id
	if err := context.ShouldBindJSON(&id); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := models.RemoveYTLinkByID(id.ID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Updated", "Challange": id.ID})

}
