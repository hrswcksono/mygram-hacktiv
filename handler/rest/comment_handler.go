package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hrswcksono/mygram-hacktiv/dto"
	"github.com/hrswcksono/mygram-hacktiv/entity"
	"github.com/hrswcksono/mygram-hacktiv/pkg/helper"
	"github.com/hrswcksono/mygram-hacktiv/service"
)

type commentResthandler struct {
	service service.CommentService
}

func NewCommentHanlder(commentService service.CommentService) commentResthandler {
	return commentResthandler{
		service: commentService,
	}
}

// Add Comment godoc
// @Tags comment
// @Description Add data comment
// @ID Add-comment
// @Accept json
// @Produce json
// @Param RequestBody body dto.CreateCommentRequest true "request body json"
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Success 201 {object} dto.CreateCommentResponse
// @Router /comments [post]
func (u commentResthandler) AddComment(c *gin.Context) {
	var comment dto.CreateCommentRequest

	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"msg": "invalid JSON request",
			"err": "BAD_REQUEST",
		})
		return
	}

	var userId int

	value, _ := c.MustGet("userData").(entity.User)

	userId = value.ID

	data, err1 := u.service.AddComment(&comment, userId)

	if err1 != nil {
		c.JSON(err1.Status(), err1)
		return
	}

	c.JSON(http.StatusCreated, data)
}

// GetAll Comment godoc
// @Tags comment
// @Description Get all data comment
// @ID Get-all-comment
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Success 200 {array} dto.GetCommentResponse
// @Router /comments [get]
func (u commentResthandler) ReadAllComment(c *gin.Context) {
	comment, err := u.service.ReadAllComment()

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, comment)
}

// Update Comment godoc
// @Tags comment
// @Description Update data comment
// @ID Update-comment
// @Accept json
// @Produce json
// @Param RequestBody body dto.EditCommentRequest true "request body json"
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param commentId path int true "comment's id"
// @Success 200 {object} dto.EditCommentResponse
// @Router /comments/{commentId} [put]
func (u commentResthandler) EditComment(c *gin.Context) {
	commentId, err := helper.GetParamId(c, "commentId")

	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"msg": err.Error(),
			"err": "BAD_REQUEST",
		})
		return
	}

	var comment dto.EditCommentRequest

	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"msg": "invalid JSON request",
			"err": "BAD_REQUEST",
		})
		return
	}

	data, err1 := u.service.EditComment(&comment, commentId)

	if err1 != nil {
		c.JSON(err1.Status(), err1)
		return
	}

	c.JSON(http.StatusOK, data)
}

// Delete Comment godoc
// @Tags comment
// @Description Delete data comment
// @ID Delete-comment
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param commentId path int true "comment's id"
// @Success 200 {object} dto.DeleteResponse
// @Router /comments/{commentId} [delete]
func (u commentResthandler) DeleteComment(c *gin.Context) {
	commentId, err := helper.GetParamId(c, "commentId")

	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"msg": err.Error(),
			"err": "BAD_REQUEST",
		})
		return
	}

	err1 := u.service.DeleteComment(commentId)

	if err1 != nil {
		c.JSON(err1.Status(), err1)
		return
	}

	c.JSON(http.StatusOK, dto.DeleteResponse{
		Message: "Your comment has been successfully deleted",
	})
}
