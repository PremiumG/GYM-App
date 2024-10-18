package static

import (
	"GYM-App/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Exercises(context *gin.Context) {
	// var allUniqueLastUserWorkouts []models.CompleteUserExercises

	// userId := context.GetInt64("userId") //middleware shit

	// allBodyWeight, err := models.GetWeight(userId)
	// if err != nil {
	// 	context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// }

	// context.JSON(http.StatusOK, allBodyWeight)

	allListedExercises, _ := models.GetAllLatestExercises()

	context.HTML(http.StatusOK, "exercises.tmpl", gin.H{
		"allListedExercises": allListedExercises,
	})
}
