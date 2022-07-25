package controllers

import (
	"github.com/artur244/first-go-rest-api/database"
	"github.com/artur244/first-go-rest-api/models"
	"github.com/artur244/first-go-rest-api/services"
	"github.com/gin-gonic/gin"
)

func CreateUser(ctx *gin.Context) {
	db := database.GetDatabase()

	var user models.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": "Cannot bind JSON: " + err.Error(),
		})
	}

	user.Password = services.SHA256Encoder(user.Password)

	err = db.Create(&user).Error
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": "Não foi possivel criar o usuário: " + err.Error(),
		})
	}

	ctx.Status(204)
}
