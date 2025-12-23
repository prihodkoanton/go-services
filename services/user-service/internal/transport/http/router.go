package http

import (
	"github.com/gin-gonic/gin"
	"github.com/prihodkoanton/go-services/services/user-service/internal/service"

	_ "github.com/prihodkoanton/go-services/services/user-service/internal/transport/http/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter(userService *service.UserService) *gin.Engine {
	router := gin.Default()
	userHandler := NewUserHandler(userService)

	api := router.Group("/api/v1")
	{
		api.POST("/users", userHandler.CreateUser)
		api.GET("/users/:id", userHandler.GetUserById)
		api.GET("/users/emails/:email", userHandler.GetUserByEmail)
		api.PUT("/users/:id", userHandler.UpdateUser)
		api.DELETE("/users/:id", userHandler.DeleteUser)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL("/swagger/doc.json")))

	return router
}
