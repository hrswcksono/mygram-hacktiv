package rest

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hrswcksono/mygram-hacktiv/dto"
	"github.com/hrswcksono/mygram-hacktiv/entity"
	"github.com/hrswcksono/mygram-hacktiv/pkg/helper"
	"github.com/hrswcksono/mygram-hacktiv/service"
)

type userRestHandler struct {
	service service.UserService
}

func newUserHandler(userService service.UserService) userRestHandler {
	return userRestHandler{
		service: userService,
	}
}

// Register godoc
// @Tags auth
// @Description Register to mygram
// @ID register-user
// @Accept json
// @Produce json
// @Param RequestBody body dto.RegisterRequest true "request body json"
// @Success 201 {object} dto.RegisterResponse
// @Router /users/register [post]
func (u userRestHandler) Register(c *gin.Context) {
	var user dto.RegisterRequest

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"msg": "invalid JSON request",
			"err": "BAD_REQUEST",
		})
		return
	}

	data, err := u.service.Register(&user)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusCreated, data)
}

// Login godoc
// @Tags auth
// @Description Login to mygram
// @ID login-user
// @Accept json
// @Produce json
// @Param RequestBody body dto.LoginRequest true "request body json"
// @Success 200 {object} dto.LoginResponse
// @Router /users/login [post]
func (u userRestHandler) Login(c *gin.Context) {
	var user *dto.LoginRequest

	fmt.Println(user)

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"msg": "invalid JSON request",
			"err": "BAD_REQUEST",
		})
		return
	}

	data, err := u.service.Login(user)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, data)
}

// Update User godoc
// @Tags user
// @Description Update data user
// @ID Update-user
// @Accept json
// @Produce json
// @Param RequestBody body dto.UserEditRequest true "request body json"
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param userId path int true "user's id"
// @Success 200 {object} dto.UserEditResponse
// @Router /users/{userId} [put]
func (u userRestHandler) UpdateUser(c *gin.Context) {
	userId, err := helper.GetParamId(c, "userId")

	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"msg": err.Error(),
			"err": "BAD_REQUEST",
		})
		return
	}

	var editRequest *dto.UserEditRequest

	if err := c.ShouldBindJSON(&editRequest); err != nil {
		c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"msg": "invalid JSON request",
			"err": "BAD_REQUEST",
		})
		return
	}

	user, errs := u.service.UpdateUser(editRequest, userId)

	if errs != nil {
		c.JSON(errs.Status(), errs)
		return
	}

	c.JSON(http.StatusOK, user)

}

// Delete User godoc
// @Tags user
// @Description Delete data user
// @ID Delete-user
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Success 200 {object} dto.DeleteResponse
// @Router /users/ [delete]
func (u userRestHandler) DeleteUser(c *gin.Context) {

	var userId int

	value, _ := c.MustGet("userData").(entity.User)

	userId = value.ID

	err := u.service.DeleteUser(userId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"msg": err.Error(),
			"err": "INTERNAL_SERVER_ERROR",
		})
		return
	}

	c.JSON(err.Status(), dto.DeleteResponse{
		Message: "Your account has been successfully deleted",
	})
}
