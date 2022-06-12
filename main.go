package main

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const (
	Admin      = iota // 0
	NormalUser = iota // 1
)

type User struct {
	gorm.Model
	Name     string
	Surname  string
	Email    string
	Password string
	Role     int
	Groups   []Group `gorm:"many2many:GroupRefer"`
}
type Group struct {
	gorm.Model
	Name string
	// Tasks      []Task `gorm:"many2many:task_id"`
	// GroupRefer uint
	// Users []User `gorm:"many2many:ID"`
}

type Task struct {
	ID        int64 `gorm:"column:task_id"`
	CreatedAt time.Time
	Name      string
	Details   string
	StartDate int64
	EndDate   int64
	// Group     Group `gorm:"many2many:ID;"`
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Group{})
	db.AutoMigrate(&Task{})
	// // create example user
	// user := User{
	// 	Name:     "John",
	// 	Surname:  "Doe",
	// 	Email:    "tomasz@palinski.pl",
	// 	Password: "password",
	// 	Role:     Admin,
	// }
	// group := Group{
	// 	Name: "Group 1",
	// }
	// db.Create(&group)
	app := fiber.New(fiber.Config{
		BodyLimit: 1024 * 1024 * 512, // 512MB
	})
	app.Use(logger.New())
	app.Use(requestid.New())
	api := app.Group("/api")
	v1 := api.Group("/v1")
	v1.Get("/getusers", func(c *fiber.Ctx) error {
		var users []User
		db.Find(&users)
		return c.JSON(users)
	})
	v1.Post("/adduser", func(c *fiber.Ctx) error {
		var user User
		if err := c.BodyParser(&user); err != nil {
			return err
		}
		db.Create(&user)
		return c.JSON(user)
	})
	v1.Post("/addgroup", func(c *fiber.Ctx) error {
		var group Group
		if err := c.BodyParser(&group); err != nil {
			return err
		}
		db.Create(&group)
		return c.JSON(group)
	})
	v1.Get("/getgroups", func(c *fiber.Ctx) error {
		var grupa []Group
		db.Find(&grupa)
		return c.JSON(grupa)
	})
	v1.Post("/addusertogroup", func(c *fiber.Ctx) error {
		var data struct {
			UserId  int64 `json:"userId"`
			GroupId int64 `json:"groupId"`
		}
		if err := c.BodyParser(&data); err != nil {
			return err
		}
		// add user to group
		var user User
		db.First(&user, data.UserId)
		var group Group
		db.First(&group, data.GroupId)
		user.Groups = append(user.Groups, group)
		db.Save(&user)
		return c.JSON(user)
	})
	log.Fatal(app.Listen(":3000"))
}
