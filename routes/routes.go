package routes

import (
	"filemaod/controllers"

	"github.com/gofiber/fiber/v2"
)

func Configure(app *fiber.App) {
	app.Static("/", "public")
	app.Static("/changeAdmin", "public")
	api := app.Group("/api")
	v1 := api.Group("/v1")
	v1.Get("/getusers", controllers.GetUsers)
	v1.Post("/adduser", controllers.AddUser)
	v1.Post("/removeuser", controllers.RemoveUser)
	v1.Post("/addgroup", controllers.AddGroup)
	v1.Get("/getgroups", controllers.GetGroups)
	v1.Post("/addgrouptouser", controllers.AssignUserToGroup)
	v1.Post("/removeuserfromgroup", controllers.RemoveUserFromGroup)
	v1.Post("/addtask", controllers.CreateNewTask)
	v1.Post("/assigntask", controllers.AssignTaskToGroup)
	v1.Get("/gettasks", controllers.GetAllGroups)
	v1.Post("/deletetask", controllers.DeleteTask)
	v1.Post("/login", controllers.LoginUser)
	v1.Get("/havetocreatefirstuser", controllers.HaveToCreateFirstUser)
}
