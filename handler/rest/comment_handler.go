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

	data, err := u.service.AddComment(&comment, userId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"msg": err.Error(),
			"err": "INTERNAL_SERVER_ERROR",
		})
		return
	}

	c.JSON(http.StatusCreated, data)
}

func (u commentResthandler) ReadAllComment(c *gin.Context) {
	comment, err := u.service.ReadAllComment()

	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"msg": err.Error(),
			"err": "INTERNAL_SERVER_ERROR",
		})
		return
	}

	c.JSON(http.StatusOK, comment)
}

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

	data, err := u.service.EditComment(&comment, commentId)

	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"msg": err.Error(),
			"err": "INTERNAL_SERVER_ERROR",
		})
		return
	}

	c.JSON(http.StatusOK, data)
}

func (u commentResthandler) DeleteComment(c *gin.Context) {
	commentId, err := helper.GetParamId(c, "commentId")

	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"msg": err.Error(),
			"err": "BAD_REQUEST",
		})
		return
	}

	err = u.service.DeleteComment(commentId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"msg": err.Error(),
			"err": "INTERNAL_SERVER_ERROR",
		})
		return
	}

	c.JSON(http.StatusCreated, dto.DeleteResponse{
		Message: "Your comment has been successfully deleted",
	})
}
