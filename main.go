package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	Admin      = iota // 0
	NormalUser = iota // 1
)

// Create structs for our data models
type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
	Role     int
	Groups   []*Group `gorm:"many2many:user_groups;"`
}
type Group struct {
	gorm.Model
	Name  string
	Users []*User `gorm:"many2many:user_groups;"`
	Tasks []*Task `gorm:"many2many:group_tasks;"`
}

type Task struct {
	gorm.Model
	Name        string
	Description string
	Done        bool
	Group       *[]Group `gorm:"many2many:group_tasks;"`
}

func main() {
	dsn := "postgres://postgres:postgrespw@localhost:49155"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&User{}, &Group{}, &Task{})
	// db.AutoMigrate(&Task{})

	app := fiber.New(fiber.Config{
		BodyLimit: 1024 * 1024 * 512, // 512MB
	})
	app.Use(logger.New())
	app.Use(requestid.New())
	api := app.Group("/api")
	v1 := api.Group("/v1")
	v1.Get("/getusers", func(c *fiber.Ctx) error {
		// Get all users and his groups
		var users []User
		db.Preload("Groups").Find(&users)
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
		db.Preload("Users").Preload("Tasks").Find(&grupa)
		return c.JSON(grupa)
	})
	v1.Post("/addgrouptouser", func(c *fiber.Ctx) error {
		var data struct {
			UserID  uint
			GroupID uint
		}
		if err := c.BodyParser(&data); err != nil {
			return err
		}
		
		var user User
		// db.First(&user, data.UserID, 1)
		db.Model(&User{}).Where("id = ?", data.UserID).First(&user)
		var group Group
		// db.First(&group, data.GroupID, 1)
		db.Model(&Group{}).Where("id = ?", data.UserID).First(&group)
		if group. == nil {
			return c.Status(fiber.StatusBadRequest).SendString("Group not found")
		}

		user.Groups = append(user.Groups, &group)
		db.Save(&user)
		return c.JSON(user)

	})
	v1.Post("/addtask", func(c *fiber.Ctx) error {
		var task Task
		if err := c.BodyParser(&task); err != nil {
			return err
		}
		db.Create(&task)
		return c.JSON(task)
	})
	v1.Post("/assigntask", func(c *fiber.Ctx) error {
		var data struct {
			TaskID  uint
			GroupID uint
		}
		if err := c.BodyParser(&data); err != nil {
			return err
		}
		var task Task
		// db.First(&user, data.UserID, 1)
		db.Model(&Task{}).Preload("Group").Where("id = ?", data.TaskID).First(&task)
		var group Group
		// db.First(&group, data.GroupID, 1)
		db.Model(&Group{}).Where("id = ?", data.GroupID).First(&group)

		group.Tasks = append(group.Tasks, &task)
		db.Save(&task)
		return c.JSON(task)
	})
	v1.Get("/gettasks", func(c *fiber.Ctx) error {
		var task []Task
		db.Preload("Group").Find(&task)
		return c.JSON(task)
	})
	log.Fatal(app.Listen(":3000"))
}
