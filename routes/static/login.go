package static

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//	func Index(context *gin.Context) {
//		context.HTML(http.StatusOK, "index.html", gin.H{
//			"Title":   "gabriel",
//			"Message": "pozdravljen",
//		})
//	}
func Login(context *gin.Context) {
	message := context.Query("message")
	if message == "" {
		context.HTML(http.StatusOK, "login.tmpl", gin.H{
			"Message": nil,
		})
	}
	context.HTML(http.StatusOK, "login.tmpl", gin.H{
		"Message": message,
	})
}
