package models

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

const (
	Admin      = iota // 0
	NormalUser = iota // 1
)

type Claims struct {
	jwt.RegisteredClaims
	User User
}
type User struct {
	gorm.Model
	LastLoginTime time.Time `gorm:"column:last_login_time"`
	Name          string
	Surname       string
	Email         string
	Password      string
	Role          int
	// Groups   []*Group `gorm:"many2many:user_groups;"`
}
type Group struct {
	gorm.Model
	Name  string `gorm:"unique"`
	Users []User `gorm:"many2many:user_groups;"`
	Tasks []Task `gorm:"many2many:group_tasks;"`
}

type Task struct {
	gorm.Model
	Name        string
	Description string
	Hidden      bool
	// Group       *[]Group `gorm:"many2many:group_tasks;"`
}
type NoneType struct {
	NotNew bool
}
