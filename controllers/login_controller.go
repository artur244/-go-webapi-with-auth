package controllers

import (
	"github.com/artur244/first-go-rest-api/database"
	"github.com/artur244/first-go-rest-api/models"
	"github.com/artur244/first-go-rest-api/services"
	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	db := database.GetDatabase()

	var login models.Login

	err := ctx.ShouldBindJSON(&login)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": "Cannot bind JSON: " + err.Error(),
		})

		return
	}

	var user models.User

	err = db.Where("email = ?", login.Email).First(&user).Error
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": "Não foi possivel encontrar o usuário: " + err.Error(),
		})

		return
	}

	if user.Password != services.SHA256Encoder(login.Password) {
		ctx.JSON(401, gin.H{
			"error": "Credenciais inválidas",
		})

		return
	}

	token, err := services.NewJwtService().GenerateToken(user.ID)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})

		return
	}

	ctx.JSON(200, gin.H{
		"token": token,
	})
}
