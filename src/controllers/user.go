package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/peidrao/instago/src/database"
	"github.com/peidrao/instago/src/domain/models"
	"github.com/peidrao/instago/src/utils"
)

func RegisterUser(context *gin.Context) {
	var user models.User

	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	validate := models.NewValidator()

	err := validate.Struct(user)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			context.Abort()
			return
		}
	}

	user.Password, _ = utils.HashPassword(user.Password)
	user.Active = true
	log.Println("\n\n\n", user)
	record := database.Instance.Create(&user)

	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}
