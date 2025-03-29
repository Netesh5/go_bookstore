package models

import (
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model

	name        string `gorm:"json:name"`
	author      string `json:"author"`
	publication string `json:"publication"`
}

func init() {
	config.connect()
	db = config.getDB()

}
