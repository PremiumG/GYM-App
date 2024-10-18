package routes

import (
	"GYM-App/middlewares"
	"GYM-App/routes/static"

	"github.com/gin-gonic/gin"
)

func RegisterStaticRoutes(server *gin.Engine) {
	server.GET("/firspage", static.Login) // "/" will check if cookie already there if not it will redirect to "firstpage" if correct cookie is there it will continue with "/"
	server.GET("/", middlewares.AlreadyAuthenticated, static.GetDashboard)

	server.GET("/dashboard", middlewares.Authenticate, static.GetDashboard)
	server.GET("/extra", middlewares.Authenticate, static.GetExtra)

	server.GET("/exercises", middlewares.Authenticate, static.Exercises)
}
