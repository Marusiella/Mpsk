package database

import (
	"filemaod/models"
	"log"

	"golang.org/x/crypto/bcrypt"
)

var HaveToCreateFirstUser = true

func FirstUser() {
	// chck if there is a first user in the database and if not create one and set HaveToCreateFirstUser to true but if there is more than one user in the database set HaveToCreateFirstUser to false
	var user models.User
	DB.First(&user)
	hash, err := bcrypt.GenerateFromPassword([]byte("admin"), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	if user.ID == 0 {
		log.Println("There is no user in the database, creating first user")
		HaveToCreateFirstUser = true
		user.Name = "admin"
		user.Password = string(hash)
		user.Email = "admin@admin.pl"
		user.Role = models.Admin
		DB.Create(&user)
	} else {
		HaveToCreateFirstUser = false
	}

}
