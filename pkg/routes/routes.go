package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pavanmuttange/go-folder-structure/controllers"
	"github.com/pavanmuttange/go-folder-structure/services"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	userService := &services.UserService{DB: db}
	userController := &controllers.UserController{Service: userService}

	userGroup := r.Group("/users")
	{
		userGroup.POST("/", userController.CreateUser)
		userGroup.GET("/", userController.GetUsers)
		userGroup.GET("/:id", userController.GetUserByID)
		userGroup.PUT("/:id", userController.UpdateUser)
		userGroup.DELETE("/:id", userController.DeleteUser)
	}

	return r
}
