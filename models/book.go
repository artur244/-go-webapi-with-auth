package models

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Name      string    `json:"name"`
	Desc      string    `json:"desc"`
	Preco     float64   `json:"preco"`
	Autor     string    `json:"autor"`
	ImagemURL string    `json:"imagem_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt gorm.DeletedAt
}
