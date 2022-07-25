package controllers

import (
	"strconv"

	"github.com/artur244/first-go-rest-api/database"
	"github.com/artur244/first-go-rest-api/models"
	"github.com/gin-gonic/gin"
)

func ShowBook(ctx *gin.Context) {
	id := ctx.Param("id")

	newid, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": "ID precisa ser inteiro",
		})

		return
	}

	db := database.GetDatabase()

	var book models.Book
	err = db.First(&book, newid).Error
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": "Não foi possivel encontrar o livro: " + err.Error(),
		})

		return
	}

	ctx.JSON(200, book)
}

func CreateBook(ctx *gin.Context) {
	db := database.GetDatabase()

	var book models.Book

	err := ctx.ShouldBindJSON(&book)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": "Cannot bind JSON: " + err.Error(),
		})

		return
	}

	err = db.Create(&book).Error
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": "Não foi possivel criar o livro: " + err.Error(),
		})

		return
	}

	ctx.JSON(200, book)
}

func ShowBooks(ctx *gin.Context) {
	db := database.GetDatabase()

	var books []models.Book

	err := db.Find(&books).Error
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": "Não foi possivel buscar grupos: " + err.Error(),
		})

		return
	}

	ctx.JSON(200, books)
}

func UpdateBook(ctx *gin.Context) {
	db := database.GetDatabase()

	var book models.Book

	err := ctx.ShouldBindJSON(&book)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": "Cannot bind JSON: " + err.Error(),
		})

		return
	}

	err = db.Save(&book).Error
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": "Não foi possivel atualizar o livro: " + err.Error(),
		})

		return
	}

	ctx.JSON(200, book)
}

func DeleteBook(ctx *gin.Context) {
	id := ctx.Param("id")

	newid, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": "ID precisa ser inteiro",
		})

		return
	}

	db := database.GetDatabase()

	err = db.Delete(&models.Book{}, newid).Error
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": "Não foi possivel deletar o livro",
		})

		return
	}

	ctx.Status(204)
}
