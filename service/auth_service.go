package service

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hrswcksono/mygram-hacktiv/entity"
	"github.com/hrswcksono/mygram-hacktiv/pkg/auth_helper"
	"github.com/hrswcksono/mygram-hacktiv/pkg/helper"
	"github.com/hrswcksono/mygram-hacktiv/repository/comment_repository"
	"github.com/hrswcksono/mygram-hacktiv/repository/photo_repository"
	"github.com/hrswcksono/mygram-hacktiv/repository/social_media_repository"
	"github.com/hrswcksono/mygram-hacktiv/repository/user_repository"
)

type AuthService interface {
	Authentication() gin.HandlerFunc
	UserAuthorization() gin.HandlerFunc
	PhotoAuthorization() gin.HandlerFunc
	CommentAuthorization() gin.HandlerFunc
	SMediaAuthorization() gin.HandlerFunc
}

type authService struct {
	userRepo    user_repository.UserRepository
	photoRepo   photo_repository.PhotoRepository
	commentRepo comment_repository.CommentRepository
	smediaRepo  social_media_repository.SocialMediaRepository
}

func NewAuthService(userRepo user_repository.UserRepository, photoRepo photo_repository.PhotoRepository, commentRepo comment_repository.CommentRepository, smediaRepo social_media_repository.SocialMediaRepository) AuthService {
	return &authService{
		userRepo:    userRepo,
		photoRepo:   photoRepo,
		commentRepo: commentRepo,
		smediaRepo:  smediaRepo,
	}
}

func (a *authService) Authentication() gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		var user = &entity.User{}

		mapClaims, err := auth_helper.VerifyToken(ctx)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"err_message": err.Error(),
			})
		}

		if exp, ok := mapClaims["exp"].(float64); !ok {
			panic("login to proceed")
		} else {
			if int64(exp)-time.Now().Unix() <= 0 {
				panic("login to proceed")
			}
		}

		if v, ok := mapClaims["id"].(float64); !ok {
			panic("login to proceed")
		} else {
			user.ID = int(v)
		}

		if v, ok := mapClaims["email"].(string); !ok {
			panic("login to proceed")
		} else {
			user.Email = v
		}

		ctx.Set("userData", *user)
		ctx.Next()
	})
}

func (a *authService) UserAuthorization() gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		var userData entity.User

		if value, ok := ctx.MustGet("userData").(entity.User); !ok {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"err_message": "unauthorized",
			})
			return
		} else {
			userData = value
		}

		userIdParam, err := helper.GetParamId(ctx, "userId")

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"err_message": "invalid params",
			})
			return
		}

		user, err1 := a.userRepo.GetUserByID(userIdParam)

		if err1 != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"err_message": "not found",
			})
			return
		}

		if user.ID != userData.ID {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"err_message": "forbidden access",
			})
			return
		}

		ctx.Next()
	})
}

func (a *authService) PhotoAuthorization() gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		var userData entity.User

		if value, ok := ctx.MustGet("userData").(entity.User); !ok {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"err_message": "unauthorized",
			})
			return
		} else {
			userData = value
		}

		photoParam, err := helper.GetParamId(ctx, "photoId")

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"err_message": "invalid params",
			})
			return
		}

		photo, _ := a.photoRepo.GetPhotoByID(photoParam)

		if photo.ID == 0 {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"err_message": "not found",
			})
			return
		}

		if photo.UserID != userData.ID {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"err_message": "forbidden access",
			})
			return
		}

		ctx.Next()
	})
}

func (a *authService) CommentAuthorization() gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		var userData entity.User

		if value, ok := ctx.MustGet("userData").(entity.User); !ok {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"err_message": "unauthorized",
			})
			return
		} else {
			userData = value
		}

		commentParam, err := helper.GetParamId(ctx, "commentId")

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"err_message": "invalid params",
			})
			return
		}

		data, _ := a.commentRepo.GetCommentByID(commentParam)

		if data.ID == 0 {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"err_message": "not found",
			})
			return
		}

		if data.UserID != userData.ID {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"err_message": "forbidden access",
			})
			return
		}

		ctx.Next()
	})
}

func (a *authService) SMediaAuthorization() gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		var userData entity.User

		if value, ok := ctx.MustGet("userData").(entity.User); !ok {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"err_message": "unauthorized",
			})
			return
		} else {
			userData = value
		}

		smediaParam, err := helper.GetParamId(ctx, "socialMediaId")

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"err_message": "invalid params",
			})
			return
		}

		data, _ := a.smediaRepo.GetSocialMediaByID(smediaParam)

		if data.ID == 0 {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"err_message": "not found",
			})
			return
		}

		if data.UserID != userData.ID {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"err_message": "forbidden access",
			})
			return
		}

		ctx.Next()
	})
}
