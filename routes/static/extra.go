package static

import (
	"GYM-App/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetExtra(context *gin.Context) {

	challenges, err := models.GetAllChallenges()
	if err != nil {
		context.HTML(http.StatusOK, "dashboard.tmpl", gin.H{}) //change status
	}

	ytLinks, err := models.GetAllYTLink()
	if err != nil {
		context.HTML(http.StatusOK, "dashboard.tmpl", gin.H{}) //change status
	}

	context.HTML(http.StatusOK, "extra.tmpl", gin.H{
		"items":   challenges,
		"ytItems": ytLinks,
	})

}
