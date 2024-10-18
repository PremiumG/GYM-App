package static

import (
	"GYM-App/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetDashboard(context *gin.Context) {
	challengeData := models.GetRandomChallenge()
	ytData := models.GetRandomYTLink()
	allExerciseWeight, _ := models.GetAllWeightsEX()
	// userId := context.GetInt64("userId")
	context.HTML(http.StatusOK, "dashboard.tmpl", gin.H{
		"challenge": challengeData,
		"ytLink":    ytData,
		"exWeight":  allExerciseWeight,
	})

}
