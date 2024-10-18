package main

import (
	"GYM-App/db"
	"GYM-App/routes"
	"GYM-App/utils"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	//gin.SetMode(gin.ReleaseMode) //remove for debug/prod mode

	//--------------------------------------------------------------------------------
	err := godotenv.Load() //load .env
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initializethe  the sqlite DB
	db.InitDB()

	//Get reset flag to reset the database for testing-------------------------------------------------------------------------------- go run . -reset=true
	//go run . -username=xxxxx -pass=xxxxx
	utils.RunFlags()

	//--------------------------------------------------------------------------------

	server := gin.Default()

	//load templates and static
	server.LoadHTMLGlob("templates/*")
	server.Static("/assets", "./assets")

	routes.RegisterStaticRoutes(server)
	routes.RegisterRoutes(server)

	srv := &http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: server,
	}

	go func() {
		// server.Run("127.0.0.1:8080") //localhost:8080

		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("listen: %s\n", err)
		}
	}()

	//--------------------------------------------------------------------------------
	//Gracefuly quit app with ctrl C
	utils.QuitGracefuly(srv)

}
