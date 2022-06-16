package controllers

import (
	"filemaod/database"
	"filemaod/models"
	"filemaod/secrets"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func GetUsers(c *fiber.Ctx) error {
	token := c.Cookies("JWT")
	user := &models.Claims{}
	_, err := jwt.ParseWithClaims(token, user, func(t *jwt.Token) (interface{}, error) {
		return []byte(secrets.SECRET_JWT), nil
	})
	if err != nil {
		return c.Status(401).SendString("Unauthorized")
	}
	if user.User.Role != models.Admin {
		return c.Status(403).SendString("Forbidden")
	}
	if user.User.ID == 0 {
		return c.Status(401).SendString("Unauthorized")
	}
	if user.User.Role == models.Admin {
		var users []models.User
		database.DB.Find(&users)
		return c.JSON(users)
	} else {
		return c.Status(403).SendString("Forbidden")
	}
}
func AddUser(c *fiber.Ctx) error {
	token := c.Cookies("JWT")
	user := &models.Claims{}
	_, err := jwt.ParseWithClaims(token, user, func(t *jwt.Token) (interface{}, error) {
		return []byte(secrets.SECRET_JWT), nil
	})
	if err != nil {
		return c.Status(401).SendString("Unauthorized")
	}
	if user.User.Role != models.Admin {
		return c.Status(403).SendString("Forbidden")
	}
	if user.User.ID == 0 {
		return c.Status(401).SendString("Unauthorized")
	}

	if user.User.Role == models.Admin {
		var user models.User
		if err := c.BodyParser(&user); err != nil {
			return err
		}
		database.DB.Create(&user)
		return c.JSON(fiber.Map{"message": "User created successfully"})
	} else {
		return c.Status(403).SendString("Forbidden")
	}
}
func RemoveUser(c *fiber.Ctx) error {
	token := c.Cookies("JWT")
	user := &models.Claims{}
	_, err := jwt.ParseWithClaims(token, user, func(t *jwt.Token) (interface{}, error) {
		return []byte(secrets.SECRET_JWT), nil
	})
	if err != nil {
		return c.Status(401).SendString("Unauthorized")
	}
	if user.User.Role != models.Admin {
		return c.Status(403).SendString("Forbidden")
	}
	if user.User.ID == 0 {
		return c.Status(401).SendString("Unauthorized")
	}
	var data struct {
		UserID uint `json:"user_id"`
	}
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	if user.User.Role == models.Admin {
		var user models.User
		database.DB.First(&user, data.UserID)
		database.DB.Delete(&user)
		return c.JSON(fiber.Map{"message": "User deleted successfully"})
	} else {
		return c.Status(403).SendString("Forbidden")
	}
}
func AddGroup(c *fiber.Ctx) error {
	token := c.Cookies("JWT")
	user := &models.Claims{}
	_, err := jwt.ParseWithClaims(token, user, func(t *jwt.Token) (interface{}, error) {
		return []byte(secrets.SECRET_JWT), nil
	})
	if err != nil {
		return c.Status(401).SendString("Unauthorized")
	}
	if user.User.Role != models.Admin {
		return c.Status(403).SendString("Forbidden")
	}
	if user.User.ID == 0 {
		return c.Status(401).SendString("Unauthorized")
	}

	if user.User.Role == models.Admin {
		var group models.Group
		if err := c.BodyParser(&group); err != nil {
			return err
		}
		database.DB.Create(&group)
		return c.JSON(group)
	} else {
		return c.Status(403).SendString("Forbidden")
	}
}

func GetGroups(c *fiber.Ctx) error {
	token := c.Cookies("JWT")
	user := &models.Claims{}
	_, err := jwt.ParseWithClaims(token, user, func(t *jwt.Token) (interface{}, error) {
		return []byte(secrets.SECRET_JWT), nil
	})
	if err != nil {
		return c.Status(401).SendString("Unauthorized")
	}
	if user.User.Role != models.Admin {
		return c.Status(403).SendString("Forbidden")
	}
	if user.User.ID == 0 {
		return c.Status(401).SendString("Unauthorized")
	}
	if user.User.Role == models.Admin {
		var grupa []models.Group
		database.DB.Preload("Users").Preload("Tasks").Find(&grupa)
		return c.JSON(grupa)
	} else {
		return c.Status(403).SendString("Forbidden")
	}

}
func AssignUserToGroup(c *fiber.Ctx) error {
	token := c.Cookies("JWT")
	user := &models.Claims{}
	_, err := jwt.ParseWithClaims(token, user, func(t *jwt.Token) (interface{}, error) {
		return []byte(secrets.SECRET_JWT), nil
	})
	if err != nil {
		return c.Status(401).SendString("Unauthorized")
	}
	if user.User.Role != models.Admin {
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
	if user.User.Role == models.Admin {

		var user models.User
		// database.DB.First(&user, data.UserID, 1)
		database.DB.Model(&models.User{}).Where("id = ?", data.UserID).First(&user)
		var group models.Group
		// database.DB.First(&group, data.GroupID, 1)
		database.DB.Model(&models.Group{}).Where("id = ?", data.GroupID).First(&group)
		if user.ID == 0 || group.ID == 0 {
			return c.Status(400).SendString("User or group not found")
		}
		group.Users = append(group.Users, user)
		database.DB.Save(&group)
		return c.JSON(user)
	} else {
		return c.Status(403).SendString("Forbidden")
	}
}

func RemoveUserFromGroup(c *fiber.Ctx) error {
	token := c.Cookies("JWT")
	user := &models.Claims{}
	_, err := jwt.ParseWithClaims(token, user, func(t *jwt.Token) (interface{}, error) {
		return []byte(secrets.SECRET_JWT), nil
	})
	if err != nil {
		return c.Status(401).SendString("Unauthorized")
	}
	if user.User.Role != models.Admin {
		return c.Status(403).SendString("Forbidden")
	}
	if user.User.ID == 0 {
		return c.Status(401).SendString("Unauthorized")
	}
	if user.User.Role == models.Admin {
		var data struct {
			UserID  uint
			GroupID uint
		}
		if err := c.BodyParser(&data); err != nil {
			return err
		}
		var user models.User
		// database.DB.First(&user, data.UserID, 1)
		database.DB.Model(&models.User{}).Where("id = ?", data.UserID).First(&user)
		var group models.Group
		// database.DB.First(&group, data.GroupID, 1)
		database.DB.Model(&models.Group{}).Where("id = ?", data.UserID).First(&group)
		if user.ID == 0 || group.ID == 0 {
			return c.Status(400).SendString("User or group not found")
		}
		for i, v := range group.Users {
			if v.ID == user.ID {
				group.Users = append(group.Users[:i], group.Users[i+1:]...)
				break
			}
		}
		database.DB.Save(&group)
		return c.JSON(user)
	} else {
		return c.Status(403).SendString("Forbidden")
	}
}
func CreateNewTask(c *fiber.Ctx) error {
	token := c.Cookies("JWT")
	user := &models.Claims{}
	_, err := jwt.ParseWithClaims(token, user, func(t *jwt.Token) (interface{}, error) {
		return []byte(secrets.SECRET_JWT), nil
	})
	if err != nil {
		return c.Status(401).SendString("Unauthorized")
	}
	if user.User.Role != models.Admin {
		return c.Status(403).SendString("Forbidden")
	}
	if user.User.ID == 0 {
		return c.Status(401).SendString("Unauthorized")
	}
	if user.User.Role == models.Admin {
		var task models.Task
		if err := c.BodyParser(&task); err != nil {
			return err
		}
		database.DB.Create(&task)
		database.DB.First(&task)
		return c.JSON(task)
	}
	return c.Status(401).SendString("Unauthorized")
}
func AssignTaskToGroup(c *fiber.Ctx) error {
	token := c.Cookies("JWT")
	user := &models.Claims{}
	_, err := jwt.ParseWithClaims(token, user, func(t *jwt.Token) (interface{}, error) {
		return []byte(secrets.SECRET_JWT), nil
	})
	if err != nil {
		return c.Status(401).SendString("Unauthorized")
	}
	if user.User.Role != models.Admin {
		return c.Status(403).SendString("Forbidden")
	}
	if user.User.ID == 0 {
		return c.Status(401).SendString("Unauthorized")
	}
	if user.User.Role == models.Admin {
		var data struct {
			TaskID  uint
			GroupID uint
		}
		if err := c.BodyParser(&data); err != nil {
			return err
		}
		var task models.Task
		// database.DB.First(&user, data.UserID, 1)
		database.DB.Model(&models.Task{}).Where("id = ?", data.TaskID).First(&task)
		var group models.Group
		// database.DB.First(&group, data.GroupID, 1)
		database.DB.Model(&models.Group{}).Where("id = ?", data.GroupID).First(&group)
		if task.ID == 0 || group.ID == 0 {
			return c.Status(400).SendString("Task or group not found")
		}
		group.Tasks = append(group.Tasks, task)
		database.DB.Save(&group)
		return c.JSON(task)
	} else {
		return c.Status(403).SendString("Forbidden")
	}
}

func GetAllGroups(c *fiber.Ctx) error {
	var group []models.Group
	// database.DB.Find(&task)
	token := c.Cookies("JWT")
	user := &models.Claims{}
	_, err := jwt.ParseWithClaims(token, user, func(t *jwt.Token) (interface{}, error) {
		return []byte(secrets.SECRET_JWT), nil
	})
	if err != nil || user.User.ID == 0 {
		return c.Status(401).SendString("Unauthorized")
	}
	log.Println(user.User.ID, user.User.Role)
	if user.User.Role == models.Admin {
		database.DB.Preload("Tasks").Find(&group)
		return c.JSON(group)
	} else {
		var grupa []models.Group
		database.DB.Preload("Users").Preload("Tasks").Where("id = ?", user.User.ID).Find(&grupa)
		return c.JSON(grupa)
	}
	// return c.Status(401).SendString("Unauthorized")
}

func DeleteTask(c *fiber.Ctx) error {
	token := c.Cookies("JWT")
	user := &models.Claims{}
	_, err := jwt.ParseWithClaims(token, user, func(t *jwt.Token) (interface{}, error) {
		return []byte(secrets.SECRET_JWT), nil
	})
	if err != nil {
		return c.Status(401).SendString("Unauthorized")
	}
	if user.User.Role != models.Admin {
		return c.Status(403).SendString("Forbidden")
	}
	if user.User.ID == 0 {
		return c.Status(401).SendString("Unauthorized")
	}
	if user.User.Role == models.Admin {
		var data struct {
			TaskID uint
		}
		if err := c.BodyParser(&data); err != nil {
			return err
		}
		var task models.Task
		// database.DB.First(&user, data.UserID, 1)
		database.DB.Model(&models.Task{}).Where("id = ?", data.TaskID).First(&task)
		if task.ID == 0 {
			return c.Status(400).SendString("Task not found")
		}
		database.DB.Delete(&task)
		return c.JSON(task)
	} else {
		return c.Status(403).SendString("Forbidden")
	}
}

func LoginUser(c *fiber.Ctx) error {
	c.Accepts("application/json")
	c.Accepts("json", "text")
	c.AcceptsEncodings("compress", "br")
	var data struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.BodyParser(&data); err != nil {
		log.Println(string(c.Body()))
		return err
	}
	var user models.User
	database.DB.Where("email = ? AND password = ?", data.Email, data.Password).First(&user) // 2a@a.pl    asasas
	if user.ID == 0 {
		return c.Status(400).SendString("User not found")
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, models.Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(secrets.LIFE_TIME).Unix(),
		},
		User: user,
	})
	token, err := t.SignedString([]byte(secrets.SECRET_JWT))
	if err != nil {
		return err
	}
	cookie := fiber.Cookie{
		Name:    "JWT",
		Value:   token,
		Expires: time.Now().Add(secrets.LIFE_TIME),
		Path:    "/",
	}
	c.Cookie(&cookie)
	return c.JSON(fiber.Map{"result": "success"})
}
