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
		result := database.DB.Create(&group)
		if result.Error != nil {
			return c.Status(500).SendString("Internal server error")
		}
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
		UserID    uint
		GroupName string
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
		database.DB.Model(&models.Group{}).Where("name = ?", data.GroupName).First(&group)
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
			UserID    uint
			GroupName string
		}
		if err := c.BodyParser(&data); err != nil {
			return err
		}
		var user models.User
		// database.DB.First(&user, data.UserID, 1)
		database.DB.Model(&models.User{}).Where("id = ?", data.UserID).First(&user)
		var group models.Group
		// database.DB.First(&group, data.GroupID, 1)
		database.DB.Model(&models.Group{}).Where("name = ?", data.GroupName).First(&group)
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
			TaskID    uint
			GroupName string
		}
		if err := c.BodyParser(&data); err != nil {
			return err
		}
		var task models.Task
		// database.DB.First(&user, data.UserID, 1)
		database.DB.Model(&models.Task{}).Where("id = ?", data.TaskID).First(&task)
		var group models.Group
		// database.DB.First(&group, data.GroupID, 1)
		database.DB.Model(&models.Group{}).Where("name = ?", data.GroupName).First(&group)
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
	var groups []models.Group
	// database.DB.Find(&task)
	token := c.Cookies("JWT")
	user := &models.Claims{}
	_, err := jwt.ParseWithClaims(token, user, func(t *jwt.Token) (interface{}, error) {
		return []byte(secrets.SECRET_JWT), nil
	})
	if err != nil || user.User.ID == 0 {
		return c.Status(401).SendString("Unauthorized")
	}
	if user.User.Role == models.Admin {
		database.DB.Preload("Tasks").Find(&groups)
		return c.JSON(groups)
	} else {
		// find user's groups but know that groups have users and tasks
		database.DB.Preload("Users").Preload("Tasks").Find(&groups)
		var group models.Group
		for _, v := range groups {
			for _, u := range v.Users {
				if u.ID == user.User.ID {
					log.Println("found user", u.ID, "in group", v.ID)
					group = v
					log.Println("group", group.ID, "has", len(group.Tasks), "tasks")
					break

				}
			}
		}

		return c.JSON(group)
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
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(secrets.LIFE_TIME)),
		},
		User: user,
	})
	token, err := t.SignedString([]byte(secrets.SECRET_JWT))
	if err != nil {
		return err
	}
	database.DB.Model(&models.User{}).Where("email = ? AND password = ?", data.Email, data.Password).Update("last_login_time", time.Now())
	cookie := fiber.Cookie{
		Name:    "JWT",
		Value:   token,
		Expires: time.Now().Add(secrets.LIFE_TIME),
		Path:    "/",
	}
	c.Cookie(&cookie)
	return c.JSON(fiber.Map{"result": "success", "HaveToCreateNewUser": database.HaveToCreateFirstUser})
}

// func HaveToCreateFirstUser(c *fiber.Ctx) error {
// 	token := c.Cookies("JWT")
// 	user := &models.Claims{}
// 	_, err := jwt.ParseWithClaims(token, user, func(t *jwt.Token) (interface{}, error) {
// 		return []byte(secrets.SECRET_JWT), nil
// 	})
// 	if err != nil {
// 		return c.Status(401).SendString("Unauthorized")
// 	}
// 	if user.User.Role != models.Admin {
// 		return c.Status(403).SendString("Forbidden")
// 	}
// 	if user.User.ID == 0 {
// 		return c.Status(401).SendString("Unauthorized")
// 	}
// 	if user.User.Role == models.Admin {
// 		if database.HaveToCreateFirstUser {
// 			return c.Status(200).JSON(fiber.Map{"result": true})
// 		} else {
// 			return c.Status(200).JSON(fiber.Map{"result": false})
// 		}
// 	} else {
// 		return c.Status(403).SendString("Forbidden")
// 	}

// }

func ChangeAdmin(c *fiber.Ctx) error {
	token := c.Cookies("JWT")
	log.Println("token", token)
	user := &models.Claims{}
	_, err := jwt.ParseWithClaims(token, user, func(t *jwt.Token) (interface{}, error) {
		return []byte(secrets.SECRET_JWT), nil
	})
	log.Println(user)
	if err != nil {
		return c.Status(401).SendString("Unauthorized")
	}
	if user.User.Role != models.Admin {
		return c.Status(403).SendString("Forbidden")
	}

	if user.User.ID == 0 {
		return c.Status(401).SendString("Unauthorized")
	}
	log.Println("user", user.User.ID, "is admin")
	log.Println("user", user.User.Email, "is admin")
	if user.User.Role == models.Admin && user.User.Email == "admin@admin.pl" {

		database.DB.Where("email LIKE ?", "admin@admin.pl").Delete(&user.User)
		database.HaveToCreateFirstUser = false
		return c.JSON(fiber.Map{"result": "success"})
	} else {
		return c.Status(403).SendString("Forbidden")
	}
}
func GetInformationAboutMe(c *fiber.Ctx) error {
	token := c.Cookies("JWT")
	user := &models.Claims{}
	_, err := jwt.ParseWithClaims(token, user, func(t *jwt.Token) (interface{}, error) {
		return []byte(secrets.SECRET_JWT), nil
	})
	if err != nil {
		return c.Status(401).SendString("Unauthorized")
	}
	if user.User.ID == 0 {
		return c.Status(401).SendString("Unauthorized")
	}
	return c.JSON(user.User)
}
