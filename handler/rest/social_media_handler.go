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

// Add Social Media godoc
// @Tags socialmedia
// @Description Add data social media
// @ID Add-social-media
// @Accept json
// @Produce json
// @Param RequestBody body dto.CreateSMediaRequest true "request body json"
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Success 201 {object} dto.CreateSMediaResponse
// @Router /socialmedias [post]
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

	data, err1 := u.service.AddSMedia(&smedia, userId)

	if err1 != nil {
		c.JSON(err1.Status(), err1)
		return
	}

	c.JSON(http.StatusCreated, data)
}

// GetAll Social Media godoc
// @Tags socialmedia
// @Description get all data social media
// @ID Get-all-social-media
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Success 200 {object} dto.GetSMediaResponse
// @Router /socialmedias [get]
func (u smediaRestHandler) ReadAllSMedia(c *gin.Context) {
	smedia, err := u.service.ReadAllSMedia()

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, smedia)
}

// Update Social Media godoc
// @Tags socialmedia
// @Description Update data social media
// @ID Update-social-media
// @Accept json
// @Produce json
// @Param RequestBody body dto.EditSMediaRequest true "request body json"
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param socialMediaId path int true "social media's id"
// @Success 200 {object} dto.EditSMediaResponse
// @Router /socialmedias/{socialMediaId} [put]
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

	data, err1 := u.service.EditSMedia(&smedia, smediaId)

	if err1 != nil {
		c.JSON(err1.Status(), err1)
		return
	}

	c.JSON(http.StatusOK, data)
}

// Delete Social Media godoc
// @Tags socialmedia
// @Description Delete data social media
// @ID Delete-social-media
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param socialMediaId path int true "social media's id"
// @Success 200 {object} dto.DeleteResponse
// @Router /socialmedias/{socialMediaId} [delete]
func (u smediaRestHandler) DeleteSMedia(c *gin.Context) {
	smediaId, err := helper.GetParamId(c, "socialMediaId")

	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"msg": err.Error(),
			"err": "BAD_REQUEST",
		})
		return
	}

	err1 := u.service.DeleteSMedia(smediaId)

	if err1 != nil {
		c.JSON(err1.Status(), err1)
		return
	}

	c.JSON(http.StatusOK, dto.DeleteResponse{
		Message: "Your social media has been successfully deleted",
	})
}
