package rest

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/hrswcksono/mygram-hacktiv/database"
	"github.com/hrswcksono/mygram-hacktiv/repository/user_repository/user_pg"
	"github.com/hrswcksono/mygram-hacktiv/service"
)

const port = ":8080"

func StartApp() {
	database.StartDB()

	db := database.GetDB()

	userRepo := user_pg.NewUserPG(db)
	userService := service.NewUserService(userRepo)
	userRestHandler := newUserHandler(userService)

	authService := service.NewAuthService(userRepo)

	route := gin.Default()

	userRoute := route.Group("/users")
	{
		userRoute.POST("/login", userRestHandler.Login)
		userRoute.POST("/register", userRestHandler.Register)
		userRoute.Use(authService.Authentication())
		userRoute.PUT("/:userId", authService.UserAuthorization(), userRestHandler.UpdateUser)
		userRoute.DELETE("/", userRestHandler.DeleteUser)
	}

	fmt.Println("Server running on PORT =>", port)
	route.Run(port)
}
