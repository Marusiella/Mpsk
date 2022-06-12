package main

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	Admin      = iota // 0
	NormalUser = iota // 1
)
const LIFE_TIME = time.Hour * 1 // 1 hour
const SECRET = "ljfmyhrmq1"

type Claims struct {
	jwt.StandardClaims
	User User
}
type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
	Role     int
	// Groups   []*Group `gorm:"many2many:user_groups;"`
}
type Group struct {
	gorm.Model
	Name  string
	Users []User `gorm:"many2many:user_groups;"`
	Tasks []Task `gorm:"many2many:group_tasks;"`
}

type Task struct {
	gorm.Model
	Name        string
	Description string
	Done        bool
	// Group       *[]Group `gorm:"many2many:group_tasks;"`
}

func main() {
	dsn := "postgres://postgres:postgrespw@localhost:49153"
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
		db.Find(&users)
		return c.JSON(users)

	})
	v1.Post("/adduser", func(c *fiber.Ctx) error {
		var user User
		if err := c.BodyParser(&user); err != nil {
			return err
		}
		db.Create(&user)
		return c.JSON(fiber.Map{"message": "User created successfully"})
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
		if user.ID == 0 || group.ID == 0 {
			return c.Status(400).SendString("User or group not found")
		}
		group.Users = append(group.Users, user)
		db.Save(&group)
		return c.JSON(user)

	})
	v1.Post("/addtask", func(c *fiber.Ctx) error {
		token := c.Get("JWT")
		user := &Claims{}
		_, err := jwt.ParseWithClaims(token, user, func(t *jwt.Token) (interface{}, error) {
			return []byte(SECRET), nil
		})
		if err != nil {
			return c.Status(401).SendString("Unauthorized")
		}
		if user.User.Role != Admin {
			return c.Status(403).SendString("Forbidden")
		}
		if user.User.ID == 0 {
			return c.Status(401).SendString("Unauthorized")
		}
		if user.User.Role == Admin {
			var task Task
			if err := c.BodyParser(&task); err != nil {
				return err
			}
			db.Create(&task)
			return c.JSON(task)
		}
		return c.Status(401).SendString("Unauthorized")
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
		db.Model(&Task{}).Where("id = ?", data.TaskID).First(&task)
		var group Group
		// db.First(&group, data.GroupID, 1)
		db.Model(&Group{}).Where("id = ?", data.GroupID).First(&group)
		if task.ID == 0 || group.ID == 0 {
			return c.Status(400).SendString("Task or group not found")
		}
		group.Tasks = append(group.Tasks, task)
		db.Save(&group)
		return c.JSON(task)
	})
	v1.Get("/gettasks", func(c *fiber.Ctx) error {
		var task []Task
		db.Find(&task)
		return c.JSON(task)
	})
	v1.Post("/login", func(c *fiber.Ctx) error {
		var data struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}
		if err := c.BodyParser(&data); err != nil {
			return err
		}
		var user User
		db.Where("email = ? AND password = ?", data.Email, data.Password).First(&user) // 2a@a.pl    asasas
		if user.ID == 0 {
			return c.Status(400).SendString("User not found")
		}

		t := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add(LIFE_TIME).Unix(),
			},
			User: user,
		})
		token, err := t.SignedString([]byte(SECRET))
		if err != nil {
			return err
		}
		return c.JSON(fiber.Map{"token": token})
	})
	log.Fatal(app.Listen(":3000"))
}
