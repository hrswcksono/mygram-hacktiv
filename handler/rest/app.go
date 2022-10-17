package rest

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/hrswcksono/mygram-hacktiv/database"
	"github.com/hrswcksono/mygram-hacktiv/repository/comment_repository/comment_pg"
	"github.com/hrswcksono/mygram-hacktiv/repository/photo_repository/photo_pg"
	"github.com/hrswcksono/mygram-hacktiv/repository/social_media_repository/social_media_pg"
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

	photoRepo := photo_pg.NewPhotoPG(db)
	photoService := service.NewPhotoService(photoRepo)
	photoRestHandler := NewPhotoHandler(photoService)

	commentRepo := comment_pg.NewCommentPG(db)
	commentService := service.NewCommentService(commentRepo)
	commentRestHandler := NewCommentHanlder(commentService)

	smediaRepo := social_media_pg.NewSMediaPG(db)
	smediaService := service.NewSMediaService(smediaRepo)
	smediaRestHandler := NewSMediaHandler(smediaService)

	authService := service.NewAuthService(userRepo, photoRepo, commentRepo, smediaRepo)

	route := gin.Default()

	userRoute := route.Group("/users")
	{
		userRoute.POST("/login", userRestHandler.Login)
		userRoute.POST("/register", userRestHandler.Register)
		userRoute.Use(authService.Authentication())
		userRoute.PUT("/:userId", authService.UserAuthorization(), userRestHandler.UpdateUser)
		userRoute.DELETE("/", userRestHandler.DeleteUser)
	}

	photoRoute := route.Group("/photos")
	{
		photoRoute.Use(authService.Authentication())
		photoRoute.POST("/", photoRestHandler.AddPhoto)
		photoRoute.GET("/", photoRestHandler.ReadAllPhoto)
		photoRoute.PUT("/:photoId", authService.PhotoAuthorization(), photoRestHandler.EditPhoto)
		photoRoute.DELETE("/:photoId", authService.PhotoAuthorization(), photoRestHandler.DeletePhoto)
	}

	commentRoute := route.Group("/comments")
	{
		commentRoute.Use(authService.Authentication())
		commentRoute.POST("/", commentRestHandler.AddComment)
		commentRoute.GET("/", commentRestHandler.ReadAllComment)
		commentRoute.PUT("/:commentId", authService.CommentAuthorization(), commentRestHandler.EditComment)
		commentRoute.DELETE("/:commentId", authService.CommentAuthorization(), commentRestHandler.DeleteComment)
	}

	smediaRoute := route.Group("/socialmedias")
	{
		smediaRoute.Use(authService.Authentication())
		smediaRoute.POST("/", smediaRestHandler.AddSMedia)
		smediaRoute.GET("/", smediaRestHandler.ReadAllSMedia)
		smediaRoute.PUT("/:socialMediaId", authService.SMediaAuthorization(), smediaRestHandler.EditSMedia)
		smediaRoute.DELETE("/:socialMediaId", authService.SMediaAuthorization(), smediaRestHandler.DeleteSMedia)
	}

	fmt.Println("Server running on PORT =>", port)
	route.Run(port)
}
