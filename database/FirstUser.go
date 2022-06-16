package database

import (
	"filemaod/models"
	"log"
)

func FirstUser() {
	var test models.NoneType

	if DB.Model(&models.NoneType{}).First(&test).Error != nil {
		DB.Create(&models.NoneType{NotNew: true})
		DB.Create(&models.User{
			Name:     "admin",
			Email:    "admin@admin.pl",
			Password: "admin",
			Role:     0,
		})
	} else {
		log.Println("Already created")
	}
}
