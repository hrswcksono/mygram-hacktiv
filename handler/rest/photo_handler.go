package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hrswcksono/mygram-hacktiv/dto"
	"github.com/hrswcksono/mygram-hacktiv/entity"
	"github.com/hrswcksono/mygram-hacktiv/pkg/helper"
	"github.com/hrswcksono/mygram-hacktiv/service"
)

type photoRestHandler struct {
	service service.PhotoService
}

func NewPhotoHandler(photoService service.PhotoService) photoRestHandler {
	return photoRestHandler{
		service: photoService,
	}
}

func (u photoRestHandler) AddPhoto(c *gin.Context) {

	var photo dto.CreatePhotoRequest

	if err := c.ShouldBindJSON(&photo); err != nil {
		c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"msg": "invalid JSON request",
			"err": "BAD_REQUEST",
		})
		return
	}

	var userId int

	value, _ := c.MustGet("userData").(entity.User)

	userId = value.ID

	data, err := u.service.AddPhoto(&photo, userId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"msg": err.Error(),
			"err": "INTERNAL_SERVER_ERROR",
		})
		return
	}

	c.JSON(http.StatusCreated, data)
}

func (u photoRestHandler) ReadAllPhoto(c *gin.Context) {
	photo, err := u.service.ReadAllPhoto()

	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"msg": err.Error(),
			"err": "INTERNAL_SERVER_ERROR",
		})
		return
	}

	c.JSON(http.StatusOK, photo)
}

func (u photoRestHandler) EditPhoto(c *gin.Context) {
	photoId, err := helper.GetParamId(c, "photoId")

	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"msg": err.Error(),
			"err": "BAD_REQUEST",
		})
		return
	}

	var photo dto.EditPhotoRequest

	if err := c.ShouldBindJSON(&photo); err != nil {
		c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"msg": "invalid JSON request",
			"err": "BAD_REQUEST",
		})
		return
	}

	data, err := u.service.EditPhoto(&photo, photoId)

	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"msg": err.Error(),
			"err": "INTERNAL_SERVER_ERROR",
		})
		return
	}

	c.JSON(http.StatusOK, data)
}

func (u photoRestHandler) DeletePhoto(c *gin.Context) {
	photoId, err := helper.GetParamId(c, "photoId")

	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"msg": err.Error(),
			"err": "BAD_REQUEST",
		})
		return
	}

	err = u.service.DeletePhoto(photoId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"msg": err.Error(),
			"err": "INTERNAL_SERVER_ERROR",
		})
		return
	}

	c.JSON(http.StatusCreated, dto.DeleteResponse{
		Message: "Your photo has been successfully deleted",
	})
}
