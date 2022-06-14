package main

import (
	"database/sql"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const (
	Admin      = iota // 0
	NormalUser = iota // 1
)
const LIFE_TIME = time.Hour * 1 // 1 hour
const SECRET_JWT = "ljfmyhrmq1"

var SECRET_PASSWORD = []byte("jkrtok9k")

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
	IDGroup   uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
	Name      string
	Users     []User `gorm:"many2many:user_groups;"`
	Tasks     []Task `gorm:"many2many:group_tasks;"`
}

type Task struct {
	gorm.Model
	Name        string
	Description string
	Done        bool
	// Group       *[]Group `gorm:"many2many:group_tasks;"`
}
type NoneType struct {
	NotNew bool
}

func main() {
	// dsn := "postgres://postgres:postgrespw@localhost:49155"
	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db, err := gorm.Open(sqlite.Open("./db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&User{}, &Group{}, &Task{}, &NoneType{})
	var test NoneType

	if db.Model(&NoneType{}).First(&test).Error != nil {
		db.Create(&NoneType{NotNew: true})
		db.Create(&User{
			Name:     "admin",
			Email:    "admin@admin.pl",
			Password: "admin",
			Role:     Admin,
		})
	} else {
		log.Println("Already created")
	}

	app := fiber.New(fiber.Config{
		BodyLimit: 1024 * 1024 * 512, // 512MB
	})
	app.Use(logger.New())
	app.Use(requestid.New())
	api := app.Group("/api")
	v1 := api.Group("/v1")
	v1.Get("/getusers", func(c *fiber.Ctx) error {
		token := c.Get("JWT")
		user := &Claims{}
		_, err := jwt.ParseWithClaims(token, user, func(t *jwt.Token) (interface{}, error) {
			return []byte(SECRET_JWT), nil
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
			var users []User
			db.Find(&users)
			return c.JSON(users)
		} else {
			return c.Status(403).SendString("Forbidden")
		}

	})
	v1.Post("/adduser", func(c *fiber.Ctx) error {
		token := c.Get("JWT")
		user := &Claims{}
		_, err := jwt.ParseWithClaims(token, user, func(t *jwt.Token) (interface{}, error) {
			return []byte(SECRET_JWT), nil
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
			var user User
			if err := c.BodyParser(&user); err != nil {
				return err
			}
			db.Create(&user)
			return c.JSON(fiber.Map{"message": "User created successfully"})
		} else {
			return c.Status(403).SendString("Forbidden")
		}
	})
	v1.Post("/addgroup", func(c *fiber.Ctx) error {
		token := c.Get("JWT")
		user := &Claims{}
		_, err := jwt.ParseWithClaims(token, user, func(t *jwt.Token) (interface{}, error) {
			return []byte(SECRET_JWT), nil
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
			var group Group
			if err := c.BodyParser(&group); err != nil {
				return err
			}
			db.Create(&group)
			return c.JSON(group)
		} else {
			return c.Status(403).SendString("Forbidden")
		}
	})
	v1.Get("/getgroups", func(c *fiber.Ctx) error {
		token := c.Get("JWT")
		user := &Claims{}
		_, err := jwt.ParseWithClaims(token, user, func(t *jwt.Token) (interface{}, error) {
			return []byte(SECRET_JWT), nil
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
			var grupa []Group
			db.Preload("Users").Preload("Tasks").Find(&grupa)
			return c.JSON(grupa)
		} else {
			return c.Status(403).SendString("Forbidden")
		}

	})
	v1.Post("/addgrouptouser", func(c *fiber.Ctx) error {
		token := c.Get("JWT")
		user := &Claims{}
		_, err := jwt.ParseWithClaims(token, user, func(t *jwt.Token) (interface{}, error) {
			return []byte(SECRET_JWT), nil
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

		var data struct {
			UserID  uint
			GroupID uint
		}
		if err := c.BodyParser(&data); err != nil {
			return err
		}
		if user.User.Role == Admin {

			var user User
			// db.First(&user, data.UserID, 1)
			db.Model(&User{}).Where("id = ?", data.UserID).First(&user)
			var group Group
			// db.First(&group, data.GroupID, 1)
			db.Model(&Group{}).Where("id = ?", data.UserID).First(&group)
			if user.ID == 0 || group.IDGroup == 0 {
				return c.Status(400).SendString("User or group not found")
			}
			group.Users = append(group.Users, user)
			db.Save(&group)
			return c.JSON(user)
		} else {
			return c.Status(403).SendString("Forbidden")
		}
	})
	//kopiowanie
	v1.Post("/addtask", func(c *fiber.Ctx) error {
		token := c.Get("JWT")
		user := &Claims{}
		_, err := jwt.ParseWithClaims(token, user, func(t *jwt.Token) (interface{}, error) {
			return []byte(SECRET_JWT), nil
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
		token := c.Get("JWT")
		user := &Claims{}
		_, err := jwt.ParseWithClaims(token, user, func(t *jwt.Token) (interface{}, error) {
			return []byte(SECRET_JWT), nil
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
			if task.ID == 0 || group.IDGroup == 0 {
				return c.Status(400).SendString("Task or group not found")
			}
			group.Tasks = append(group.Tasks, task)
			db.Save(&group)
			return c.JSON(task)
		} else {
			return c.Status(403).SendString("Forbidden")
		}
	})
	v1.Get("/gettasks", func(c *fiber.Ctx) error {
		var group []Group
		// db.Find(&task)
		token := c.Get("JWT")
		user := &Claims{}
		_, err := jwt.ParseWithClaims(token, user, func(t *jwt.Token) (interface{}, error) {
			return []byte(SECRET_JWT), nil
		})
		if err != nil || user.User.ID == 0 {
			return c.Status(401).SendString("Unauthorized")
		}
		if user.User.Role != Admin {
			return c.Status(403).SendString("Forbidden")
		}

		if user.User.Role == NormalUser {
			var grupa []Group
			db.Preload("Users").Preload("Tasks").Where("ID = ?", user.User.ID).Find(&grupa)
			return c.JSON(grupa)
		}
		if user.User.Role == Admin {
			db.Preload("Tasks").Find(&group)
			return c.JSON(group)
		}
		return c.Status(401).SendString("Unauthorized")
		// return c.JSON(group)
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
		token, err := t.SignedString([]byte(SECRET_JWT))
		if err != nil {
			return err
		}
		return c.JSON(fiber.Map{"token": token})
	})
	v1.Get("/ttt", func(c *fiber.Ctx) error {
		token := c.Get("JWT")
		user := &Claims{}
		_, err := jwt.ParseWithClaims(token, user, func(t *jwt.Token) (interface{}, error) {
			return []byte(SECRET_JWT), nil
		})
		if err != nil {
			return c.Status(401).SendString("Unauthorized")
		}
		var grupa []Group
		db.Preload("Users").Preload("Tasks").Find(&grupa)
		return c.JSON(grupa)
	})
	log.Fatal(app.Listen(":3000"))
}
