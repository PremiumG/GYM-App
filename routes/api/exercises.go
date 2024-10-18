package api

import (
	"GYM-App/models"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Name struct {
	Id string
}

func RemoveExercise(context *gin.Context) {
	var bruh Name
	if err := context.ShouldBindJSON(&bruh); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := models.RemoveExerciseByName(bruh.Id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete"})
		return
	}
	err = os.Remove("./assets/imagesExercise/" + bruh.Id + ".png")

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete pic"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "removed", "Exercise": bruh.Id})

}

func AddExercise(context *gin.Context) {
	var recivedExercise models.Exercise

	recivedExercise.Name = context.PostForm("name")
	recivedExercise.Weight, _ = strconv.ParseInt(context.PostForm("current_weight"), 10, 64)
	recivedExercise.B_weight, _ = strconv.ParseInt(context.PostForm("b_weight"), 10, 64)
	recivedExercise.I_weight, _ = strconv.ParseInt(context.PostForm("i_weight"), 10, 64)
	recivedExercise.A_weight, _ = strconv.ParseInt(context.PostForm("a_weight"), 10, 64)

	file, err := context.FormFile("img")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "No file is received"})
		return
	}
	imgName := "./assets/imagesExercise/" + recivedExercise.Name + ".png"
	if err := context.SaveUploadedFile(file, imgName); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save the file"})
		return
	}
	models.AddExercise(recivedExercise, imgName)
	context.JSON(http.StatusOK, gin.H{"message": "added", "Exercise": "a"})

}

type Kg struct {
	Name   string
	Weight string
}

func UpdateExercise(context *gin.Context) {
	var nu Kg

	if err := context.ShouldBindJSON(&nu); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	number, _ := strconv.ParseInt(nu.Weight, 10, 64)
	fmt.Println(nu.Name, number)
	err := models.UpdateExerciseByName(nu.Name, number)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "updated", "Exercise": nu.Name})

}
