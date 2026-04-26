package user

import "github.com/gin-gonic/gin"

func RegisterUserRoutes(rg *gin.RouterGroup, h UserHandler) {
	userRoutes := rg.Group("/users")

	userRoutes.POST("/", h.CreateUser)
	userRoutes.GET("/:userId", h.GetUserById)
}