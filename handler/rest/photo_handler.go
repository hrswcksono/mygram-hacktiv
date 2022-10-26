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

// Add Photo godoc
// @Tags photo
// @Description Add data photo
// @ID Add-photo
// @Accept json
// @Produce json
// @Param RequestBody body dto.CreatePhotoRequest true "request body json"
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Success 201 {object} dto.CreatedPhotoResponse
// @Router /photos [post]
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
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusCreated, data)
}

// GetAll Photo godoc
// @Tags photo
// @Description Get all data photo
// @ID Get-all-photo
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Success 200 {array} dto.GetPhotoResponse
// @Router /photos [get]
func (u photoRestHandler) ReadAllPhoto(c *gin.Context) {
	photo, err := u.service.ReadAllPhoto()

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, photo)
}

// Update Photo godoc
// @Tags photo
// @Description Update data photo
// @ID Update-photo
// @Accept json
// @Produce json
// @Param RequestBody body dto.EditPhotoRequest true "request body json"
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param photoId path int true "photo's id"
// @Success 200 {object} dto.EditPhotoResponse
// @Router /photos/{photoId} [put]
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

	data, err1 := u.service.EditPhoto(&photo, photoId)

	if err1 != nil {
		c.JSON(err1.Status(), err1)
		return
	}

	c.JSON(http.StatusOK, data)
}

// Delete Photo godoc
// @Tags photo
// @Description Delete data photo
// @ID Delete-photo
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param photoId path int true "photo's id"
// @Success 200 {object} dto.DeleteResponse
// @Router /photos/{photoId} [delete]
func (u photoRestHandler) DeletePhoto(c *gin.Context) {
	photoId, err := helper.GetParamId(c, "photoId")

	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"msg": err.Error(),
			"err": "BAD_REQUEST",
		})
		return
	}

	err1 := u.service.DeletePhoto(photoId)

	if err1 != nil {
		c.JSON(err1.Status(), err1)
		return
	}

	c.JSON(http.StatusOK, dto.DeleteResponse{
		Message: "Your photo has been successfully deleted",
	})
}
