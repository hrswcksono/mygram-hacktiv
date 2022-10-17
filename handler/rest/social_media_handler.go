package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hrswcksono/mygram-hacktiv/dto"
	"github.com/hrswcksono/mygram-hacktiv/entity"
	"github.com/hrswcksono/mygram-hacktiv/pkg/helper"
	"github.com/hrswcksono/mygram-hacktiv/service"
)

type smediaRestHandler struct {
	service service.SMediaService
}

func NewSMediaHandler(smediaService service.SMediaService) smediaRestHandler {
	return smediaRestHandler{
		service: smediaService,
	}
}

func (u smediaRestHandler) AddSMedia(c *gin.Context) {
	var smedia dto.CreateSMediaRequest

	if err := c.ShouldBindJSON(&smedia); err != nil {
		c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"msg": "invalid JSON request",
			"err": "BAD_REQUEST",
		})
		return
	}

	var userId int

	value, _ := c.MustGet("userData").(entity.User)

	userId = value.ID

	data, err := u.service.AddSMedia(&smedia, userId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"msg": err.Error(),
			"err": "INTERNAL_SERVER_ERROR",
		})
		return
	}

	c.JSON(http.StatusCreated, data)
}

func (u smediaRestHandler) ReadAllSMedia(c *gin.Context) {
	smedia, err := u.service.ReadAllSMedia()

	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"msg": err.Error(),
			"err": "INTERNAL_SERVER_ERROR",
		})
		return
	}

	c.JSON(http.StatusOK, smedia)
}

func (u smediaRestHandler) EditSMedia(c *gin.Context) {
	smediaId, err := helper.GetParamId(c, "socialMediaId")

	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"msg": err.Error(),
			"err": "BAD_REQUEST",
		})
		return
	}

	var smedia dto.EditSMediaRequest

	if err := c.ShouldBindJSON(&smedia); err != nil {
		c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"msg": "invalid JSON request",
			"err": "BAD_REQUEST",
		})
		return
	}

	data, err := u.service.EditSMedia(&smedia, smediaId)

	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"msg": err.Error(),
			"err": "INTERNAL_SERVER_ERROR",
		})
		return
	}

	c.JSON(http.StatusOK, data)
}

func (u smediaRestHandler) DeleteSMedia(c *gin.Context) {
	smediaId, err := helper.GetParamId(c, "socialMediaId")

	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"msg": err.Error(),
			"err": "BAD_REQUEST",
		})
		return
	}

	err = u.service.DeleteSMedia(smediaId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"msg": err.Error(),
			"err": "INTERNAL_SERVER_ERROR",
		})
		return
	}

	c.JSON(http.StatusCreated, dto.DeleteResponse{
		Message: "Your social media has been successfully deleted",
	})
}
