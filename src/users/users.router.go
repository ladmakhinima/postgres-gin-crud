package users

import "github.com/gin-gonic/gin"

func LoadUserRoutes(server *gin.Engine) {
	usersRoute := server.Group("/api/users")
	usersRoute.POST("/", CreateUserAction)
	usersRoute.GET("/", LoadAllUsersAction)
	usersRoute.GET("/:id", LoadUserByIdAction)
	usersRoute.DELETE("/:id", DeleteUserByIdAction)
	usersRoute.PATCH("/:id", UpdateUserByIdAction)
}
