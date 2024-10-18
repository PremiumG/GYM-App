package routes

import (
	"GYM-App/middlewares"
	"GYM-App/routes/api"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {

	// server.GET("/events/:id", api.GetEvent)

	// authenticated := server.Group("/") //middleware lahko samo das umes navadnega propertija
	// authenticated.Use(middlewares.Authenticate)

	// authenticated.POST("/events", middlewares.Authenticate, api.CreateEvents)

	// authenticated.PUT("/events/:id", api.UpdateEvent)

	// authenticated.DELETE("/events/:id", api.DeleteEvent)

	// authenticated.POST("/events/:id/register", api.RegisterForEvent)
	// authenticated.DELETE("/events/:id/register", api.CancelRegistration)

	// server.POST("/signup", api.Signup)
	server.POST("/login", api.LoginUser)

	// server.POST("/updateUserExercise/:id", api.UpdateUserExerciseById)

	server.POST("/addBodyWeight", middlewares.Authenticate, api.AddBodyWeight)
	server.GET("/getAllBodyWeight", middlewares.Authenticate, api.GetAllWeight)
	server.POST("/removeBodyWeight", middlewares.Authenticate, api.RemoveWeight)

	server.POST("/setChallenge", middlewares.Authenticate, api.SetChallenge)
	server.POST("/removeChallenge", middlewares.Authenticate, api.RemoveChallenge)

	server.POST("/setYTLink", middlewares.Authenticate, api.SetYTLink)
	server.POST("/removeYTLink", middlewares.Authenticate, api.RemoveYTLink)

	server.POST("/removeExercise", middlewares.Authenticate, api.RemoveExercise)
	server.POST("/addExercise", middlewares.Authenticate, api.AddExercise)
	server.POST("/updateExercise", middlewares.Authenticate, api.UpdateExercise)

}
