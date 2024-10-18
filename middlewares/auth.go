package middlewares

import (
	"GYM-App/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")
	if token == "" {
		// Alternatively, check for the token in a cookie
		cookie, err := context.Cookie("token")
		if err == nil {
			token = cookie
		}
	}

	if token == "" {
		context.Redirect(http.StatusSeeOther, "/?message=Please log in.")

		// context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorizated"})
		context.Abort()
		return
	}

	userId, err := utils.VerifyToken(token)

	if err != nil {
		context.Redirect(http.StatusSeeOther, "/?message=Sesion expired. Please log in.")

		// context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorizated"})
		context.Abort()
		return
	}
	context.Set("userId", userId)
	context.Next()
}

func AlreadyAuthenticated(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")
	if token == "" {
		// Alternatively, check for the token in a cookie
		cookie, err := context.Cookie("token")
		if err == nil {
			token = cookie
		}
	}

	if token == "" {
		context.Redirect(http.StatusSeeOther, "/firspage?message=Please log in.")

		// context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorizated"})
		context.Abort()
		return
	}

	userId, err := utils.VerifyToken(token)

	if err != nil {
		context.Redirect(http.StatusSeeOther, "/firspage?message=Sesion expired. Please log in.")

		// context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorizated"})
		context.Abort()
		return
	}
	context.Set("userId", userId)

	context.Next()
}
