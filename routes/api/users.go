package api

import (
	"GYM-App/models"
	"GYM-App/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginUser(context *gin.Context) {
	var user models.User

	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := user.ValidateCredentials() //returns custum err

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "could not authenticate"}) //redirect or tell that in js for wrong credentials
		return
	}

	token, err := utils.GenerateToken(int(user.ID))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not save token"}) //redirect or tell that in js for wrong credentials
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Login sucessful", "token": token})

}
